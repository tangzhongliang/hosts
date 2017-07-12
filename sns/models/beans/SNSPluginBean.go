package beans

type SnsPluginAccount struct {
	Name          string `xorm:"pk notnull 'name'"`
	Email         string `xorm:"'email'"`
	AccountId     string `xorm:"unique notnull 'account_id'"`
	AccountSecret string `xorm:"notnull 'account_secret'"`
	Password      string `xorm:"notnull 'password'"`
}
type SnsPlugin struct {
	Id               int64
	AccountName      string `xorm:"unique notnull 'account_name'"`
	PluginName       string `xorm:"notnull 'plugin_name'"`
	PluginSecret     string `xorm:"notnull 'plugin_secret'"`
	PluginWebhookUrl string
	PluginButtonUrl  string
	PluginIcon       string
	PluginIconName   string
}
type SnsPluginConfig struct {
	Id          int64
	AccountName string `xorm:"unique notnull 'account_name'"`
}
