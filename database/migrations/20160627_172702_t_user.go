package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TUser_20160627_172702 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TUser_20160627_172702{}
	m.Created = "20160627_172702"
	migration.Register("TUser_20160627_172702", m)
}

// Run the migrations
func (m *TUser_20160627_172702) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE t_user(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(300) NOT NULL,PRIMARY KEY (`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8")
}

// Reverse the migrations
func (m *TUser_20160627_172702) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `t_user`")
}
