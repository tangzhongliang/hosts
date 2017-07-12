package routers

import (
	"sns/controllers/web"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &web.MainController{})
	beego.NewNamespace("/index",
		beego.NSCond(func(ctx *context.Context) bool {
			//		enable domain check
			//		if ctx.Input.Domain() == "api.beego.me" {
			//			return true
			//		}
			return true
		}),
		beego.NSRouter("/plugin-list", &web.MainController{}, "get:Get"),
		beego.NSRouter("/epuser/plugin-list", &web.MainController{}, "get:Get"),
		beego.NSRouter("/epuser/login", &web.MainController{}, "get:Get"),
	)
	beego.NewNamespace("/dev", beego.NSCond(func(ctx *context.Context) bool {
		//		enable domain check
		//		if ctx.Input.Domain() == "api.beego.me" {
		//			return true
		//		}
		return true
	}),
		beego.NSRouter("/1.0/reference", &web.MainController{}, "get:Get"),
	)
	beego.NewNamespace("/console", beego.NSCond(func(ctx *context.Context) bool {
		//		enable domain check
		//		if ctx.Input.Domain() == "api.beego.me" {
		//			return true
		//		}
		return true
	}),
		beego.NSRouter("/1.0/reference", &web.MainController{}, "get:Get"),
	)
}
