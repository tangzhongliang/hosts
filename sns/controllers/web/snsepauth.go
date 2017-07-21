package web

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"sns/common/snsglobal"
	"sns/common/snsstruct"
	"sns/controllers/snscommon"
	"sns/controllers/snsep"
	"sns/controllers/snsinterface/snseper"
	// "sns/util/snslog"
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

func (this *SnsEpController) GetEpCheckUrl() {
	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	ret := make(map[string]interface{})
	var data snsstruct.SnsEpEmailBindRequest
	err := json.Unmarshal(body, &data)

	if err != nil {
		ret["ok"] = false
	} else {
		ret["ok"] = true
		ret["url"] = snsep.GetEpCheckUrl(data)
	}

	this.Data["json"] = ret
	this.ServeJSON()
}
