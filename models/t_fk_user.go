package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TFkUser struct {
	Id     int    `orm:"column(id);auto"`
	Name   string `orm:"column(name);size(300);null"`
	UserId *TUser `orm:"column(user_id);rel(fk)"`
}

func (t *TFkUser) TableName() string {
	return "t_fk_user"
}

func init() {
	orm.RegisterModel(new(TFkUser))
}

// AddTFkUser insert a new TFkUser into database and returns
// last inserted Id on success.
func AddTFkUser(m *TFkUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTFkUserById retrieves TFkUser by Id. Returns error if
// Id doesn't exist
func GetTFkUserById(id int) (v *TFkUser, err error) {
	o := orm.NewOrm()
	v = &TFkUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err

}

// GetAllTFkUser retrieves all TFkUser matches certain condition. Returns empty list if
// no records exist
func GetAllTFkUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	//	o := orm.NewOrm()
	//	var arc []*TFkUser
	//	o.QueryTable("t_fk_user").RelatedSel().All(&arc) //使用RelatedSel将关联的arctype也查出来，也就是left join arctype as T1 on T1.id=go_archives.arctype_id
	//	fmt.Println(arc)
	//	for _, v := range arc {
	//		ml = append(ml, v)
	//	}
	//	return ml, nil

	o := orm.NewOrm()
	qs := o.QueryTable(new(TFkUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TFkUser
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTFkUser updates TFkUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateTFkUserById(m *TFkUser) (err error) {
	o := orm.NewOrm()
	v := TFkUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTFkUser deletes TFkUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTFkUser(id int) (err error) {
	o := orm.NewOrm()
	v := TFkUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TFkUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
