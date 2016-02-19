package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type KeyWordController struct {
	beego.Controller
}

func (this *KeyWordController) Get() {
	textMessages, err := octopus.FindAllTextMessage()
	if err != nil {
		this.TplNames = "keyWord.html"
	}
	this.Data["TextMessage"] = textMessages
	newsMessagesLists, err, _ := octopus.FindAllNewsMessageList()
	if err != nil {
		this.TplNames = "keyWord.html"
	}
	this.Data["NewsMessagesLists"] = newsMessagesLists
	keyWords, err := octopus.FindKeyWords()
	if err != nil {
		this.TplNames = "keyWord.html"
	}
	this.Data["KeyWords"] = keyWords
	this.TplNames = "keyWord.html"
}
func (this *KeyWordController) Post() {
	keyWordName := this.GetString("name")
	modelMessage := this.GetString("modelMessage")
	messageType := this.GetString("type")

	_, err := octopus.InsertOneKeyWord(keyWordName, modelMessage, messageType)
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
func (this *KeyWordController) Delete() {
	data := this.Ctx.Input.RequestBody
	keyWordNames := make(map[string]string)
	err := json.Unmarshal(data, &keyWordNames)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	err = octopus.DeleteKeyWord(keyWordNames)
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
