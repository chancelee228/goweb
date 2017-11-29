package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
)

type TopicController struct {
	beego.Controller
}
func (this *TopicController) Post(){
	if !CheckAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	title:=this.Input().Get("title")
	content:=this.Input().Get("content")
	var err error
	err=models.AddTopic(title,content)
	if nil != err{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}
func (this *TopicController) Get(){
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	var err error
	this.Data["Topics"],err = models.GetAllTopics(false)
	if nil != err{
		beego.Error(err)
	}
	return
}
func (this *TopicController) Add()  {
	this.TplName = "topic_add.html"
	//this.Ctx.WriteString("add topic")
}

func (this *TopicController) View(){
	this.TplName = "topic_view.html"
	topic,err := models.GetTopicById(this.Ctx.Input.Param("0"))
	if nil != err{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}