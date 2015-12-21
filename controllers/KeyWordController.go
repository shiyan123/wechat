package controllers

import (
	"github.com/astaxie/beego"
)

type KeyWordController struct {
	beego.Controller
}

func (this *KeyWordController) Get() {

	this.TplNames = "keyWord.html"
}
