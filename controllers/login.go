package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)
type LoginController struct {
	beego.Controller
}
func(this *LoginController)Get(){
	isExist := this.Input().Get("exit") == "true"
	if isExist{
		this.Ctx.SetCookie("username","",-1,"/")
		this.Ctx.SetCookie("password","",-1,"/")
		this.Redirect("/",301)
		return
	}
	this.TplName = "login.html"
}
func (this *LoginController) Post(){
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autologin:= this.Input().Get("autologin") == "on"
	if beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password{
		maxAge := 0
		if autologin{
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("username",username,maxAge,"/")
		this.Ctx.SetCookie("password",password,maxAge,"/")
	}
	this.Redirect("/",301)
	return
}
func CheckAccount(ctx *context.Context) bool {
	ck,err :=ctx.Request.Cookie("username")
	if nil != err{
		return false
	}
	username := ck.Value
	ck,err = ctx.Request.Cookie("password")
	if nil != err{
		return false
	}
	password:=ck.Value
	return beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password
}