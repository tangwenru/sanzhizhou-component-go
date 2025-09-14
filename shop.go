package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Shop struct {
}

func init() {

}

func (this *Shop) GetOrderLastSyncTime(userToken string, shopId int64) (int64, error) {
	query := sanzhizhouComponentConfig.ShopGetOrderLastSyncTimeQuery{
		ShopId: shopId,
	}
	orderLastSyncTimeResult := sanzhizhouComponentConfig.GetOrderLastSyncTimeResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getOrderLastSyncTime", &query, &orderLastSyncTimeResult)

	if err != nil {
		fmt.Println("Shop GetOrderLastSyncTime err:", string(bytesResult), err)
	}

	if !orderLastSyncTimeResult.Success {
		return 0, errors.New(orderLastSyncTimeResult.Message)
	}

	return orderLastSyncTimeResult.Data.OrderLastSyncTime, nil
}

func (this *Shop) List(
	userToken string,
	showAll bool,
) (*sanzhizhouComponentConfig.ShopListData, error) {
	query := sanzhizhouComponentConfig.ShopListResultQuery{
		ShowAll: showAll,
	}
	shopDictResult := sanzhizhouComponentConfig.ShopListResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/list", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetOrderLastSyncTime err:", string(bytesResult), err)
	}

	if !shopDictResult.Success {
		return &sanzhizhouComponentConfig.ShopListData{}, errors.New(shopDictResult.Message)
	}

	return &shopDictResult.Data, nil
}

func (this *Shop) Dict(
	userToken string,
	shopIdList *[]int64,
	showExtend bool,
) (*map[int64]sanzhizhouComponentConfig.ShopBaseInfo, error) {
	query := sanzhizhouComponentConfig.ShopDictQuery{
		ShopIdList: *shopIdList,
		ShowExtend: showExtend,
	}
	shopDictResult := sanzhizhouComponentConfig.ShopDictResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/dict", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetOrderLastSyncTime err:", string(bytesResult), err)
	}

	if !shopDictResult.Success {
		return &map[int64]sanzhizhouComponentConfig.ShopBaseInfo{}, errors.New(shopDictResult.Message)
	}

	return &shopDictResult.Data, nil
}

func (this *Shop) Detail(
	userToken string,
	shopId int64,
) (*sanzhizhouComponentConfig.ShopDetail, error) {
	query := sanzhizhouComponentConfig.ShopDetailQuery{
		Id: shopId,
	}
	shopDetailResult := sanzhizhouComponentConfig.ShopDetailResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/detail", &query, &shopDetailResult)

	if err != nil {
		fmt.Println("Shop Detail err:", string(bytesResult), err)
	}

	if !shopDetailResult.Success {
		return &sanzhizhouComponentConfig.ShopDetail{}, errors.New(shopDetailResult.Message)
	}

	return &shopDetailResult.Data, nil
}
