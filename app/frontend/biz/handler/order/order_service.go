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

package order

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xvxiaoman8/gomall/app/frontend/biz/service"
	"github.com/xvxiaoman8/gomall/app/frontend/biz/utils"
	common "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/xvxiaoman8/gomall/app/frontend/hertz_gen/frontend/order"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "order", hertzUtils.H{"error": err})
		return
	}

	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}

// UpdateOrderPage 处理修改订单页面请求（GET）
// @router /order/update/:order_id [GET]
func UpdateOrderPage(ctx context.Context, c *app.RequestContext) {
	var req order.UpdateOrderPageReq
	if err := c.BindAndValidate(&req); err != nil {
		// 直接渲染错误页面
		c.HTML(consts.StatusOK, "error", utils.WarpResponse(ctx, c, hertzUtils.H{
			"error": "请求参数错误",
		}))
		return
	}
	// 调用服务层获取订单数据
	resp, err := service.NewUpdateOrderPageService(ctx, c).Run(&req)
	if err != nil {
		// 渲染错误页面（带订单ID上下文）
		c.HTML(consts.StatusOK, "error", utils.WarpResponse(ctx, c, hertzUtils.H{
			"error":    "加载订单失败",
			"order_id": req.OrderId,
		}))
		return
	}
	// 成功时渲染修改页面模板
	c.HTML(consts.StatusOK, "update_order", utils.WarpResponse(ctx, c, resp))
}

// UpdateOrder 处理订单修改提交（POST）
// @router /order/update/:order_id [POST]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var req order.UpdateOrderReq
	OrderId := c.Param("order_id")
	if err := c.BindAndValidate(&req); err != nil {
		// 保持当前页面并显示错误信息
		c.HTML(consts.StatusOK, "update_order", utils.WarpResponse(ctx, c, hertzUtils.H{
			"error":    "表单数据错误",
			"form":     req, // 保留已填写的表单数据
			"order_id": req.OrderId,
		}))
		return
	}

	_, err := service.NewUpdateOrderService(ctx, c).Run(OrderId, &req)
	if err != nil {
		// 将之前传递上来的错误原因打印出来
		c.HTML(consts.StatusOK, "error", utils.WarpResponse(ctx, c, hertzUtils.H{
			"error":    err.Error(),
			"order_id": req.OrderId,
		}))
		return
	}

	// 成功时使用get访问order页面
	c.Redirect(consts.StatusOK, []byte("/order"))

}

// DeleteOrder .
// @router /order/delete/:order_id [POST]
func DeleteOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.DeleteOrderReq
	OrderId := c.Param("order_id")
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewDeleteOrderService(ctx, c).Run(OrderId, &req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusOK, []byte("/order"))
}
