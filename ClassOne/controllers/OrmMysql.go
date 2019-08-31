package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ClassOne/models"
)

type OrmController struct {
	beego.Controller
}


func (this *OrmController) ShowOrm()  {


	/*o := orm.NewOrm()
	//2.插入对象
	user := User{}
	user.Name = "hanmeimei"
	//3.执行插入操作
	_, err := o.Insert(&user)
	if err != nil{
		beego.Info("插入失败", err)
		return
	}
	this.Ctx.WriteString("插入成功")
*/
	//查询
		//1.获取orm对象
		o := orm.NewOrm()
		//2.获取查询对象
		user := models.User{Id:1}

		//3.执行查询操作
		err := o.Read(&user)
		if err != nil{
			beego.Info("查询失败", err)
			return
		}
		beego.Info(user.Name)
		this.Ctx.WriteString("查询成功")
//修改


}

func (this *OrmController) ShowInsert() {
	//1.获取连接对象，orm
	o := orm.NewOrm()
	//2.插入对象
	user := models.User{Id:1}
	user.Name = "hanmeimei"
	//3.执行插入操作
	_, err := o.Insert(&user)
	if err != nil{
		beego.Info("插入失败", err)
		return
	}
	this.Ctx.WriteString("插入成功")
}

func (this *OrmController)ShowUpdate() {
	//1.获取orm对象
	o := orm.NewOrm()
	//2.获取更新对象
	user := models.User{Id:1}
	//3.查询对象
	err := o.Read(&user)
	if err != nil{
		beego.Info("查询数据错误", err)
		return
	}
	//4.给查询到的对象赋值
	user.Name = "lilei"
	//5.更新操作
	_, err = o.Update(&user)
	//删除操作 o.delete
	if err != nil{
		beego.Info("更新数据错误", err)
		return
	}
	this.Ctx.WriteString("更新成功")
}