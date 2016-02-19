package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type CustomMenuController struct {
	beego.Controller
}

func (this *CustomMenuController) Get() {
	firstMenus, err := octopus.FindFirstMenu()
	secondMenus, err := octopus.FindSecondMenu()
	if err != nil {
		this.TplNames = "customMenu.html"
	}
	this.Data["SecondMenus"] = secondMenus
	this.Data["FirstMenus"] = firstMenus
	this.TplNames = "customMenu.html"
}

func (this *CustomMenuController) Post() {
	var err error
	menuName := this.GetString("name")
	menuType := this.GetString("type")
	skipUrl := this.GetString("url")
	menuSort, _ := this.GetInt("sort")
	firstMenuName := this.GetString("firstMenuName")
	if firstMenuName != "" {
		messageTitle := this.GetString("messageTitle")
		messageType := this.GetString("messageType")
		_, err = octopus.InsertOneSecondMenu(firstMenuName, menuName, menuType, messageType, messageTitle, skipUrl, menuSort)
	} else {
		_, err = octopus.InsertOneFirstMenu(menuName, menuType, skipUrl, menuSort)
	}
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
func (this *CustomMenuController) Delete() {
	data := this.Ctx.Input.RequestBody
	menuNames := make(map[string]string)
	err := json.Unmarshal(data, &menuNames)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	err = octopus.DeleteFirstMenu(menuNames)
	if err != nil {
		err = octopus.DeleteSecondMenu(menuNames)
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
	message := &Message{
		Message: "操作成功",
	}
	this.Data["json"] = message
	this.ServeJson()
}
func (this *CustomMenuController) Put() {
	_, err := octopus.SendCustomMenuToWechat()
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
