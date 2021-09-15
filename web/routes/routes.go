package routes

import (
	"github.com/kataras/iris/mvc"

	"iris_server/bootstrap"
	"iris_server/services"
	"iris_server/web/controllers"
	"iris_server/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()
	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)

	// 用户相关
	UserService := services.NewUserService() //注册service
	user := mvc.New(b.Party("/user"))
	user.Register(UserService)
	user.Handle(new(controllers.UserController)) //注册控制器

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))
}
