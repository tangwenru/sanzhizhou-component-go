package sanzhizhouComponent

import (
	"errors"
	"fmt"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type AppKey struct {
}

type AppKeyDetailResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Code    string       `json:"code"`
	Data    AppKeyDetail `json:"data"`
}

type AppKeyDetail struct {
	KeyType   string `json:"keyType"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

func init() {

}

func (this *AppKey) DetailByUserId(userToken, keyType string) (*AppKeyDetail, error) {
	appKeyDetailResult := AppKeyDetailResult{}
	query := map[string]string{
		"keyType": keyType,
	}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "appKey/detail", &query, &appKeyDetailResult)

	appKeyDetail := AppKeyDetail{}
	if err != nil {
		fmt.Println("app key info err:", err)
		return &appKeyDetail, err
	}

	if !appKeyDetailResult.Success {
		return &appKeyDetail, errors.New(appKeyDetailResult.Message)
	}

	return &appKeyDetailResult.Data, err
}
