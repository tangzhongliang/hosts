package snsstruct

import (
	"sns/models"
)

type ServiceMessage struct {
}
type Attachment struct {
	CallbackId     string   `json:"callback_id"`
	AttachmentType string   `json:"attachment_type"`
	Actions        []Action `json:"actions"`
	ResponseUrl    string   `json:"response_url"`
	Text           string   `json:"text"`
}
type Action struct {
	Name    string       `json:"name"`
	Value   string       `json:"value"`
	Type    string       `json:"type"`
	Text    string       `json:"text"`
	Options []Option4Btn `json:"options"`
}
type ActionResponse struct {
	Name            string               `json:"name"`
	Value           string               `json:"value"`
	Type            string               `json:"type"`
	Text            string               `json:"text"`
	Options         []Option4Btn         `json:"options"`
	SelectedOptions []SelectedOption4Btn `json:"selected_options"`
}
type SnsEpResponse struct {
	CallbackId string           `json:"callback_id"`
	Actions    []ActionResponse `json:"actions"`
}
type Option4Btn struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}
type SelectedOption4Btn struct {
	Value string `json:"value"`
}

type PluginToEpMessageData struct {
	Text        string
	Link        string
	File        string
	IsToAll     bool
	Attachments []*Attachment
}
type PluginToEpMessageEmail struct {
	TargetUserEmail []string
	Platforms       []string
}
type PluginToEpMessage struct {
	TargetUserIds []string
	TargetUsers   []models.SnsEpAccount
	TargetEmails  PluginToEpMessageEmail
	PluginId      string
	Message       PluginToEpMessageData
}
type EpToPluginMessageData struct {
	ChannelId     string
	Text          string
	MessageTs     string
	MessageType   string
	File          string
	SnsEpResponse SnsEpResponse
}
type EpToPluginMessage struct {
	UserId   string
	User     models.SnsEpAccount `json:"-"`
	PluginId string
	Message  EpToPluginMessageData
}
type ServiceMessageResponse struct {
	ErrDefine
	Ok                  bool   `json:"ok"`
	Context             string `json:"context"`
	PluginSendMessageId int
}

type PluginTokenResponse struct {
	ErrDefine
	AccessToken string
}
