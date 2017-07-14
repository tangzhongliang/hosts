package snscommon

import (
	"sns/common/snsstruct"
	"sns/controllers/snsep"
	// "sns/controllers/snsplugin"
	"net/http"
	"sns/models"
	"sns/util/snserror"
	"sns/util/snslog"
	"sync"
)

var (
	ConversationSateMap     = make(map[models.SnsEpAccount]models.SnsPlugin)
	ConversationSateMapLock = &sync.RWMutex{}
)

func GetConversationPlugin(account models.SnsEpAccount) (plugin models.SnsPlugin, ok bool) {
	ConversationSateMapLock.RLock()
	defer ConversationSateMapLock.RUnlock()
	plugin, ok = ConversationSateMap[models.SnsEpAccount{AccountId: account.AccountId, AccountType: account.AccountType}]
	return
}

func SetConversationPlugin(account models.SnsEpAccount, plugin models.SnsPlugin) (plugin models.SnsPlugin, ok bool) {
	ConversationSateMapLock.Lock()
	defer ConversationSateMapLock.Unlock()
	ConversationSateMap[models.SnsEpAccount{AccountId: account.AccountId, AccountType: account.AccountType}] = plugin
	return
}

func SendMessageToPluginByPost(url, json string) (err error) {
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest("POST", url, []byte(json))
	req.Header.Add("Content-Type", "application/json")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		snserror.LogAndPanic(err)
	}
	return
}

func SendMessageToEp(accounts []models.SnsEpAccount) bool {
	snslog.If("SendMessageToEp/ send to %d;%+v", len(accounts), accounts)
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
	if !msg.Message.IsToAll && len(msg.TargetEmails.TargetUserEmail) == 0 && len(msg.TargetUsers) == 0 {
		ret = CreateServiceMessageResponse(1003, "no user to send")
		return
	}

	// -----------------------------send msg to ep users
	accounts := msg.TargetUsers
	if !msg.Message.IsToAll {
		emailaccounts := snsep.GetSnsEpByEmail(msg.TargetEmails.TargetUserEmail, msg.TargetEmails.Platforms)
		//	------------------------- Check SNSEpAuth For Plugin
		// CheckSNSEpAuthForPlugin
		snslog.I("DispatchMessageToEP/ email accounts", emailaccounts)
		accounts = append(accounts, emailaccounts...)
		snslog.I("DispatchMessageToEP/ IsToAll true all users", accounts)
	} else {
		accounts = snsep.GetSnsEpByPluginId(msg.PluginId)
		snslog.I("DispatchMessageToEP/ IsToAll false all users", accounts)
	}
	SendMessageToEp(accounts)
	if len(accounts) == 0 {
		return CreateServiceMessageResponse(1000, "accounts null")
	}
	return CreateServiceMessageResponse(1000, "")
}

func DispatchMessageToPlugin(msg snsstruct.EpToPluginMessage) {
	var plugin models.SnsPlugin
	err := models.Query(&plugin, &models.SnsPlugin{PluginId: msg.PluginId})
	snserror.LogAndPanic(err)
	SendMessageToPluginByPost(plugin.PluginWebhookUrl, msg)
}

func CreateServiceMessageResponse(code int, errMsg string) snsstruct.ServiceMessageResponse {
	if code == 1000 {
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
