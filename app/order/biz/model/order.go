// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}

func GetOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (order Order, err error) {
	err = db.Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

func UpdateOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string, updates map[string]interface{}) (err error) {
	err = db.Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Updates(updates).Error
	return
}

func DeleteOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (err error) {
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 先删除关联的 OrderItem 记录
	if err := tx.Where("order_id_refer = ?", orderId).Delete(&OrderItem{}).Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}

	// 再删除 Order 记录
	if err := tx.Where(&Order{UserId: userId, OrderId: orderId}).Delete(&Order{}).Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
