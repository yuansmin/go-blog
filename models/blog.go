package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Blog struct {
	Id      int
	Title   string
	Content string
}

func init() {
	// register model
	orm.RegisterModel(new(Blog))

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "/Users/fancy/work/learn/go/src/myblog/blog.db", 30)
}
