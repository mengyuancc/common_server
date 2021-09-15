package main

import (
	"iris_server/bootstrap"
	"iris_server/web/middleware/identity"
	"iris_server/web/routes"
)

/**
 * 初始化操作
 */
func newApp() *bootstrap.Bootstrapper {
	//初始化iris
	app := bootstrap.New("Server Api", "Mengyuan")
	app.Bootstrap()
	//引入路由
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	// 设置端口 运行
	app.Listen(":8099")
}
