package models

//EPUserRole:user group team
// related SNSEPAccountEmail one-one
type SnsEpAccount struct {
	BaseModel
	Name         string `gorm:"not null;index"`
	Email        string
	AccountId    string `gorm:"primary_key"` //create id when account is team
	EPType       string `gorm:"primary_key"`
	AccountType  string `gorm:"not null"`
	ForeverToken string `gorm:"type:varchar(2000)"` //check is granted account which is granted
}

type SnsEpAccountEmail struct {
	BaseModel
	AccountId     string `gorm:"primary_key"`
	AccountEPType string `gorm:"primary_key"`
	Email         string
}
