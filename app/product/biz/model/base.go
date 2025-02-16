package model

import "time"

// Base 定义基础模型结构体
type Base struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
