package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
)
type LoginController struct {
	beego.Controller
}
func(this *LoginController)Get(){
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
func checkAccount(ctx *context.Context,input *url.Values){
	
}