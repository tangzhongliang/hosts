package web

import (
	"github.com/astaxie/beego"
	"sns/common/snsglobal"
	"sns/controllers/snscommon"
	"sns/controllers/snsinterface/snseper"
	"strings"
)

type SnsEpController struct {
	beego.Controller
}

func GetLastString(uri string) (snstype string) {
	snstypeIndex := strings.LastIndex(uri, "/")
	snstype = uri[snstypeIndex : len(uri)+1]
	return
}

func (this *SnsEpController) Login() {
	snstype := GetLastString(this.Ctx.Request.RequestURI)
	snsEpAuther := snsglobal.SBeanFactory.New(snstype).(snseper.SNSEPAccounAuther)
	snsEpAuther.SnsCheckLoginResponse(&this.Controller)
}

func (this *SnsEpController) Auth() {
	snstype := GetLastString(this.Ctx.Request.RequestURI)
	snsEpAuther := snsglobal.SBeanFactory.New(snstype).(snseper.SNSEPAccounAuther)
	snsEpAuther.SnsCheckAuthResponse(&this.Controller)
}

func (this *SnsEpController) Notify() {
	snstype := GetLastString(this.Ctx.Request.RequestURI)
	snsEpSender := snsglobal.SBeanFactory.New(snstype).(snseper.SNSEPAccounAuther)
	epToPluginMessage := snsEpSender.ParseMessageFromWebhook(&this.Controller)
	snscommon.DispatchMessageToPlugin(epToPluginMessage)
}
