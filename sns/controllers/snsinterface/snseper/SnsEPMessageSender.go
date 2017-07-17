package snseper

import (
	"sns/common/snsstruct"
)

type SnsEPMessageSender interface {
	//send message struct json
	SendAttachmentByChannel(token string, msg snsstruct.PluginToEpMessage)
	SendAttachmentByUser(token string, msg snsstruct.PluginToEpMessage)

	//send file only
	//	SendFileByChannel(token, message, url, channelId string)
	//	SendFileByUser(token, message, url, userId string)
}
