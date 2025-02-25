package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	redislock "github.com/jefferyjob/go-redislock"
	"github.com/xvxiaoman8/gomall/app/checkout/biz/dal/redis"
	"github.com/xvxiaoman8/gomall/app/checkout/infra/rpc"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/order"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/payment"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product"
	"strconv"
	"time"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	fmt.Println(req)
	if rpc.CartClient == nil {
		fmt.Println("CartClient is nil")
	}
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})

	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return nil, err
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return nil, err
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Cart.Items {
		var (
			cost       float32
			stockInt   int
			stockInt32 int32
			stockStr   string
			ok         bool
		)
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		// 初始化库存
		err = InitStock(s.ctx, int(p.Id), 100)
		if err != nil {
			klog.Error(err)
			return nil, err
		}
		fmt.Println("init stock success")
		//查redis库存
		//加上分布式锁
		//十秒超时
		ctx, cancel := context.WithCancel(s.ctx)
		lock := redis.NewRedisLock(strconv.Itoa(int(p.Id)), ctx, redislock.WithTimeout(time.Duration(10)*time.Second), redislock.WithAutoRenew())
		err = lock.Lock()
		if err != nil {
			klog.Error(err)
			cancel()
			return nil, err
		}
		unlockSignal := make(chan bool, 1)
		//开启协程进行锁的自动续期，如果业务结束则退出
		go func() {
			defer cancel()
			select {
			case <-unlockSignal:
				return
			}
		}()

		stock, err := redis.RedisDo(s.ctx, "GET", strconv.Itoa(int(p.Id))+"_stock")
		if err != nil {
			goto ERR
		}
		// 类型断言为 string
		stockStr, ok = stock.(string)
		if !ok {
			err = errors.New("stock is not a string")
			goto ERR
		}
		// 将 string 转换为 int
		stockInt, err = strconv.Atoi(stockStr)
		if err != nil {
			goto ERR
		}
		stockInt32 = int32(stockInt)
		//检查库存
		if stockInt32 <= cartItem.Quantity {
			err = errors.New("stock is not enough")
			goto ERR
		}
		//减库存
		stockInt32 -= cartItem.Quantity
		_, err = redis.RedisDo(s.ctx, "SET", strconv.Itoa(int(p.Id))+"_stock", stockInt)
		if err != nil {
			goto ERR
		}
		//释放锁
		err = lock.UnLock()
		if err != nil {
			klog.Error(err)
			panic(err)
		}
		unlockSignal <- true

		cost = p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	ERR:
		klog.Error(err)
		unlockSignal <- true
		return nil, err
	}

	//创建订单
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	var orderId string
	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}

	//支付订单
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
func InitStock(ctx context.Context, ID int, amount int) error {
	_, err := redis.RedisDo(ctx, "SET", strconv.Itoa(ID)+"_stock", amount)
	return err
}
