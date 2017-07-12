package snsglobal

import (
	"sns/common/snsfactory"
	// "sns/
	//	"sns/util/snserror"
)

var (
	SBeanFactory = snsfactory.RegisterStructMaps{}
	// SDBEngine    = New()
)

func init() {
	SBeanFactory.Register("snsfactory.Test", &snsfactory.Test{})
}
