package beans

//EPUserRole:user group team
// related SNSEPAccountEmail one-one
type SnsEpAccount struct {
	Name         string `xorm:"notnull 'name'"`
	Email        string `xorm:"'email'"`
	Id           string `xorm:"pk notnull 'id'"` //create id when account is team
	EPType       string `xorm:"pk notnull 'ep_type'"`
	AccountType  string `xorm:"notnull default 'user' 'account_type'"`
	ForeverToken string `xorm:"varchar(2000)"` //check is granted account which is granted
}

type SnsEpAccountEmail struct {
	AccountId     string `xorm:"pk notnull 'account_id'"`
	AccountEPType string `xorm:"pk notnull 'account_ep_type'"`
	Email         string
}
