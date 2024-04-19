package model

import (
	"gorm.io/gorm"
	"myginblog/utils/errMsg"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100); not null" json:"title"`
	Cid      int      `gorm:"type:int; not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(100); not null" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

//// CheckArticle 检查分类是否存在
//func CheckArticle(name string) (code int) {
//	var category Category
//	db.Select("id").Where("name = ?", name).First(&category)
//	if category.ID >= 1 {
//		return errMsg.ERROR_CATENAME_USED
//	}
//	return errMsg.SUCCESS
//}

// CreateArticle 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// GetCateArticle 查询分类下的所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var article Article
	var total int64
	err = db.Preload("Category").Where("cid = ?", id).Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&articles).Error
	db.Model(&article).Where("cid = ?", id).Count(&total)
	if err != nil {
		return articles, errMsg.ERROR, 0
	}
	return articles, errMsg.SUCCESS, total
}

// GetArticle 查询文章列表
func GetArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var article Article
	var total int64
	//db.Model(&article).Count(&total)
	if title == "" {
		err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
		db.Model(&article).Count(&total)
	} else {
		err = db.Preload("Category").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
		db.Model(&article).Where("title LIKE ?", title+"%").Count(&total)
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errMsg.ERROR, 0
	}
	return articles, errMsg.SUCCESS, total
}

// GetSignalArtInfo 查询单个文章信息
func GetSignalArtInfo(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, errMsg.ERROR_ART_NOT_EXIST
	}
	return article, errMsg.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, article *Article) int {
	var data Article
	articleMap := make(map[string]interface{})
	articleMap["title"] = article.Title
	articleMap["cid"] = article.Cid
	articleMap["content"] = article.Content
	articleMap["desc"] = article.Desc
	articleMap["img"] = article.Img
	err = db.Model(&data).Where(" id = ?", id).Updates(articleMap).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}
