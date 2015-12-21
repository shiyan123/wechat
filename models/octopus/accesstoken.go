package octopus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	"wechat/models/octopus/wxBaseParameter"
)

type AccessToken struct {
	Id          int64
	AccessToken string
	ExpiresIn   int
	CreateTime  time.Time
}
type Data struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func FindAccessToken() (*AccessToken, error) {
	o := orm.NewOrm()
	accessToken := &AccessToken{Id: 1}
	err := o.Read(accessToken)
	if err == orm.ErrNoRows {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return accessToken, nil
	}
}
func GetOrUpdateAccessToken() (*AccessToken, error) {
	accesstoken, err := FindAccessToken()
	if err != nil && accesstoken == nil {
		accounts, err := FindAllAccount()
		if err != nil {
			return nil, err
		}
		data, err := GetWxAccessToken(accounts[0].Appid, accounts[0].Appsecret)
		if err != nil {
			return nil, err
		}
		accesstoken, err = InsertOneAccessToken(data.AccessToken, data.ExpiresIn)
		if err != nil {
			return nil, err
		}
		return accesstoken, nil
	}
	if accesstoken.CreateTime.Add(2*time.Hour).After(time.Now().Add(-8*time.Hour)) == false {
		accounts, err := FindAllAccount()
		if err != nil {
			return nil, err
		}
		data, err := GetWxAccessToken(accounts[0].Appid, accounts[0].Appsecret)
		if err != nil {
			return nil, err
		}
		accesstoken, err = UpdateOneAccessToken(data.AccessToken, data.ExpiresIn)
		if err != nil {
			return nil, err
		}
		return accesstoken, nil
	}
	return accesstoken, nil
}
func InsertOneAccessToken(accesstoken string, expiresIn int) (*AccessToken, error) {
	o := orm.NewOrm()
	accessToken := &AccessToken{
		AccessToken: accesstoken,
		ExpiresIn:   expiresIn,
		CreateTime:  time.Now(),
	}
	_, err := o.Insert(accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
func UpdateOneAccessToken(accesstoken string, expiresIn int) (*AccessToken, error) {
	o := orm.NewOrm()
	accessToken := &AccessToken{
		Id:          1,
		AccessToken: accesstoken,
		ExpiresIn:   expiresIn,
		CreateTime:  time.Now(),
	}
	_, err := o.Update(accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
func GetWxAccessToken(appid, appsecret string) (*Data, error) {
	url := wxBaseParameter.GetWxChatHttpUrl().AccessTokenUrl
	newUrl := strings.Replace(strings.Replace(url, "APPID", appid, -1), "APPSECRET", appsecret, -1)
	req := httplib.Get(newUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	fmt.Println(str)
	var str2 Data
	err = json.Unmarshal([]byte(str), &str2)
	if err != nil {
		return nil, err
	}
	return &str2, nil
}
