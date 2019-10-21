package main

import (
	_ "meal/routers"
	_ "meal/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
