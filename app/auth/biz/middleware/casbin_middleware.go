package middleware

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) error {
			// 从上下文中获取角色和请求信息
			role, _ := ctx.Value("role").(string)
			method := "your_method_name" // 根据请求获取具体方法
			path := "your_resource_path"

			// 执行权限检查
			ok, err := enforcer.Enforce(role, path, method)
			if err != nil || !ok {
				return errors.New("access denied")
			}

			return next(ctx, req, resp)
		}
	}
}
