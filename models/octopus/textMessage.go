package octopus

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type TextMessage struct {
	Id             int64
	MessageTitle   string
	MessageContent string
	CreateTime     time.Time
}

func InsertOneTextMessage(messageTitle, messageContent string) (*TextMessage, error) {
	o := orm.NewOrm()
	textMessage := &TextMessage{
		MessageTitle:   messageTitle,
		MessageContent: messageContent,
		CreateTime:     time.Now(),
	}
	_, err := o.Insert(textMessage)
	if err != nil {
		return nil, err
	}
	return textMessage, nil
}
func FindOneTextMessage(messageTitle string) (*TextMessage, error) {
	fmt.Println(messageTitle)
	o := orm.NewOrm()
	textMessage := &TextMessage{
		MessageTitle: messageTitle,
	}
	err := o.Read(textMessage, "message_title")
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return textMessage, nil
}
func FindAllTextMessage() ([]TextMessage, error) {
	o := orm.NewOrm()
	var textMessages []TextMessage
	num, err := o.Raw("select * from text_message").QueryRows(&textMessages)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return textMessages, nil
	}
}
func DeleteTextMessage(messageTitles map[string]string) error {
	o := orm.NewOrm()
	for _, v := range messageTitles {
		num, _ := o.QueryTable("text_message").Filter("message_title", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
