package util

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/siconghe/blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"testing"
)



func AssertErrShouldBeNil(err error,t *testing.T) {
	if err != nil {
		t.Errorf("Err应该为空！错误信息:%v",err)
	}
}
var mock sqlmock.Sqlmock
func MockDB() {
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if nil != err {
		log.Fatalf("Init sqlmock failed, err %v", err)
	}
	//结合gorm、sqlmock
	models.DB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn: db,
	}), &gorm.Config{ NamingStrategy: schema.NamingStrategy{
		TablePrefix: "blog_",   // 表名前缀，`User`表为`t_users`
		SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
	},})
	if nil != err {
		log.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
	models.DB.AutoMigrate(&models.Tag{})
}