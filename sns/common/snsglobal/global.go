package snsglobal

import (
	"sns/common/snsfactory"
	"sns/models/beans"
	"sns/models/snsdao"
	//	"sns/util/snserror"
)

var (
	SBeanFactory = snsfactory.RegisterStructMaps{}
	SDBEngine    = snsdao.New()
)

func init() {
	SBeanFactory.Register("snsfactory.Test", &snsfactory.Test{})
	SDBEngine.CreateTables(&beans.SNSEPAccount{})
}
