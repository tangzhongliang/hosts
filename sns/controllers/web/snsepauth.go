package web

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"regexp"
	"sns/common/snsglobal"
	"sns/common/snsstruct"
	"sns/controllers/snscommon"
	"sns/controllers/snsep"
	"sns/controllers/snsinterface/snseper"
	"sns/controllers/snsplugin"
	"sns/models"
	"sns/util/snserror"
	"sns/util/snslog"
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
func (this *SnsEpController) Active() {
	activeId := this.GetString(":activeid")
	email, ok := snsglobal.SEmailAuthIdSyncMap.Get(activeId)
	if ok {
		//	--------------------------email active success;save email to db
		snslog.Df("(this *SnsEpController) Active()/ activeId%s;email%s", activeId, email)
		snsEpAccountEmail := email.(models.SnsEpAccountEmail)
		models.InsertOrUpdate(&snsEpAccountEmail)
		this.Redirect("/", 302)
	} else {
		//	--------------------------email active failed;redirect to bind email
		snslog.Df("(this *SnsEpController) Active()/ %s", activeId)
		this.Redirect("/pages/user/bindemail", 302)
	}
}

func (this *SnsEpController) Login() {
	snstype := GetLastString(this.Ctx.Request.RequestURI)
	snsEpAuther := snsglobal.SBeanFactory.New(snstype).(snseper.SNSEPAccounAuther)
	epAccount, _ := snsEpAuther.SnsCheckLoginResponse(&this.Controller)

	//	-------------------------------send active email when email is not null
	email := this.GetString("state")
	formater := regexp.MustCompile("[\\w]+@[\\w.]+")
	if len(formater.FindIndex([]byte(email))) > 0 {
		snslog.Df("(this *SnsEpController) Login()/ account%s;email%s", epAccount, email)
		snscommon.SendActiveEmail(models.SnsEpAccountEmail{AccountEPType: epAccount.AccountType, AccountId: epAccount.AccountId, Email: email})
	}
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
	epToPluginMessage, err := snsplugin.ParseToPluginMessage(epToPluginMessage)
	snserror.LogAndPanic(err)
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
