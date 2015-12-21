package models

import (
	"github.com/astaxie/beego/orm"
	"wechat/models/octopus"
)

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(octopus.WechatAccount), new(octopus.AccessToken), new(octopus.FirstMenu), new(octopus.SecondMenu), new(octopus.TextMessage), new(octopus.WelcomeMessage))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/wechat?charset=utf8") //密码为空格式
}
