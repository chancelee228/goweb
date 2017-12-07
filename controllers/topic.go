package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
	//"fmt"
	"time"
	"fmt"
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
	category:=this.Input().Get("category")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@")
	fmt.Println(category)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@")
	var err error
	tid:=this.Input().Get("tid")
	if len(tid) == 0{
		err=models.AddTopic(title,category,content)
	}else {
		err=models.ModifyTopic(tid,title,category,content)
	}
	if nil != err{
		beego.Error(err)
	}
	this.Redirect("/topic",302)
}

type ShowTopic struct {
	Id int64
	Title string
	Updated time.Time
	Views int64
	Category string
	ReplyCount int64
}
func (this *TopicController) Get(){
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	var err error
	this.Data["Topics"],err = models.GetAllTopics(false)

	/*
	s,err:= models.GetAllTopics(false)
	fmt.Println("==========1111111111============")
	fmt.Println(s[0].Category.Title)
	fmt.Println("==========1111111111============")
	ret := make([]ShowTopic,len(s))
	for i:=0;i<len(s);i++{
		ret[i] = ShowTopic{
			Id:s[i].Id,
			Title:s[i].Title,
			Updated:s[i].Updated,
			Views:s[i].Views,
			Category:s[i].Category.Title,
			ReplyCount:s[i].ReplyCount,
		}
	}
	this.Data["Topics"] = ret
	*/
	if nil != err{
		beego.Error(err)
	}
	return
}
func (this *TopicController) Add()  {
	this.TplName = "topic_add.html"
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	var err error
	this.Data["Categories"],err = models.GetAllCategories()
	//this.Ctx.WriteString("add topic")
	if nil != err{
		beego.Error(err)
	}
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

func (this *TopicController) Modify(){
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic,err:=models.GetTopicById(tid)
	if nil != err{
		beego.Error(err)
		this.Redirect("/",302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.Data["Categories"],err = models.GetAllCategories()
	if nil != err{
		beego.Error(err)
	}
}

func (this *TopicController) Delete(){
	if !CheckAccount(this.Ctx){
		this.Redirect("/login",302)
		return
	}
	err:=models.DeleteTopicById(this.Input().Get("tid"))
	if nil != err{
		beego.Error(err)
	}
	this.Redirect("/",302)
}