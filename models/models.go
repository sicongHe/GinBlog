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


type Tag struct {
	gorm.Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (t *Tag)CreateTable() {
	if !DB.Migrator().HasTable(t) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(t)
	}
}

func InitDB() {
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


func GetTags(pageNum int, pageSize int, maps interface{})(tags []Tag){
	DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{})(count int64) {
	DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}