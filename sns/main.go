package main

import (
	_ "sns/common/snsglobal"
	_ "sns/routers"

	_ "github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
