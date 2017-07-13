package dao

import (
	"fmt"
	_ "sns/common/snsglobal"
	"sns/models"
	"sns/util/snserror"
	"sns/util/snslog"
	"testing"
	"time"
)

var err error

func TestNew(t *testing.T) {
	var account models.SnsEpAccount

	err = account.InsertOrUpdate(&models.SnsEpAccount{AccountId: "1111", EPType: "slack", Email: "asdfasdf"})
	snserror.LogAndPanic(err)
	err = account.InsertOrUpdate(&models.SnsEpAccount{AccountId: "2222", EPType: "slack2", Email: "asdfasdf"})
	snserror.LogAndPanic(err)
	err = account.InsertOrUpdate(&models.SnsEpAccount{AccountId: "3333", EPType: "slack", Email: "asdfasdf"})
	snserror.LogAndPanic(err)
	now := time.Now()
	for i := 0; i < 1000; i++ {
		err = account.InsertOrUpdate(&models.SnsEpAccount{AccountId: fmt.Sprintf("1111%d", i), EPType: "slack", Email: fmt.Sprintf("aaaaa%d", i)})
		snserror.LogAndPanic(err)
	}
	snslog.I(time.Now().Sub(now).Nanoseconds())
}

func TestFind(t *testing.T) {
	var account models.SnsEpAccount
	var accounts []models.SnsEpAccount
	var accountInfo models.SnsEpAccount
	err = account.Query(&accounts, &models.SnsEpAccount{AccountId: "1111", EPType: "slack"})
	snslog.I(accounts)
	snserror.LogAndPanic(err)
	err = account.QueryByKey(&accountInfo, &models.SnsEpAccount{AccountId: "1111", EPType: "slack", Email: "vbvvvv"})
	snserror.LogAndPanic(err)
	snslog.I(accounts)
}

func TestUpdate(t *testing.T) {
	var account models.SnsEpAccount

	//	for i := 0; i < 10000; i++ {
	err = account.InsertOrUpdate(&models.SnsEpAccount{AccountId: "1111", EPType: "slack", Email: "dfdfdfdfdf"})
	snserror.LogAndPanic(err)
	//	}

}

// func TestExist(t *testing.T) {
// 	var account models.SnsEpAccount
// 	ret := account.Exist(&models.SnsEpAccount{AccountId: "1111", EPType: "slack", Email: "asdfasdfasdfa"})
// 	if !ret {
// 		panic(nil)
// 	}

// }

func TestDelete(t *testing.T) {
	var account models.SnsEpAccount
	err = account.DeleteByStruct(&models.SnsEpAccount{AccountId: "3333", EPType: "slack"})
	snserror.LogAndPanic(err)
	// models.GetDB().Exec("delete from sns_ep_accounts where account_id = ?  and ep_type = ? ", "1111", "slack")
}
