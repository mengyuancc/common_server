package main

import (
	"bootstrap"
	"web/middleware/identity"
	"web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	// 初始化iris
	app := bootstrap.New("Server Api", "Mengyuan")
	app.Bootstrap()
	// 引入路由
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	// 设置端口 运行
	app.Listen(":8082")
}
