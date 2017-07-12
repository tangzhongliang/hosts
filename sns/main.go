package main

import (
	_ "sns/common/snsglobal"
	_ "sns/models"
	_ "sns/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
