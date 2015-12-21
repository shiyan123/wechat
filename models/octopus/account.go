package octopus

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type WechatAccount struct {
	Id           int64
	AccountName  string
	Token        string
	WechatNumber string
	OldId        string
	AccountType  string
	Email        string
	Description  string
	Appid        string
	Appsecret    string
	CreateTime   time.Time
}

func FindAllAccount() ([]WechatAccount, error) {
	o := orm.NewOrm()
	var accounts []WechatAccount
	num, err := o.Raw("select * from wechat_account").QueryRows(&accounts)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return accounts, nil
	}
}

func InsertOneWechatAccount(accountName, token, wechatNumber, oldId, accountType, email, description, appid, appsecret string) error {
	o := orm.NewOrm()
	account := &WechatAccount{
		AccountName:  accountName,
		Token:        token,
		WechatNumber: wechatNumber,
		OldId:        oldId,
		AccountType:  accountType,
		Email:        email,
		Description:  description,
		Appid:        appid,
		Appsecret:    appsecret,
		CreateTime:   time.Now(),
	}
	_, err := o.Insert(account)
	if err != nil {
		return err
	}
	return nil
}
