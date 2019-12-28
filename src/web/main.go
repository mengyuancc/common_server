package main

import (
	"bootstrap"
	"web/middleware/identity"
	"web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Server Api", "my")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
