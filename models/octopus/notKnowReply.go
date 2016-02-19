package octopus

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type NotKnowReply struct {
	Id           int64
	MessageTitle string
	MessageType  string
	CreateTime   time.Time
}

func InsertOneNotKnowReply(messageTitle, messageType string) (*NotKnowReply, error) {
	o := orm.NewOrm()
	var Type string
	if messageType == "text" {
		Type = "文本消息"
	} else {
		Type = "图文消息"
	}
	notKnowReply := &NotKnowReply{
		MessageTitle: messageTitle,
		MessageType:  Type,
		CreateTime:   time.Now(),
	}
	_, err := o.Insert(notKnowReply)
	if err != nil {
		return nil, err
	}
	return notKnowReply, nil
}

func FindOneNotKnowReply(messageTitle string) (*NotKnowReply, error) {
	o := orm.NewOrm()
	notKnowReply := &NotKnowReply{
		MessageTitle: messageTitle,
	}
	err := o.Read(notKnowReply, "not_know_reply")
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return notKnowReply, nil
}
func DeleteNotKnowReply(messageTitles map[string]string) error {
	o := orm.NewOrm()
	for _, v := range messageTitles {
		num, _ := o.QueryTable("not_know_reply").Filter("message_title", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
func FindNotKnowReplys() ([]NotKnowReply, error) {
	o := orm.NewOrm()
	var notKnowReplys []NotKnowReply
	num, err := o.Raw("select * from not_know_reply").QueryRows(&notKnowReplys)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return notKnowReplys, nil
	}
}
