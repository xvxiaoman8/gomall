package middleware

import "github.com/cloudwego/kitex/pkg/endpoint"

// ChainMiddleware 组合多个中间件为一个中间件链
func ChainMiddleware(middlewares ...endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
