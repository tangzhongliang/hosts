package snsep

import (
	// "fmt"
	"sns/models"
	"sns/util/snserror"
	"strings"
)

func GetSnsEpByEmail(emails []string, types []string) (epAccounts []models.SnsEpAccount) {
	for _, email := range emails {

		var epAccountEmails []models.SnsEpAccountEmail
		// models.Query(&epAccountEmails, "email = ? and account_ep_type in (?)", email, types)
		models.GetDB().Where("email = ? and account_ep_type in (?)", email, types).Find(&epAccountEmails)

		for _, value := range epAccountEmails {
			epAccounts = append(epAccounts, models.SnsEpAccount{AccountId: value.AccountId, AccountType: value.AccountEPType})
		}
	}
	return
}

func GetSnsEpByPluginId(pluginId string) (accounts []models.SnsEpAccount) {
	//	---------------------------------find account and plugin
	var snsPluginEpAccounts []models.SnsPluginEpAccount
	err := models.Query(&snsPluginEpAccounts, &models.SnsPluginEpAccount{PluginId: pluginId})
	snserror.LogAndPanic(err)

	//	---------------------------------- pack snsPluginEpAccounts into account array
	for _, pluginEpAccount := range snsPluginEpAccounts {
		accounts = append(accounts, models.SnsEpAccount{AccountId: pluginEpAccount.EpAccountId, AccountType: pluginEpAccount.EpAccountType})
	}
	return
}
func GetSNSToken(string) string {
	return ""
}
