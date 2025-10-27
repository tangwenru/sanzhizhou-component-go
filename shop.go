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

func (this *Shop) GetLastSyncTime(userToken string, shopId int64) (*sanzhizhouComponentConfig.GetLastSyncTime, error) {
	query := sanzhizhouComponentConfig.ShopGetLastSyncTimeQuery{
		ShopId: shopId,
	}
	lastSyncTimeResult := sanzhizhouComponentConfig.GetLastSyncTimeResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getLastSyncTime", &query, &lastSyncTimeResult)

	if err != nil {
		fmt.Println("Shop GetLastSyncTime err:", string(bytesResult), err)
	}

	if !lastSyncTimeResult.Success {
		return &sanzhizhouComponentConfig.GetLastSyncTime{}, errors.New(lastSyncTimeResult.Message)
	}

	return &lastSyncTimeResult.Data, nil
}

func (this *Shop) List(
	userToken string,
	commercePlatformId int64,
	showAll bool,
) (*sanzhizhouComponentConfig.ShopListData, error) {
	query := sanzhizhouComponentConfig.ShopListResultQuery{
		ShowAll:            showAll,
		CommercePlatformId: commercePlatformId,
	}
	shopDictResult := sanzhizhouComponentConfig.ShopListResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/list", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetLastSyncTime err:", string(bytesResult), err)
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

	//去重
	newShopIdList := sanzhizhouComponentLib.RemoveDuplicatesIn64(shopIdList)

	query := sanzhizhouComponentConfig.ShopDictQuery{
		ShopIdList: *newShopIdList,
		ShowExtend: showExtend,
	}
	shopDictResult := sanzhizhouComponentConfig.ShopDictResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/dict", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetLastSyncTime err:", string(bytesResult), err)
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

func (this *Shop) ListByDomesticWarehouseId(
	userToken string,
	domesticWarehouseId int64,
) (*[]int64, error) {
	query := sanzhizhouComponentConfig.ShopByDomesticWarehouseIdQuery{
		DomesticWarehouseId: domesticWarehouseId,
	}
	shopDetailResult := sanzhizhouComponentConfig.ListByDomesticWarehouseIdResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/listByDomesticWarehouseId", &query, &shopDetailResult)

	if err != nil {
		fmt.Println("Shop Detail err:", string(bytesResult), err)
	}

	if !shopDetailResult.Success {
		return &[]int64{}, errors.New(shopDetailResult.Message)
	}

	return &shopDetailResult.Data, nil
}
