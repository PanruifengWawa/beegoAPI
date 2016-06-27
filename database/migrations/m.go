package main

import (
	"os"

	"github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/godb")
}

func main() {
	task := "upgrade"
	switch task {
	case "upgrade":
		if err := migration.Upgrade(0); err != nil {
			os.Exit(2)
		}
	case "rollback":
		if err := migration.Rollback("TUser_20160627_172702"); err != nil {
			os.Exit(2)
		}
	case "reset":
		if err := migration.Reset(); err != nil {
			os.Exit(2)
		}
	case "refresh":
		if err := migration.Refresh(); err != nil {
			os.Exit(2)
		}
	}
}
