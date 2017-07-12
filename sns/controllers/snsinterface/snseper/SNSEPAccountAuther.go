package snsepmessagecenter

import (
	"sns/models/beans"

	"github.com/astaxie/beego"
)

type SNSEPAccounAuther interface {
	//send message struct json

	GetSnsCheckLoginUrl(emailEncode string) string
	SnsCheckLoginResponse(controller *beego.Controller) bool

	GetAuthUrl()
	SnsCheckAuthResponse(controller *beego.Controller) bool

	GetSnsByEmail(email) []beans.SNSEPUser
	GetSNSToken(string)
	//send file only
	//	SendFileByChannel(token, message, url, channelId string)
	//	SendFileByUser(token, message, url, userId string)
}
