package models

import (
	"gorm.io/gorm"
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

func GetTags(pageNum int, pageSize int, maps interface{})(tags []Tag){
	DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTag(name string) (tag Tag){
	DB.Where("name = ?", name).First(&tag)
	return
}

func GetTagByID(id int)(tag Tag) {
	DB.Where("id = ?",id).First(&tag)
	return
}

func GetTagTotal(maps interface{})(count int64) {
	DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func AddTag(name string,state int,createdBy string) bool {
	DB.Create(&Tag{Name:name,State: state,CreatedBy: createdBy})
	return true
}

func ExistTagByName(name string) bool {
	var tag Tag
	DB.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	DB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func EditTag(id int,data interface{}) bool{
	DB.Model(&Tag{}).Where("id = ?",id).Updates(data)
	return true
}

func DeleteTag(id int) bool{
	DB.Where("id = ?",id).Delete(&Tag{})
	return true
}

func CleanAllTag() bool {
	DB.Unscoped().Where("deleted_at != ? ", 0).Delete(&Tag{})
	return true
}