package service

import (
	"errors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"strconv"
)

type PaymentProcessor interface {
	Validate() error             // 验证支付参数
	Execute() (uuid.UUID, error) // 发起支付，返回交易ID
	Name() string                // 支付方式名称（如 "alipay", "creditcard"）
}

type CreditCardProcessor struct {
	Card creditcard.Card
}

func (p *CreditCardProcessor) Validate() error {
	return p.Card.Validate()
}

func (p *CreditCardProcessor) Execute() (uuid.UUID, error) {
	// 调用信用卡支付网关接口...
	return uuid.NewRandom()
}

func (p *CreditCardProcessor) Name() string {
	return "creditcard"
}

func CreatePaymentProcessor(method string, info *payment.CreditCardInfo) (PaymentProcessor, error) {
	switch method {
	case "creditcard":
		card := creditcard.Card{
			Number: info.CreditCardNumber,
			Cvv:    strconv.Itoa(int(info.CreditCardCvv)),
			Month:  strconv.Itoa(int(info.CreditCardExpirationMonth)),
			Year:   strconv.Itoa(int(info.CreditCardExpirationYear)),
		}
		return &CreditCardProcessor{Card: card}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}
