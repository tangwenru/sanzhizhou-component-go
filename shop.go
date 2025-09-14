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

type ShopBaseInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	NameRemark     string `json:"nameRemark"`
	GroupName      string `json:"groupName"`
	Currency       string `json:"currency"`
	PublishSetting string `json:"publishSetting"` // 可能为空，依赖接口参数
	ApiInfo        string `json:"apiInfo"`        // 可能为空，依赖接口参数
}

type ShopDictQuery struct {
	ShopIdList []int64 `json:"shopIdList"`
	ShowExtend bool    `json:"showExtend"`
}

type ShopDictPostQuery struct {
}

type ShopDictResult struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message,omitempty"`
	Data    map[int64]ShopBaseInfo `json:"data"`
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

func (this *Shop) Dict(
	userToken string,
	shopIdList *[]int64,
	showExtend bool,
) (*map[int64]ShopBaseInfo, error) {
	query := ShopDictQuery{
		ShopIdList: *shopIdList,
		ShowExtend: showExtend,
	}
	shopDictResult := ShopDictResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/dict", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetOrderLastSyncTime err:", string(bytesResult), err)
	}

	if !shopDictResult.Success {
		return &map[int64]ShopBaseInfo{}, errors.New(shopDictResult.Message)
	}

	return &shopDictResult.Data, nil
}
