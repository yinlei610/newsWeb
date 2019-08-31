package controllers

import (
	"github.com/astaxie/beego"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}



func (this *MysqlController) ShowMysql()  {
	//1. 打开数据库
	conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	if err != nil{
		beego.Info("连接错误",err)
		return
	}

	//3. 关闭数据库
	defer conn.Close()

	/*_, err = conn.Exec("CREATE TABLE userInfo(id INT,name VARCHAR(11))", 1)
	if err != nil{
		beego.Info("创建表错误",err)
		this.Data["errmsg"] = ""
		this.TplName = "err.html"
		return
	}*/
	//2. 操作数据库
	/*_, err = conn.Exec("CREATE TABLE userInfo(id INT,name VARCHAR(11))")
	if err != nil{
		beego.Info("创建表错误",err)
		return
	}
	//插入数据
	_, err = conn.Exec("INSERT userInfo(id,name) VALUE (?,?)", 1, "lilei")
	if err != nil{
		beego.Info("插入数据错误",err)
		return
	}*/

	//作业1：更新，删除
	/*_, err = conn.Exec("insert into userInfo (id,name) values (?,?)", 2, "zhangsan")
	if err != nil{
		beego.Info("插入数据错误",err)
		return
	}*/

	/*_, err = conn.Exec("DELETE FROM userInfo WHERE id=2")
	if err != nil{
		beego.Info("删除数据错误",err)
		return
	}*/

	//查询
	rows, err := conn.Query("select id from userInfo")
	var id int
	for rows.Next(){
		rows.Scan(&id)
		beego.Info(id)
	}


	this.Ctx.WriteString("插入数据成功")

}
