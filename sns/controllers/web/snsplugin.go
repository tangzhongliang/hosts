package web

import (
	"encoding/json"
	"io/ioutil"
	"sns/common/snsglobal"
	"sns/common/snsstruct"
	"sns/controllers/snscommon"
	"sns/controllers/snsplugin"
	"sns/models"

	"github.com/astaxie/beego"
	// "strings"
)

type SnsPluginController struct {
	beego.Controller
}

func (this *SnsPluginController) PluginToEpMessage() {
	pluginToEpMessage, err := snsplugin.ParseFromPluginMessage(string(this.Ctx.Input.RequestBody))
	var res snsstruct.ServiceMessageResponse
	if err != nil {
		res = snsstruct.ServiceMessageResponse{ErrDefine: snsglobal.SErrConfig.GetError(snsglobal.CErrCommon, "message_format_err"), Context: string(this.Ctx.Input.RequestBody)}
	} else {
		token, err2 := snscommon.GetBearerFromRequest(this.Ctx.Request)
		if err2 != nil {
			res = snsstruct.ServiceMessageResponse{ErrDefine: snsglobal.SErrConfig.GetError(snsglobal.CErrCommon, "message_format_err"), Context: string(this.Ctx.Input.RequestBody)}
		}
		res = snscommon.DispatchMessageToEP(pluginToEpMessage, token)
	}
	this.Data["json"] = &res
	this.ServeJSON()
}

func (this *SnsPluginController) RegisterPluginAccount() {

}

func (this *SnsPluginController) LoginPluginAccount() {

}

func (this *SnsPluginController) AddAndEditPlugin() {

}

func (this *SnsPluginController) RequestPluginToken() {
	body, _ := ioutil.ReadAll(this.Ctx.Request.Body)
	var plugin models.SnsPlugin
	err := json.Unmarshal([]byte(body), &plugin)
	//	---------------------check message format
	var res snsstruct.PluginTokenResponse
	if err != nil {
		res = snsstruct.PluginTokenResponse{
			ErrDefine: snsglobal.SErrConfig.GetError(snsglobal.CErrCommon, "message_format_err"),
		}
	} else {
		res = snsplugin.RequestPluginToken(plugin)
	}
	this.Data["json"] = res
	this.ServeJSON()
}
