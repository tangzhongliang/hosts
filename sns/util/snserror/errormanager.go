package snserror

import (
	"sns/util/snslog"
)

func LogAndPanic(err error) {
	if err != nil {
		snslog.E(err)
		panic(err)
	}

}
