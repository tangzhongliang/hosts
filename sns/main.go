package main

import (
	_ "sns/common/snsglobal"
	"sns/models"
	_ "sns/routers"
	"sns/util/snserror"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true
	InitDbData()
	beego.Run()
}
func InitDbData() {
	var err error

	//	---------------------------insert ep account
	err = models.InsertOrUpdate(&models.SnsEpAccount{AccountId: "snstest111", EPType: "slack"})
	snserror.LogAndPanic(err)
	err = models.InsertOrUpdate(&models.SnsEpAccount{AccountId: "snstest222", EPType: "slack"})
	snserror.LogAndPanic(err)
	err = models.InsertOrUpdate(&models.SnsEpAccount{AccountId: "snstest333", EPType: "line"})
	snserror.LogAndPanic(err)

	//	--------------------------insert account email combine
	err = models.InsertOrUpdate(&models.SnsEpAccountEmail{AccountEPType: "slack", AccountId: "snstest111", Email: "1@qq.com"})
	snserror.LogAndPanic(err)
	err = models.InsertOrUpdate(&models.SnsEpAccountEmail{AccountEPType: "line", AccountId: "snstest333", Email: "1@qq.com"})
	snserror.LogAndPanic(err)

	// ---------------------------insert plugin account
	pluginAccount := models.SnsPluginAccount{Name: "plugin1", Password: "pwd", AccountId: "bbb", AccountSecret: "ccc"}
	err = models.InsertOrUpdate(&pluginAccount)
	snserror.LogAndPanic(err)
	plugin := models.SnsPlugin{AccountName: pluginAccount.Name, PluginId: "bbbplugin1", PluginWebhookUrl: "PluginWebhookUrl", PluginButtonUrl: "PluginButtonUrl", PluginToken: "token"}
	err = models.InsertOrUpdate(&plugin)
	snserror.LogAndPanic(err)

	// ----------------------------insert snsPluginEpAccounts
	snsPluginEpAccount := models.SnsPluginEpAccount{PluginId: "bbbplugin1", EpAccountId: "snstest111", EpAccountType: "slack"}
	err = models.InsertOrUpdate(&snsPluginEpAccount)
	snserror.LogAndPanic(err)

	emailAccount := models.EmailAccount{Email: "tangzhongliang@rst.ricoh.com", Password: "123"}
	models.InsertOrUpdate(&emailAccount)
}
