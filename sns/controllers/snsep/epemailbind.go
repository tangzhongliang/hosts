package snsep

import (
	"github.com/astaxie/beego"

	"sns/common"
	"sns/common/snsglobal"
	"sns/common/snsstruct"
	"sns/controllers/snscommon"
	"sns/controllers/snsinterface/snseper"
)

func GetEmailBindPage(this *beego.Controller) {

}

func GetEpCheckUrl(req snsstruct.SnsEpEmailBindRequest) (url string) {
	snsglobal.SEmailAuthIdSyncMap.Lock.Lock()

	res := snscommon.ExecUntilSuccess(func() (res interface{}, ok bool) {
		res := common.CreateRandomString(20)
		_, ok := snsglobal.SEmailAuthIdSyncMap.Get(random)
		ok = !ok
		return
	})
	snsglobal.SEmailAuthIdSyncMap.Set(res.(string), req)
	snsglobal.SEmailAuthIdSyncMap.Lock.Unlock()
	switch req.AccountType {
	case "line":
		return snsglobal.SBeanFactory.New("ep_line").(snseper.SNSEPAccounAuther).GetSnsCheckLoginUrl(req.Email)
	case "slack":
		return snsglobal.SBeanFactory.New("ep_slack").(snseper.SNSEPAccounAuther).GetSnsCheckLoginUrl(req.Email)
	case "teams":
		return snsglobal.SBeanFactory.New("ep_team").(snseper.SNSEPAccounAuther).GetSnsCheckLoginUrl(req.Email)
	}
}
