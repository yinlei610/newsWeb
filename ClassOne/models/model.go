package models

import "github.com/astaxie/beego/orm"

//1.放结构体，即表设计

//2.放初始化语句

type User struct {
	Id int
	Name string
}

func init(){
	//1.连接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	//2.注册表 （在sql中为操作数据库）
	orm.RegisterModel(new(User))
	//3.生成表
	orm.RunSyncdb("default", false, true)

}
