package web

import (
// "github.com/astaxie/beego"
// "sns/common/snsglobal"
// "sns/controllers/snscommon"
// "sns/controllers/snsinterface/snseper"
// "strings"
)

type UserPageController struct {
	UserChekcController
}

func (this *UserPageController) UserBindEmail() {
	//set default
	this.Data["accountType"] = "ep_slack"
	this.Data["accountTypeText"] = "slack"
	this.Data["userId"] = ""
	//get session value
	SetValue(this, "accountType", this.GetSessionString("bindAccountType"))
	SetValue(this, "accountTypeText", "line")
	SetValue(this, "userId", this.GetSessionString("bindAccountId"))

	//use url param
	SetValue(this, "accountType", this.GetString("bindAccountType"))
	SetValue(this, "accountTypeText", "line")
	SetValue(this, "userId", this.GetString("bindAccountId"))

	this.TplName = "user_bind_email.html"
}
func SetValue(this *UserPageController, key, value string) {
	if value == "" {
		return
	}
	this.Data[key] = value
}
func (this *UserPageController) UserPage() {
	this.TplName = "user.html"
}
func (this *UserPageController) UserLoginPage() {
	this.TplName = "login.html"
}
