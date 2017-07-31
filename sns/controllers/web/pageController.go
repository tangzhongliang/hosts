package web

import (
	"github.com/astaxie/beego"
	// "sns/common/snsglobal"
	// "sns/controllers/snscommon"
	// "sns/controllers/snsinterface/snseper"
	// "strings"
)

type PageController struct {
	beego.Controller
}

func (this *PageController) UserBindEmail() {
	this.Data["accountType"] = this.GetString("accountType", "ep_line")
	this.Data["accountTypeText"] = "line"
	this.Data["userId"] = this.GetString("userId", "xxx@line.com")
	this.GetSession("bindAccountType")
	this.GetSession("bindAccountType")
	this.TplName = "user_bind_email.html"
}
