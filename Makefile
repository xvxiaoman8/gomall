export ROOT_MOD=github.com/xvxiaoman8/gomall
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module github.com/xvxiaoman8/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module  github.com/xvxiaoman8/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/xvxiaoman8/gomall/app/frontend --idl ../../idl/frontend/order_page.proto

.PHONY: gen-user
gen-user: 
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/xvxiaoman8/gomall/app/user --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module github.com/xvxiaoman8/gomall/app/product --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module github.com/xvxiaoman8/gomall/app/cart --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --service payment --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module github.com/xvxiaoman8/gomall/app/payment --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --service checkout --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module github.com/xvxiaoman8/gomall/app/checkout --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/xvxiaoman8/gomall/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/xvxiaoman8/gomall/app/order --pass "-use github.com/xvxiaoman8/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto


.PHONY: gen-auth
gen-auth:
	@cd  rpc_gen && cwgo client --type RPC --service auth --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/auth.proto
	@cd  app/auth && cwgo server --type RPC --service auth --module ${ROOT_MOD}/app/auth --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/auth.proto