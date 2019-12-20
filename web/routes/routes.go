package routes

import (
	"github.com/kataras/iris/_examples/mvc/login/web/middleware"
	"github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
	"github.com/kataras/iris/mvc"
	"superStar/services"
	"superStar/web/controllers"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()
	//映射路径
	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))
	//映射路径
	admin := mvc.New(b.Party("/admin"))
	//注册一个中间件
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))
}
