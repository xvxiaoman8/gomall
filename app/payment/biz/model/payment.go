package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserID        uint32    `json:"user_id"`
	OrderID       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}
type PaymentAmount struct {
	IsRefunded    bool    `json:"is_refunded"` // 是否已退款
	Amount        float32 `json:"amount"`
	TransactionID string  `json:"transaction_id"`
	gorm.Model
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, paymentLog *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}
func CreatePaymentAmount(db *gorm.DB, ctx context.Context, paymentAmount *PaymentAmount) error {
	return db.WithContext(ctx).Model(&PaymentAmount{}).Create(paymentAmount).Error
}
func UpdatePaymentAmount(db *gorm.DB, ctx context.Context, paymentAmount *PaymentAmount) error {
	return db.WithContext(ctx).Model(&PaymentAmount{}).Where("transaction_id = ?", paymentAmount.TransactionID).Updates(paymentAmount).Error
}
func GetPaymentAmountByTxnID(db *gorm.DB, ctx context.Context, transactionId string) (PaymentAmount, error) {
	paymentAmount := PaymentAmount{}
	err := db.WithContext(ctx).Model(&PaymentAmount{}).Where("transaction_id = ?", transactionId).First(&paymentAmount).Error
	return paymentAmount, err
}
