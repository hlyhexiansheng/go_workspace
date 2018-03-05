package main

import (
	_ "my_bee_project/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	beego.Run()
}

