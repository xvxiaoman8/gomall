package model

import (
	"context"
	"gorm.io/gorm"
)

// Category 定义分类结构体
type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products" gorm:"many2many:product_category"`
}

// TableName 定义 Category 结构体对应的数据库表名
func (c Category) TableName() string {
	return "category"
}

// CategoryQuery 定义分类查询结构体
type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// GetProductsByCategoryName 根据分类名称获取分类及其关联的产品
func (c CategoryQuery) GetProductsByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}

func GetProductsByCategoryName(db *gorm.DB, ctx context.Context, name string) (category []Category, err error) {
	err = db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	return category, err
}
