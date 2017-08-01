package web

type UserChekcController struct {
	BaseController
}

func (this *UserChekcController) Prepare() {
	this.BaseController.Prepare()
	emailAccount := this.GetSessionString("emailAccount")
	if len(emailAccount) == 0 && this.Ctx.Request.RequestURI != "/user/login" {
		this.Redirect("/user/login", 302)
	}
}
