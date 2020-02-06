package routes

import (
	"../../bootstrap"
	"../../services"
	"../../web/controllers"
	"../middleware"
	"github.com/kataras/iris/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewUserService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.IndexController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
