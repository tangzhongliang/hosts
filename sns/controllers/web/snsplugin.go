package web

import (
	"github.com/astaxie/beego"
	"sns/common/snsglobal"
	"sns/common/snsstruct"
	"sns/controllers/snscommon"
	"sns/controllers/snsplugin"
	"strings"
)

type SnsPluginController struct {
	beego.Controller
}

func (this *SnsPluginController) PluginToEpMessage() {
	pluginToEpMessage, err := snsplugin.ParseFromPluginMessage(string(this.Ctx.Input.RequestBody))
	var res snsstruct.ServiceMessageResponse
	if err != nil {
		res = snsstruct.ServiceMessageResponse{ErrCode: 1010, ErrMessage: "message formatter is incorrect", Collect: string(this.Ctx.Input.RequestBody)}
	} else {
		res = snscommon.DispatchMessageToEP(msg, token)
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

}
