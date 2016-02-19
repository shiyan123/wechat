package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type SubscribeWelcomeController struct {
	beego.Controller
}

func (this *SubscribeWelcomeController) Get() {
	textMessages, err := octopus.FindAllTextMessage()
	if err != nil {
		this.TplNames = "subscribeWelcome.html"
	}
	this.Data["TextMessage"] = textMessages
	welcomeMessages, err := octopus.FindWelcomeMessage()
	if err != nil {
		this.TplNames = "subscribeWelcome.html"
	}
	this.Data["WelcomeMessage"] = welcomeMessages
	this.TplNames = "subscribeWelcome.html"
}
func (this *SubscribeWelcomeController) Post() {
	messageTitle := this.GetString("title")
	messageType := this.GetString("type")
	_, err := octopus.InsertOneWelcomeMessage(messageTitle, messageType)
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
func (this *SubscribeWelcomeController) Delete() {
	data := this.Ctx.Input.RequestBody
	messageTitles := make(map[string]string)
	err := json.Unmarshal(data, &messageTitles)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	err = octopus.DeleteWelcomeMessage(messageTitles)
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
