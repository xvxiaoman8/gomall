package service

import (
	"context"
	"errors"
	"fmt"
	creditcard "github.com/durango/go-credit-card"
	"github.com/xvxiaoman8/gomall/app/payment/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/payment/biz/model"
	payment "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"strconv"
	"time"
)

type RefundService struct {
	ctx context.Context
} // NewRefundService new RefundService
func NewRefundService(ctx context.Context) *RefundService {
	return &RefundService{ctx: ctx}
}

// Run create note info
func (s *RefundService) Run(req *payment.RefundReq) (resp *payment.RefundResp, err error) {
	paymentprocessor, err := CreateRefundProcessor("creditcard", req.CreditCard)
	if err != nil {
		return nil, err
	}
	if err := paymentprocessor.ExecuteRefund(s.ctx, req.TransactionId); err != nil {
		return nil, err
	}
	resp = &payment.RefundResp{
		Message: "Refund successful",
	}
	return resp, nil
}

type RefundProcessor interface {
	ExecuteRefund(ctx context.Context, transactionId string) error
}
type CreditCardRefundProcessor struct {
	Card creditcard.Card
}

func CreateRefundProcessor(method string, info *payment.CreditCardInfo) (RefundProcessor, error) {
	switch method {
	case "creditcard":
		card := creditcard.Card{
			Number: info.CreditCardNumber,
			Cvv:    strconv.Itoa(int(info.CreditCardCvv)),
			Month:  strconv.Itoa(int(info.CreditCardExpirationMonth)),
			Year:   strconv.Itoa(int(info.CreditCardExpirationYear)),
		}
		return &CreditCardRefundProcessor{Card: card}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}

// ExecuteRefund  退款实现
func (p *CreditCardRefundProcessor) ExecuteRefund(ctx context.Context, transactionId string) error {
	// 获取原支付金额（需要根据 transactionId 查询数据库）
	paymentAmount, err := model.GetPaymentAmountByTxnID(mysql.DB, ctx, transactionId)
	if err != nil {
		return fmt.Errorf("failed to get original payment amount: %v", err)
	}
	if paymentAmount.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
		return fmt.Errorf("payment has expired")
	}
	if paymentAmount.IsRefunded {
		return fmt.Errorf("payment is refunded")
	}
	// 调用退款API
	if err := callCreditCardRefundAPI(ctx,
		paymentAmount,
	); err != nil {
		return fmt.Errorf("creditcard refund failed: %v", err)
	}
	return nil
}
func callCreditCardRefundAPI(ctx context.Context, paymentAmount model.PaymentAmount) error {
	paymentAmount.IsRefunded = true
	return model.UpdatePaymentAmount(mysql.DB, ctx, &paymentAmount)
}
