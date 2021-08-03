package models

import (
	"fmt"
	"github.com/siconghe/blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var(
	DB *gorm.DB
)




func Setup() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix
	switch dbType {
	case "mysql":DB,err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)),&gorm.Config{ NamingStrategy: schema.NamingStrategy{
		TablePrefix: tablePrefix,   // 表名前缀，`User`表为`t_users`
		SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
	},})
	default:
		fmt.Println("请自行添加其他数据库的支持")
	}

	if err != nil {
		log.Fatalf("数据库打开失败 %v",err)
	}
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
}


