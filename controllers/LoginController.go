package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.TplNames = "login.html"
}
func (this *LoginController) Post() {
	if this.GetString("email") != "admin" || this.GetString("password") != "123456" {
		this.Data["message"] = "用户名或密码错误请重新登陆"
		this.TplNames = "login.html"
	} else {
		this.Data["email"] = this.GetString("email")
		this.Data["password"] = this.GetString("password")
		this.TplNames = "index.html"
	}

}
