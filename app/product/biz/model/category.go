package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

// 定义Category结构体的TableName方法，返回表名为"category"
func (c Category) TableName() string {
	return "category"
}

// 根据分类名称获取产品
func GetProductsByCategoryName(db *gorm.DB, ctx context.Context, name string) (category []Category, err error) {
	// 使用WithContext方法将上下文传递给db对象
	err = db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	// 返回查询结果和错误信息
	return category, err
}
