package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type User struct {
}

type UserDetailResult struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Code    string     `json:"code"`
	Data    UserDetail `json:"data"`
}

type UserDetail struct {
	Id          int64  `json:"id"`
	AccountName string `json:"accountName"`
	Nickname    string `json:"nickname"`
	AvatarUrl   string `json:"avatarUrl"`
	Enabled     bool   `json:"enabled"`
	Expired     int64  `json:"expired"`
	WebSid      string `json:"webSid"`
	Role        string `json:"role"`
	AppSid      string `json:"appSid"`
	ClientSid   string `json:"clientSid"`
	DealerId    int64  `json:"dealerId"`
	//Created     int64   `json:"created"`
	Balance float64 `json:"balance"`
}

func (this *User) Detail(userToken string) (error, *UserDetail) {
	userDetailResult := UserDetailResult{}
	query := map[string]string{}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "user/info", &query, &userDetailResult)

	userDetail := UserDetail{}
	if err != nil {
		fmt.Println("shiPinLv user info err:", err)
		return err, &userDetail
	}

	if !userDetailResult.Success {
		return errors.New(userDetailResult.Message), &userDetail
	}

	return err, &userDetailResult.Data
}

// 一个用户只能登录一个客户端
func (this *User) DetailByOneClient(userToken string, productType, clientUniqueKey string) (error, *UserDetail) {
	userDetailResult := UserDetailResult{}
	query := map[string]string{
		"productType":     productType,
		"clientUniqueKey": clientUniqueKey,
	}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "user/detailByOneClient", &query, &userDetailResult)

	userDetail := UserDetail{}
	if err != nil {
		fmt.Println("shi Pin Lv user info err:", err)
		return err, &userDetail
	}

	if !userDetailResult.Success {
		return errors.New(userDetailResult.Message), &userDetail
	}

	return err, &userDetailResult.Data
}

// 一个用户只能登录一个客户端
func (this *User) Dict(staffToken string, userIdList *[]int64) (error, *sanzhizhouComponentConfig.UserDict) {
	userDetailResult := sanzhizhouComponentConfig.UserDictResult{}
	query := map[string]interface{}{
		"userIdList": userIdList,
	}
	_, err := sanzhizhouComponentLib.MainSystem(staffToken, "user/dict", &query, &userDetailResult)

	userDict := sanzhizhouComponentConfig.UserDict{}
	if err != nil {
		fmt.Println("ssz user info err:", err)
		return err, &userDict
	}

	if !userDetailResult.Success {
		return errors.New(userDetailResult.Message), &userDict
	}

	return err, &userDetailResult.Data
}
