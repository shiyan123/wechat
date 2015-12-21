package octopus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strings"
	"time"
	"wechat/models/octopus/wxBaseParameter"
)

type FirstMenu struct {
	Id         int64
	MenuName   string
	MenuType   string
	MenuKey    string
	SkipUrl    string
	MenuSort   int
	CreateTime time.Time
}
type SecondMenu struct {
	Id            int64
	FirstMenuName string
	MenuName      string
	MenuType      string
	MenuKey       string
	MessageType   string
	MessageTitle  string
	SkipUrl       string
	MenuSort      int
	CreateTime    time.Time
}
type ErrorData struct {
	Errcode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type Menu struct {
	Buttons []Button `json:"button"`
}
type Button struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Key        string      `json:"key"`
	Url        string      `json:"url"`
	SubButtons []SubButton `json:"sub_button"`
}
type SubButton struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Key  string `json:"key"`
	Url  string `json:"url"`
}

func FindFirstMenu() ([]FirstMenu, error) {
	o := orm.NewOrm()
	var firstMenus []FirstMenu
	num, err := o.Raw("select * from first_menu").QueryRows(&firstMenus)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return firstMenus, nil
	}
}
func FindSecondMenu() ([]SecondMenu, error) {
	o := orm.NewOrm()
	var secondMenus []SecondMenu
	num, err := o.Raw("select * from second_menu").QueryRows(&secondMenus)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return secondMenus, nil
	}
}
func FindOneFirstMenu() (*FirstMenu, error) {
	o := orm.NewOrm()
	firstMenu := &FirstMenu{Id: 1}
	err := o.Read(firstMenu)
	if err == orm.ErrNoRows {
		err = errors.New("查询不到")
		return nil, err
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err
	} else {
		return firstMenu, nil
	}
}
func InsertOneFirstMenu(menuName, menuType, skipUrl string, menuSort int) (*FirstMenu, error) {
	o := orm.NewOrm()
	var menuKey string
	if menuType == "click" {
		menuType = "消息触发类"
		menuKey = GetRandomString()
	} else {
		menuType = "网页链接类"
		menuKey = ""
	}
	firstMenu := &FirstMenu{
		MenuName:   menuName,
		MenuType:   menuType,
		MenuKey:    menuKey,
		SkipUrl:    skipUrl,
		MenuSort:   menuSort,
		CreateTime: time.Now(),
	}
	_, err := o.Insert(firstMenu)
	if err != nil {
		return nil, err
	}
	return firstMenu, nil
}
func InsertOneSecondMenu(firstMenuName, menuName, menuType, messageType, messageTitle string, skipUrl string, menuSort int) (*SecondMenu, error) {
	o := orm.NewOrm()
	var menuKey string
	if menuType == "click" {
		menuType = "消息触发类"
		menuKey = GetRandomString()
	} else {
		menuType = "网页链接类"
		menuKey = ""
	}
	if messageType == "text" {
		messageType = "文本消息"
	} else {
		messageType = "图文消息"
	}
	secondMenu := &SecondMenu{
		FirstMenuName: firstMenuName,
		MenuName:      menuName,
		MenuType:      menuType,
		MenuKey:       menuKey,
		MessageType:   messageType,
		MessageTitle:  messageTitle,
		SkipUrl:       skipUrl,
		MenuSort:      menuSort,
		CreateTime:    time.Now(),
	}
	_, err := o.Insert(secondMenu)
	if err != nil {
		return nil, err
	}
	return secondMenu, nil
}
func UpdateOneFirstMenu() (*FirstMenu, error) {
	o := orm.NewOrm()
	firstMenu := &FirstMenu{}
	_, err := o.Update(firstMenu)
	if err != nil {
		return nil, err
	}
	return firstMenu, nil
}
func DeleteFirstMenu(menuNames map[string]string) error {
	o := orm.NewOrm()
	for _, v := range menuNames {
		num, _ := o.QueryTable("first_menu").Filter("menuName", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
func DeleteSecondMenu(menuNames map[string]string) error {
	o := orm.NewOrm()
	for _, v := range menuNames {
		num, _ := o.QueryTable("second_menu").Filter("menuName", v).Delete()
		if num == 0 {
			return errors.New("查询不到")
		}
	}
	return nil
}
func SendCustomMenuToWechat() (*ErrorData, error) {
	url := wxBaseParameter.GetWxChatHttpUrl().CustomMenuUrl
	accessToken, err := GetOrUpdateAccessToken()
	newUrl := strings.Replace(url, "ACCESS_TOKEN", accessToken.AccessToken, -1)
	req := httplib.Post(newUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	firstMenus, err := FindFirstMenu()
	secondMenus, err := FindSecondMenu()
	var buttons []Button
	var subButtons []SubButton
	for _, f := range firstMenus {
		var menuType string
		if f.MenuType == "消息触发类" {
			menuType = "click"
		} else {
			menuType = "view"
		}
		for _, s := range secondMenus {
			if f.MenuName == s.FirstMenuName {
				var secondMenuType string
				if s.MenuType == "消息触发类" {
					secondMenuType = "click"
				} else {
					secondMenuType = "view"
				}
				subButtons = append(subButtons, SubButton{
					Name: s.MenuName,
					Type: secondMenuType,
					Key:  s.MenuKey,
					Url:  s.SkipUrl,
				})
			}
			continue
		}
		buttons = append(buttons, Button{
			Type:       menuType,
			Name:       f.MenuName,
			Key:        f.MenuKey,
			Url:        f.SkipUrl,
			SubButtons: subButtons,
		})
		subButtons = nil
	}
	menu := &Menu{
		Buttons: buttons,
	}
	a, _ := json.Marshal(menu)
	fmt.Println(string(a))
	b, _ := req.JsonBody(menu)
	fmt.Println(b)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	fmt.Println(str)
	var str2 ErrorData
	err = json.Unmarshal([]byte(str), &str2)
	if err != nil {
		return nil, err
	}
	return &str2, nil
}

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
