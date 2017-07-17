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
		// -----------------------sns ep auth
		beego.NSRouter("/ep/login/ep_line", &web.SnsEpController{}, "get:Login"),
		beego.NSRouter("/ep/login/ep_slack", &web.SnsEpController{}, "get:Login"),
		beego.NSRouter("/ep/login/ep_office", &web.SnsEpController{}, "get:Login"),
		beego.NSRouter("/ep/auth/ep_line", &web.SnsEpController{}, "get:Auth"),
		beego.NSRouter("/ep/auth/ep_slack", &web.SnsEpController{}, "get:Auth"),
		beego.NSRouter("/ep/auth/ep_office", &web.SnsEpController{}, "get:Auth"),

		// -----------------------sns ep webhook
		beego.NSRouter("/ep/notify/ep_line", &web.SnsEpController{}, "post:Notify"),
		beego.NSRouter("/ep/notify/ep_slack", &web.SnsEpController{}, "post:Notify"),
		beego.NSRouter("/ep/notify/ep_office", &web.SnsEpController{}, "post:Notify"),
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
