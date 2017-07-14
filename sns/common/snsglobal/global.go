package snsglobal

import (
	"sns/common/snsfactory"
	"sns/models"
	//	"sns/util/snserror"
)

var (
	SBeanFactory = snsfactory.RegisterStructMaps{}
	// SDBEngine    = New()
)

func init() {
	SBeanFactory.Register("snsfactory.Test", &snsfactory.Test{})
}
