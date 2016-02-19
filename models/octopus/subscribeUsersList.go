package octopus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strings"
	"wechat/models/octopus/wxBaseParameter"
)

type SubscribeUser struct {
	Id            int64
	Subscribe     int    `json:"subscribe"`
	OpenId        string `json:"openid"`
	Sex           int    `json:"sex"`
	NickName      string `json:"nickname"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	Headimgurl    string `json:"headimgurl"`
	SubscribeTime int    `json:"subscribe_time"`
	Unionid       string `json:"unionid"`
	Remark        string `json:"remark"`
	Groupid       int    `json:"groupid"`
}
type UsersList struct {
	Total      int         `json:"total"`
	Count      int         `json:"count"`
	AcceptData OpenidsData `json:"data"`
	NextOpenid string      `json:"next_openid"`
}

type OpenidsData struct {
	Openids []string `json:"openid"`
}

func FindAllSubscribeUsers() ([]SubscribeUser, error) {
	o := orm.NewOrm()
	var subscribeUsers []SubscribeUser
	num, err := o.Raw("select * from subscribe_user").QueryRows(&subscribeUsers)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return subscribeUsers, nil
	}
}
func InsertSubscribeUsers(subscribeUsers []SubscribeUser) error {
	o := orm.NewOrm()
	qs := o.QueryTable("subscribe_user")
	i, _ := qs.PrepareInsert()
	for _, subscribeUser := range subscribeUsers {
		fmt.Println(subscribeUser.NickName)
		id, err := i.Insert(&subscribeUser)
		if err != nil {
			return err
		}
		fmt.Println(id)
	}
	i.Close()
	return nil
}
func GetSubscribeUserList(nextOpenid string) error {
	url := wxBaseParameter.GetWxChatHttpUrl().SubscribeUsersListUrl
	accessToken, _ := GetOrUpdateAccessToken()
	newUrl := strings.Replace(strings.Replace(url, "ACCESS_TOKEN", accessToken.AccessToken, -1), "NEXT_OPENID", nextOpenid, -1)
	req := httplib.Get(newUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	if err != nil {
		return err
	}
	var usersList UsersList
	var subscribeUsers []SubscribeUser
	var subscribeUser SubscribeUser
	err = json.Unmarshal([]byte(str), &usersList)
	if err != nil {
		return err
	}
	sem := make(chan int, usersList.Count)
	for i := 0; i < usersList.Count; {
		go func(i int) {
			_, str1 := GetUserBaseInfo(accessToken, i, usersList)
			err = json.Unmarshal([]byte(str1), &subscribeUser)
			subscribeUsers = append(subscribeUsers, subscribeUser)
			fmt.Println(i)
			sem <- 0
		}(i)
		i++
		continue
	}
	for j := 0; j < usersList.Count; j++ {
		<-sem
	}
	err = InsertSubscribeUsers(subscribeUsers)
	if err != nil {
		return err
	}
	return nil
}

func GetUserBaseInfo(accessToken *AccessToken, i int, usersList UsersList) (error, string) {
	url := wxBaseParameter.GetWxChatHttpUrl().SubscribeUsersUrl
	newUrl := strings.Replace(strings.Replace(url, "ACCESS_TOKEN", accessToken.AccessToken, -1), "OPENID", usersList.AcceptData.Openids[i], -1)
	req := httplib.Get(newUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	if err != nil {
		return err, ""
	}
	return nil, str
}
