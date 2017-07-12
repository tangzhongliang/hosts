package dao

import (
	_ "sns/common/snsglobal"
	"sns/models/snsdao"
	"testing"
)

func TestDao(t *testing.T) {
	snsdao.Test()
}
func TestGorm(t *testing.T) {
	snsdao.TestGorm()
}
