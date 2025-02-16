package service

import (
	"context"
	"github.com/xvxiaoman8/gomall/app/product/biz/dal/mysql"
	"github.com/xvxiaoman8/gomall/app/product/biz/model"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/product"
)

// ListProductsService 用于处理产品列表服务
type ListProductsService struct {
	ctx context.Context
}

// 创建一个新的ListProductsService实例，并返回该实例的指针
func NewListProductsService(ctx context.Context) *ListProductsService {
	// 创建一个新的ListProductsService实例，并将传入的context赋值给ctx字段
	return &ListProductsService{ctx: ctx}
}

// Run 执行获取产品列表的业务逻辑
//func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
//	// Finish your business logic.
//	// 根据类别名称从数据库获取产品列表
//	c, err := model.GetProductsByCategoryName(mysql.DB, s.ctx, req.CategoryName)
//	if err != nil {
//		return nil, err
//	}
//	// 初始化产品列表响应对象
//	resp = &product.ListProductsResp{}
//	// 遍历获取到的类别及其产品，将产品信息添加到响应对象中
//	for _, v1 := range c {
//		for _, v := range v1.Products {
//			resp.Products = append(resp.Products, &product.Product{Id: uint32(v.ID), Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
//		}
//	}
//
//	return resp, nil
//}

func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	c, err := model.GetProductsByCategoryName(mysql.DB, s.ctx, req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{Id: uint32(v.ID), Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
		}
	}

	return resp, nil
}

//func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
//	if req == nil {
//		return nil, errors.New("request is nil")
//	}
//	// 根据类别名称从数据库获取产品列表
//	c, err := model.GetProductsByCategoryName(mysql.DB, s.ctx, req.CategoryName)
//	if err != nil {
//		return nil, err
//	}
//	// 初始化产品列表响应对象
//	resp = &product.ListProductsResp{
//		Products: make([]*product.Product, 0),
//	}
//	// 遍历获取到的类别及其产品，将产品信息添加到响应对象中
//	for _, v1 := range c {
//		for _, v := range v1.Products {
//			resp.Products = append(resp.Products, &product.Product{Id: uint32(v.ID), Name: v.Name, Description: v.Description, Picture: v.Picture, Price: v.Price})
//		}
//	}
//
//	return resp, nil
//}
