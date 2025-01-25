// Code generated by Kitex v0.9.1. DO NOT EDIT.

package cartservice

import (
	server "github.com/cloudwego/kitex/server"
	cart "github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen/cart"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler cart.CartService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
