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
	this.TplName = "user_bind_email.html"
}
