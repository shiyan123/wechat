package octopus

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type WelcomeMessage struct {
	Id           int64
	MessageTitle string
	MessageType  string
	CreateTime   time.Time
}

func InsertOneWelcomeMessage(messageTitle, messageType string) (*WelcomeMessage, error) {
	o := orm.NewOrm()
	var Type string
	if messageType == "text" {
		Type = "文本消息"
	} else {
		Type = "图文消息"
	}
	welcomeMessage := &WelcomeMessage{
		MessageTitle: messageTitle,
		MessageType:  Type,
		CreateTime:   time.Now(),
	}
	_, err := o.Insert(welcomeMessage)
	if err != nil {
		return nil, err
	}
	return welcomeMessage, nil
}
func DeleteWelcomeMessage(messageTitles map[string]string) error {
	o := orm.NewOrm()
	for _, v := range messageTitles {
		num, _ := o.QueryTable("welcome_message").Filter("message_title", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
func FindWelcomeMessage() ([]WelcomeMessage, error) {
	o := orm.NewOrm()
	var welcomeMessages []WelcomeMessage
	num, err := o.Raw("select * from welcome_message").QueryRows(&welcomeMessages)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return welcomeMessages, nil
	}
}
