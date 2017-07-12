package snseper

type SnsEPMessageSender interface {
	//send message struct json
	SendAttachmentByChannel(token, message, channelId string)
	SendAttachmentByUser(token, message, userId string)

	//send file only
	//	SendFileByChannel(token, message, url, channelId string)
	//	SendFileByUser(token, message, url, userId string)
}
