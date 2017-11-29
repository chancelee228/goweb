package main

import (
	_ "myBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myBlog/controllers"
	"myBlog/models"
)

func init()  {
	//注册数据库
	models.RegisterDB()
}
func main() {
	//开启orm调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default",false,true)//database name,delete it and recreate it ,print log
	//注册路由
	beego.Router("/",&controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/category",&controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/topic",&controllers.TopicController{})
	//启动
	beego.Run()
}

