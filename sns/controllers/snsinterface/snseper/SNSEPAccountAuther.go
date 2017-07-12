package snsepmessagecenter

import (
	"sns/

	"github.com/astaxie/beego"
)

type SNSEPAccounAuther interface {
	//send message struct json

	GetSnsCheckLoginUrl(emailEncode string) string
	SnsCheckLoginResponse(controller *beego.Controller) bool

	GetAuthUrl()
	SnsCheckAuthResponse(controller *beego.Controller) bool

	GetSnsByEmail(email) []SNSEPUser
	GetSNSToken(string)
	//send file only
	//	SendFileByChannel(token, message, url, channelId string)
	//	SendFileByUser(token, message, url, userId string)
}
