package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	topics,err:=models.GetAllTopics(true)
	if nil != err{
		beego.Error(err)
		return
	}
	this.Data["Topics"] = topics
}
