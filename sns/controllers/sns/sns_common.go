package sns

import (
	"sns/common/snsstruct"
	"sns/controllers/snsep"
	"sns/controllers/snsplugin"
	"sns/models"
	"sns/util/snserror"
	"sns/util/snslog"
)

func SendMessageToPluginByPost(url, json string) bool {
	return true
}

func SendMessageToEp(accounts []models.SnsEpAccount) bool {
	snslog.If("SendMessageToEp/ send to %+v", accounts)
	return true
}

func DispatchMessageToEP(msg snsstruct.PluginToEpMessage, token string) (ret snsstruct.ServiceMessageResponse) {
	// -------------------------check Plugin id
	var plugin models.SnsPlugin
	err := models.QueryByKey(&plugin, &models.SnsPlugin{PluginId: msg.PluginId})
	if err != nil {
		ret = CreateServiceMessageResponse(1001, "plugin id is invalid")
		return
	}

	//	--------------------------check token
	if plugin.PluginToken != token {
		ret = CreateServiceMessageResponse(1002, "token is invalid")
		return
	}

	//	---------------------------check user null
	if !msg.Message.IsToAll && len(msg.TargetUserEmails) && len(msg.TargetUsers) == 0 {
		ret = CreateServiceMessageResponse(1003, "no user to send")
		return
	}

	// -----------------------------send msg to ep users
	var accounts []models.SnsEpAccount
	if !msg.Message.IsToAll {
		accounts = snsep.GetSnsEpByEmail(msg.TargetEmails.TargetUserEmail, msg.TargetEmails.Platforms)
		//	------------------------- Check SNSEpAuth For Plugin
		// CheckSNSEpAuthForPlugin
		accounts = append(accounts, msg.TargetUsers)
	} else {
		accounts = snsep.GetSnsEpByPluginId(msg.PluginId)
	}
	SendMessageToEp(accounts)
	return CreateServiceMessageResponse(1000, "")
}

func DispatchMessageToPlugin(msg snsstruct.EPMessage, epUser SNSEPUser) {

}

func ParsePluginMessage(json string) snsstruct.EPMessage {

}
func CreateServiceMessageResponse(code int, errMsg string) snsstruct.ServiceMessageResponse {
	if code != 1000 {
		return snsstruct.ServiceMessageResponse{
			Ok:         true,
			ErrMessage: errMsg,
			ErrCode:    code,
		}
	} else {
		return snsstruct.ServiceMessageResponse{
			Ok:         false,
			ErrMessage: errMsg,
			ErrCode:    code,
		}
	}
}
