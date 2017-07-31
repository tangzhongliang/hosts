package snsep

import (
	"fmt"
	"math/rand"

	"persian/cloud/models"
	"persian/cloud/util/pConst"
	"persian/cloud/util/pFunction"
	"persian/cloud/util/pGlobal"
	"persian/cloud/util/pLog"
	"persian/cloud/util/pStruct"
	"strconv"

	"sync"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

type Line struct {
}

//var AutoLine Line
func GetLineBotInstance() (botInstance *linebot.Client, err error) {
	if bot == nil {
		err = InitLine()
	}
	return bot, err
}

func InitLine() (err error) {
	pGlobal.LineCallbackId = pStruct.LineMutexMap{
		Lock: new(sync.RWMutex),
		Bm:   make(map[string]string),
	}
	bot, err = linebot.New(
		beego.AppConfig.String("linebotsecret"),
		beego.AppConfig.String("linebotAccessToken"),
	)
	return err
}

//发送带button的Message
func Send(userID string, eventInfo models.Event_info, timeout int64) (err error) {
	pLog.D("*********Send---eventInfo************", eventInfo)
	//check bot
	if bot == nil {
		pLog.Ef("Line bot nil err ")
		return err
	}
	accountinfo, err1 := models.AccountInfo.GetAccountInfo(eventInfo.Creatormail)
	if err1 != nil {
		pLog.Ef("models.AccountInfo.GetAccountInfo err %v ", err1)
		return err1
	}

	//get Start time
	tStartTime, err := pFunction.GetLocalFromUTC(eventInfo.Startdatetime, accountinfo[0].Timezone)
	if err != nil {
		pLog.Ef("CheckTime err %v ", err)
		return err
	}

	//预约信息
	ConfirmMessage := pFunction.GetSlackMessage(pConst.ACT_SNS_CONFIM_MESSAGE, accountinfo[0].Language, eventInfo.Creatormail, tStartTime, eventInfo.Location, eventInfo.Summary, timeout, "")
	//message := eventInfo.Creatormail + " さん\n" + tStartTime + " " + eventInfo.Location + "にて、" + "会議（" + eventInfo.Summary + "）" + "が予約されています。\n"
	//message += "予約通り利用する場合は”利用する”、利用しない場合は”キャンセルする”を選択してください。\n"
	//message += "10分間返答がない場合、自動的に予約はキャンセルされます。"

	// create callback id
	callbackId := strconv.Itoa(RandInt(10000000, 99999999))
	//CallbackIdMap[callbackId] = eventInfo.Resourceid + "&" + eventInfo.Eventid
	isBool := pGlobal.LineCallbackId.Set(callbackId, eventInfo.Resourceid+"&"+eventInfo.Eventid)
	if isBool != true {
	}
	// set button information

	lang := accountinfo[0].Language
	if lang == "" {
		lang = "ja-JP"
	}
	pLog.D("***************i18n.Tr(accountinfo[0].Language***************", i18n.Tr(lang, "lineSettings.lineRemiderYesButtonMsg"))
	pLog.D("***************i18n.Tr(accountinfo[0].Language2***************", i18n.Tr(lang, "lineSettings.lineRemiderNoButtonMsg"))
	yesBtn := linebot.NewPostbackTemplateAction(i18n.Tr(lang, "lineSettings.lineRemiderYesButtonMsg"), "{\"value\":\"yes\",\"callback_id\":\""+callbackId+"\"}", i18n.Tr(lang, "lineSettings.lineRemiderYesButtonMsg"))
	noBtn := linebot.NewPostbackTemplateAction(i18n.Tr(lang, "lineSettings.lineRemiderNoButtonMsg"), "{\"value\":\"no\",\"callback_id\":\""+callbackId+"\"}", i18n.Tr(lang, "lineSettings.lineRemiderNoButtonMsg"))

	template := linebot.NewConfirmTemplate(ConfirmMessage, yesBtn, noBtn)
	message := linebot.NewTemplateMessage(i18n.Tr(lang, "lineSettings.lineRemiderTitleMsg"), template)
	if _, err := bot.PushMessage(userID, message).Do(); err != nil {
		pLog.Ef("Line PushMessage err %v ", err)
		return err
	}
	return err
}

//发送带button的checkInMessage
func SendCheckInMessage(userID string, eventInfo models.Event_info, urlMessage string, timeout int64) (err error) {
	//check bot
	if bot == nil {
		pLog.Ef("Line bot nil err ")
		return err
	}

	accountinfo, err1 := models.AccountInfo.GetAccountInfo(eventInfo.Creatormail)
	if err1 != nil || len(accountinfo) == 0 {
		pLog.Ef("models.AccountInfo.GetAccountInfo err %v ", err1)
		return err1
	}

	tStartTime, err := pFunction.GetLocalFromUTC(eventInfo.Startdatetime, accountinfo[0].Timezone)
	if err != nil {
		pLog.Ef("CheckTime err %v ", err)
		return err
	}
	lang := accountinfo[0].Language
	if lang == "" {
		lang = "ja-JP"
	}
	// send 预约信息
	ChickInMessage := pFunction.GetSlackMessage(pConst.ACT_SNS_CIN_MESSAGE, accountinfo[0].Language, eventInfo.Creatormail, tStartTime, eventInfo.Location, eventInfo.Summary, timeout, urlMessage)
	SendMessageOnly(userID, ChickInMessage)

	// create callback id
	callbackId := strconv.Itoa(RandInt(10000000, 99999999))
	//CallbackIdMap[callbackId] = eventInfo.Resourceid + "&" + eventInfo.Eventid
	isBool := pGlobal.LineCallbackId.Set(callbackId, eventInfo.Resourceid+"&"+eventInfo.Eventid)
	if isBool != true {
	}
	// set button information
	cancelBtn := linebot.NewPostbackTemplateAction(i18n.Tr(lang, "lineSettings.lineCheckInCancelButtonMsg"), "{\"value\":\"cancel\",\"callback_id\":\""+callbackId+"\"}", i18n.Tr(lang, "lineSettings.lineCheckInCancelButtonMsg"))

	ConfirmMessage := i18n.Tr(lang, "lineSettings.lineCheckInCancelMsg")

	template := linebot.NewButtonsTemplate("", "", ConfirmMessage, cancelBtn)
	message := linebot.NewTemplateMessage(i18n.Tr(lang, "lineSettings.lineCheckInTitleMsg"), template)
	if _, err := bot.PushMessage(userID, message).Do(); err != nil {
		pLog.Ef("Line CheckInMessage err %v ", err)
		return err
	}
	return err
}

//只发送message
func SendMessageOnly(userID string, text string) (err error) {
	//check bot
	if bot == nil {
		pLog.Ef("Line bot nil err ")
		return err
	}

	if _, err := bot.PushMessage(userID, linebot.NewTextMessage(text)).Do(); err != nil {
		pLog.Ef("Line SendMessageOnly err %v ", err)
		return err
	}
	return err
}

//发送取消会议message
func SendCancelMessage(userID string, eventInfo models.Event_info) (err error) {
	//check bot
	if bot == nil {
		pLog.Ef("Line bot nil err ")
		return err
	}

	accountinfo, err1 := models.AccountInfo.GetAccountInfo(eventInfo.Creatormail)
	if err1 != nil || len(accountinfo) == 0 {
		pLog.Ef("models.AccountInfo.GetAccountInfo err %v ", err1)
		return err1
	}

	tStartTime, err := pFunction.GetLocalFromUTC(eventInfo.Startdatetime, accountinfo[0].Timezone)
	if err != nil {
		pLog.Ef("CheckTime err %v ", err)
		return err
	}

	message := pFunction.GetSlackMessage(pConst.ACT_SNS_CONFIM_CANCEL, accountinfo[0].Language, eventInfo.Creatormail, tStartTime, eventInfo.Location, eventInfo.Summary, 0, "")
	//message := eventInfo.Creatormail + " さん\n" + "予約した会議室「(" + eventInfo.Summary + ")" + tStartTime + " " + eventInfo.Location + "」はキャンセルされました。"
	if _, err := bot.PushMessage(userID, linebot.NewTextMessage(message)).Do(); err != nil {
		pLog.Ef("Line SendCancelMessage err %v ", err)
		return err
	}
	return err
}

//SendInitMessage ...初期設定を行うリンクを送信する。（友達追加時に送信されるべきメッセージ）
func SendInitMessage(userID string, settingPageURL string) (err error) {

	url := settingPageURL + userID
	urlAction := linebot.NewURITemplateAction(i18n.Tr("en-US", "lineSettings.lineInitSettingButtonMsg"), url)
	template := linebot.NewButtonsTemplate("", "", i18n.Tr("en-US", "lineSettings.lineInitSettingMsg"), urlAction)
	msg := linebot.NewTemplateMessage(i18n.Tr("en-US", "lineSettings.lineInitSettingTitleMsg"), template)
	if _, err := bot.PushMessage(userID, msg).Do(); err != nil {
		pLog.Ef("Line SendInitMessage err %v ", err)
		return err
	}
	return err
}

func GetUserDisplayName(UserID string) string {
	userPrefile, err := bot.GetProfile(UserID).Do()
	if err != nil {
		return ""
	}
	return userPrefile.DisplayName
}
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func (this Line) LineMessageButtonCallback(c beego.Controller) {
	// get Post callback data
	fmt.Println("LineMessageButtonCallback")
	bot, err := GetLineBotInstance()
	events, err := bot.ParseRequest(c.Ctx.Request)
	if err != nil {
		pLog.Ef("LineMessageButtonCallback ParseRequest err: %v ", err)
		if err == linebot.ErrInvalidSignature {
			c.Ctx.ResponseWriter.WriteHeader(400)
		} else {
			c.Ctx.ResponseWriter.WriteHeader(500)
		}
		return
	}
	c.Ctx.ResponseWriter.WriteHeader(200)
	fmt.Printf("events%s\n", events)
	// for _, event := range events {
	// 	if event.Source.UserID == "" &&
	// 		event.Source.GroupID == "" &&
	// 		event.Source.RoomID == "" {
	// 		pLog.If("UserId or GroupId is empty")
	// 		continue
	// 	} else {
	// 		pLog.If("find userId:", event.Source.UserID, "groupID:", event.Source.GroupID, "RoomID:", event.Source.RoomID)
	// 	}
	// 	if event.Type == linebot.EventTypePostback { //ユーザーからのPostBackEvent
	// 		// data := event.Postback.Data
	// 		// pLog.D("***************************event.Postback.Data************", event.Postback.Data)
	// 		// go ParsePostbackData(event.Source.UserID, data)
	// 	} else if event.Type == linebot.EventTypeFollow { //FollowEvent（Botが友達追加された）
	// 		meettingId := pStruct.MeettingId{"", ""}
	// 		eventInfo := models.Event_info{Resourceid: "", Eventid: ""}
	// 		settingURL := "https://" + beego.AppConfig.String("domain") + "/lineBind" + "?lineUserId="
	// 		Update_LineMessage := pStruct.Update_LineMessage{event.Source.UserID, settingURL}
	// 		inputData := pStruct.MeettingStateInfo{pConst.ACT_SNS_ADD_LINE_USER, eventInfo, Update_LineMessage, 0}
	// 		// TODO(Lineへ送信)
	// 		pThread.ProcManager.Execute(pConst.JOB_TYPE_SNS, meettingId, inputData)
	// 	} else if event.Type == linebot.EventTypeJoin { //JoinEvent(botがGroupに参加)
	// 		meettingId := pStruct.MeettingId{"", ""}
	// 		eventInfo := models.Event_info{Resourceid: "", Eventid: ""}
	// 		settingURL := "https://" + beego.AppConfig.String("domain") + "/lineBind" + "?GroupId="
	// 		userId := event.Source.GroupID
	// 		if userId == "" {
	// 			userId = event.Source.RoomID
	// 		}
	// 		Update_LineMessage := pStruct.Update_LineMessage{event.Source.GroupID, settingURL}
	// 		inputData := pStruct.MeettingStateInfo{pConst.ACT_SNS_ADD_LINE_USER, eventInfo, Update_LineMessage, 0}
	// 		// TODO(Lineへ送信)
	// 		pThread.ProcManager.Execute(pConst.JOB_TYPE_SNS, meettingId, inputData)
	// 	} else if event.Type == linebot.EventTypeMessage { //Message Event
	// 		switch message := event.Message.(type) {
	// 		case *linebot.TextMessage:
	// 			if message.Text == "persian-chat" {
	// 				//send bind message
	// 				meettingId := pStruct.MeettingId{"", ""}
	// 				eventInfo := models.Event_info{Resourceid: "", Eventid: ""}
	// 				var userId string
	// 				var settingURL string
	// 				if event.Source.UserID != "" && event.Source.Type == linebot.EventSourceTypeUser {
	// 					userId = event.Source.UserID
	// 					settingURL = "https://" + beego.AppConfig.String("domain") + "/lineBind" + "?lineUserId="
	// 				} else if event.Source.GroupID != "" && event.Source.Type == linebot.EventSourceTypeGroup {
	// 					userId = event.Source.GroupID
	// 					settingURL = "https://" + beego.AppConfig.String("domain") + "/lineBind" + "?GroupId="
	// 				} else if event.Source.RoomID != "" && event.Source.Type == linebot.EventSourceTypeRoom {
	// 					userId = event.Source.RoomID
	// 					settingURL = "https://" + beego.AppConfig.String("domain") + "/lineBind" + "?GroupId="
	// 				} else {
	// 					continue
	// 				}
	// 				Update_LineMessage := pStruct.Update_LineMessage{userId, settingURL}
	// 				inputData := pStruct.MeettingStateInfo{pConst.ACT_SNS_ADD_LINE_USER, eventInfo, Update_LineMessage, 0}
	// 				// TODO(Lineへ送信)
	// 				pThread.ProcManager.Execute(pConst.JOB_TYPE_SNS, meettingId, inputData)
	// 			}
	// 		}
	// 	}
	// }
}
