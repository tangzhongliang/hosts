package sns

import (
	"sns/common/snsstruct"
	"sns/
)

func SendMessageToPluginByPost(url, json string) bool {
	return true
}

func DispatchMessageToEP(msg snsstruct.EPMessage, token string) {

}

func DispatchMessageToPlugin(msg snsstruct.EPMessage, epUser SNSEPUser) {

}

func ParsePluginMessage(json string) snsstruct.EPMessage {

}
