package models

import (
// "github.com/jinzhu/gorm"
)

type SnsPluginAccount struct {
	BaseModel
	Name          string `gorm:"primary_key"`
	Email         string
	AccountId     string `gorm:"unique;not null"`
	AccountSecret string `gorm:"not null"`
	Password      string `gorm:"not null"`
}
type SnsPlugin struct {
	BaseModelWithId
	AccountName      string `gorm:"not null"`
	PluginName       string `gorm:"not null"`
	PluginSecret     string `gorm:"not null"`
	PluginWebhookUrl string
	PluginButtonUrl  string
	PluginIcon       string
	PluginIconName   string
}
type SnsPluginConfig struct {
	BaseModelWithId
	AccountName string
}
