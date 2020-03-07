package main

import (
	_ "mybook/routers"
	_ "mybook/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
