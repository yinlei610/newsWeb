package main

import (
	_ "ClassOne/routers"
	"github.com/astaxie/beego"
	_ "ClassOne/models"
)

func main() {
	beego.Run()
}

