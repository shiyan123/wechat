package controllers

import (
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type NewsMessageListController struct {
	beego.Controller
}

func (this *NewsMessageListController) Get() {
	newsMessageLists, err, _ := octopus.FindAllNewsMessageList()
	if err != nil {
		this.TplNames = "newsMessageList.html"
	}
	this.Data["NewsMessageLists"] = newsMessageLists
	this.TplNames = "newsMessageList.html"
}
func (this *NewsMessageListController) Put() {
	_, err, cnt := octopus.FindAllNewsMessageList()
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	err = octopus.GetNewsMessageList(cnt)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	message := &Message{
		Message: "操作成功",
	}
	this.Data["json"] = message
	this.ServeJson()
}
