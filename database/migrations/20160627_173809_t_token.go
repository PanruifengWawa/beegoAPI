package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TToken_20160627_173809 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TToken_20160627_173809{}
	m.Created = "20160627_173809"
	migration.Register("TToken_20160627_173809", m)
}

// Run the migrations
func (m *TToken_20160627_173809) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE t_token(`id` int(11) NOT NULL AUTO_INCREMENT,`token` varchar(300) NOT NULL,PRIMARY KEY (`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;")
}

// Reverse the migrations
func (m *TToken_20160627_173809) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `t_token`")
}
