package model

import (
	"gorm.io/gorm"
	"myginblog/utils/errMsg"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

// CheckCategory 检查分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID >= 1 {
		return errMsg.ERROR_CATENAME_USED
	}
	return errMsg.SUCCESS
}

// CreateCategory 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// GetCate 查询单个分类信息
func GetCate(id int) (Category, int) {
	var category Category
	err = db.Limit(1).Where("ID = ?", id).Find(&category).Error
	if err != nil {
		return category, errMsg.ERROR
	}
	return category, errMsg.SUCCESS
}

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var category []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&category).Error
	db.Model(&category).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return category, total
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// EditCategory 编辑分类
func EditCategory(id int, category *Category) int {
	var data Category
	categoryMap := make(map[string]interface{})
	categoryMap["name"] = category.Name
	err = db.Model(&data).Where(" id = ?", id).Updates(categoryMap).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}
