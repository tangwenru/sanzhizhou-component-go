package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Shop struct {
}

type ShopGetOrderLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type GetOrderLastSyncTimeResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    struct {
		OrderLastSyncTime int64 `json:"orderLastSyncTime"`
	} `json:"data"`
}

func init() {

}

func (this *Shop) GetOrderLastSyncTime(userToken string, shopId int64) (int64, error) {
	query := ShopGetOrderLastSyncTimeQuery{
		ShopId: shopId,
	}
	orderLastSyncTimeResult := GetOrderLastSyncTimeResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getOrderLastSyncTime", &query, &orderLastSyncTimeResult)

	if err != nil {
		fmt.Println("Shop GetOrderLastSyncTime err:", string(bytesResult), err)
	}

	if !orderLastSyncTimeResult.Success {
		return 0, errors.New(orderLastSyncTimeResult.Message)
	}

	return orderLastSyncTimeResult.Data.OrderLastSyncTime, nil
}
