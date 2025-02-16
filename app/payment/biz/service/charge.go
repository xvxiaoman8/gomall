package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xvxiaoman8/gomall/app/payment/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/payment/biz/model"
	payment "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	paymentprocessor, err := CreatePaymentProcessor("creditcard", req.CreditCard)
	if err != nil {
		return nil, err
	}
	err = paymentprocessor.Validate()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}

	// 随机生成交易ID
	transactionId, err := paymentprocessor.Execute()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	// 保存交易记录

	err = model.CreatePaymentAmount(mysql.DB, s.ctx, &model.PaymentAmount{
		TransactionID: transactionId.String(),
		Amount:        req.Amount,
	})

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserID:        req.UserId,
		OrderID:       req.OrderId,
		TransactionID: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	//交易成功，进行应答

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
