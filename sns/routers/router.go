package routers

import (
	"sns/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.NewNamespace("/index", beego.NSCond(func(ctx *context.Context) bool {
		//		enable domain check
		//		if ctx.Input.Domain() == "api.beego.me" {
		//			return true
		//		}
		return true
	}),
		beego.NSRouter("/plugin-list", &MainController{}, "get:Get"),
		beego.NSRouter("/epuser/plugin-list", &MainController{}, "get:Get")
		beego.NSRouter("/epuser/login", &MainController{}, "get:Get")
	)
	beego.NewNamespace("/dev", beego.NSCond(func(ctx *context.Context) bool {
		//		enable domain check
		//		if ctx.Input.Domain() == "api.beego.me" {
		//			return true
		//		}
		return true
	}),
		beego.NSRouter("/1.0/reference", &MainController{}, "get:Get")
	)
	beego.NewNamespace("/console", beego.NSCond(func(ctx *context.Context) bool {
		//		enable domain check
		//		if ctx.Input.Domain() == "api.beego.me" {
		//			return true
		//		}
		return true
	}),
		beego.NSRouter("/1.0/reference", &MainController{}, "get:Get")
	)
}
