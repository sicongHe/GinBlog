package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func (a *Article)CreateTable() {
	if !DB.Migrator().HasTable(a) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(a)
	}
}

func ExistArticleByID(id int) bool{
	var article Article
	DB.Select("id").Where("id = ?",id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface {}) (count int64){
	DB.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article){
	DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article){
	DB.Where("id  = ?",id).Preload("Tag").First(&article)
	return
}

func EditArticle(id int, data interface {}){
	var article Article
	DB.Preload("Tag").Where("id = ?",id).First(&article).Updates(data)
}

func AddArticle(data map[string]interface {}) bool{
	DB.Create(&Article{
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool{
	DB.Where("id = ?",id).Delete(&Article{})
	return true
}

func CleanAllArticle() bool {
	DB.Unscoped().Where("deleted_at != ? ", 0).Delete(&Article{})

	return true
}