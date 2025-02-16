package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal/redis"
	"github.com/xvxiaoman8/gomall/app/product/model"
	product "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct { // 定义GetProductService结构体
	ctx context.Context // 定义ctx字段，用于存储请求上下文
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService { // 定义NewGetProductService函数，用于创建GetProductService实例
	return &GetProductService{ctx: ctx} // 返回新的GetProductService实例指针，上下文为传入的ctx
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) { // 定义Run方法，接收产品请求，返回产品响应和错误信息
	// Finish your business logic.
	if req.Id == 0 { // 检查请求的产品ID是否为0
		return nil, kerrors.NewBizStatusError(40000, "product id is required") // 如果ID为0，返回错误信息，表示产品ID是必需的
	} //发现货物编号错误

	p, err := model.NewCachedProductQuery(model.NewProductQuery(s.ctx, mysql.DB), redis.RedisClient).GetById(int(req.Id)) // 创建CachedProductQuery对象，并调用GetById方法获取指定ID的产品信息
	if err != nil {                                                                                                       // 检查是否有错误发生
		return nil, err // 如果有错误，返回nil和错误信息
	}
	return &product.GetProductResp{ // 创建并返回产品响应对象
		Product: &product.Product{ // 定义响应中的产品信息
			Id:          uint32(p.ID),  // 设置产品ID，类型转换为uint32
			Picture:     p.Picture,     // 设置产品图片
			Price:       p.Price,       // 设置产品价格
			Description: p.Description, // 设置产品描述
			Name:        p.Name,        // 设置产品名称
		},
	}, err // 返回产品响应和错误信息，这里err应为nil，因为前面已经处理了错误情况
}
