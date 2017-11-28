package controllers

import (
	"github.com/astaxie/beego"
	"myBlog/models"
)
type CategoryController struct {
	beego.Controller
}
func (this *CategoryController)Get(){
	op:=this.Input().Get("op")
	switch op {
	case "add":
		cname:=this.Input().Get("cname")
		if len(cname) == 0 {
			break
		}
		err:=models.AddCategory(cname)
		if nil != err{
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	case "del":
		id:=this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err:=models.DelCategoryById(id)
		if nil != err{
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	var err error
	this.Data["Categories"],err = models.GetAllCategories()
	if nil != err{
		beego.Error(err)
	}
}