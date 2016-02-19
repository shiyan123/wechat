package octopus

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type KeyWord struct {
	Id           uint64
	KeyWordName  string
	MessageTitle string
	MessageType  string
	CreateTime   time.Time
}

func InsertOneKeyWord(keyWordName, messageTitle, messageType string) (*KeyWord, error) {
	o := orm.NewOrm()
	var Type string
	if messageType == "text" {
		Type = "文本消息"
	} else {
		Type = "图文消息"
	}
	keyWord := &KeyWord{
		KeyWordName:  keyWordName,
		MessageTitle: messageTitle,
		MessageType:  Type,
		CreateTime:   time.Now(),
	}
	_, err := o.Insert(keyWord)
	if err != nil {
		return nil, err
	}
	return keyWord, nil
}

func FindOneKeyWord(keyWordName string) (*KeyWord, error) {
	o := orm.NewOrm()
	keyWord := &KeyWord{
		KeyWordName: keyWordName,
	}
	err := o.Read(keyWord, "key_word")
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return keyWord, nil
}
func DeleteKeyWord(keyWordNames map[string]string) error {
	o := orm.NewOrm()
	for _, v := range keyWordNames {
		num, _ := o.QueryTable("key_word").Filter("key_word_name", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
func FindKeyWords() ([]KeyWord, error) {
	o := orm.NewOrm()
	var keyWords []KeyWord
	num, err := o.Raw("select * from key_word").QueryRows(&keyWords)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return keyWords, nil
	}
}
