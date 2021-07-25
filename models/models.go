package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/siconghe/blog/pkg/setting"
	"log"
)

var(
	db *gorm.DB
)

type Model struct {
	gorm.Model
}
func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	section,err :=setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal("配置读取失败，配置项：[database]")
	}
	dbType = section.Key("TYPE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	host = section.Key("HOST").String()
	tablePrefix = section.Key("TABLE_PREFIX").String()
	db,err = gorm.Open(dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local\"",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Fatalf("数据库打开失败 %v",err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix+ defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
