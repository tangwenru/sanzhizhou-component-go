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
		return nil, err
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
		return nil, err
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
		shopDictResult.Message = err.Error()
		return &shopDictResult.Data, err
	}

	if !shopDictResult.Success {
		return &map[int64]sanzhizhouComponentConfig.ShopBaseInfo{}, errors.New(shopDictResult.Message)
	}

	return &shopDictResult.Data, nil
}

func (this *Shop) Detail(
	userToken string,
	shopId int64,
) (*sanzhizhouComponentConfig.ShopDetail, *sanzhizhouComponentConfig.ShopDetailResult) {
	query := sanzhizhouComponentConfig.ShopDetailQuery{
		Id: shopId,
	}
	shopDetailResult := sanzhizhouComponentConfig.ShopDetailResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/detail", &query, &shopDetailResult)

	if err != nil {
		fmt.Println("Shop Detail err:", string(bytesResult), err)
		shopDetailResult.Message = err.Error()
		return &shopDetailResult.Data, &shopDetailResult
	}

	if !shopDetailResult.Success {
		return &sanzhizhouComponentConfig.ShopDetail{}, &shopDetailResult
	}

	return &shopDetailResult.Data, &shopDetailResult
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
		return nil, err
	}

	if !shopDetailResult.Success {
		return &[]int64{}, errors.New(shopDetailResult.Message)
	}

	return &shopDetailResult.Data, nil
}

func (this *Shop) SaveStatus(
	userToken string,
	query *sanzhizhouComponentConfig.ShopSaveStatusQuery,
) error {
	saveResult := sanzhizhouComponentConfig.ShopSaveStatusQueryResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/saveStatus", &query, &saveResult)

	if err != nil {
		fmt.Println("Shop saveStatus err:", string(bytesResult), err)
		return err
	}

	if !saveResult.Success {
		return errors.New(saveResult.Message)
	}

	return nil
}

func (this *Shop) SaveLastSyncTime(
	userToken string,
	query *sanzhizhouComponentConfig.ShopSaveLastSyncTimeQuery,
) error {
	saveResult := sanzhizhouComponentConfig.ShopSaveLastSyncTimeResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/SaveLastSyncTime", &query, &saveResult)

	if err != nil {
		fmt.Println("Shop SaveLastSyncTime err:", string(bytesResult), err)
		return err
	}

	if !saveResult.Success {
		return errors.New(saveResult.Message)
	}

	return nil
}

func (this *Shop) GetPublishShopIdList(
	userToken string,
	commercePlatformId int64,
) (*[]int64, error) {
	saveResult := sanzhizhouComponentConfig.ShopGetPublishShopIdListResult{}

	query := sanzhizhouComponentConfig.ShopGetPublishShopIdListQuery{
		CommercePlatformId: commercePlatformId,
	}
	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getPublishShopIdList", &query, &saveResult)

	if err != nil {
		fmt.Println("Shop GetPublishShopIdList err:", string(bytesResult), err)
		return nil, err
	}

	if !saveResult.Success {
		return &[]int64{}, errors.New(saveResult.Message)
	}

	return &saveResult.Data, nil
}
