package middleware

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/auth"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {

			// 判断是否为分发token请求
			if _, isDeliver := req.(*auth.DeliverTokenReq); isDeliver {
				// 直接放行分发token请求
				return next(ctx, req, resp)
			}

			// 从上下文中获取角色和请求信息

			userRole, _ := ctx.Value("user_role").(string)
			resource, _ := ctx.Value("resourse").(string)
			action, _ := ctx.Value("action").(string)

			// 执行权限检查
			ok, err := enforcer.Enforce(userRole, resource, action)
			if err != nil || !ok {
				return errors.New("access denied")
			}

			return next(ctx, req, resp)
		}
	}
}
