package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "区块链"
	c.TplName = "test.html"	//渲染
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Post() {
	this.Data["test"] = "区块链一期最棒"
	this.TplName = "test.html"

}
func (this *IndexController)ShowGet()  {
	this.Data["data"] = "这是get请求对应的方法"
	this.TplName = "test.html"
}