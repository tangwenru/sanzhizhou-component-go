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

	empty := sanzhizhouComponentConfig.GetLastSyncTime{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getLastSyncTime", &query, &lastSyncTimeResult)

	if err != nil {
		fmt.Println("Shop GetLast SyncTime err:", string(bytesResult), err)
		return &empty, err
	}

	if !lastSyncTimeResult.Success {
		return &empty, errors.New(lastSyncTimeResult.Message)
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
	empty := sanzhizhouComponentConfig.ShopListData{}
	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/list", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Shop GetLast SyncTime err:", string(bytesResult), err)
		return &empty, err
	}

	if !shopDictResult.Success {
		return &empty, errors.New(shopDictResult.Message)
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

	empty := make(map[int64]sanzhizhouComponentConfig.ShopBaseInfo)

	if err != nil {
		fmt.Println("Shop GetLastSyncTime err:", string(bytesResult), err)
		shopDictResult.Message = err.Error()
		return &empty, err
	}

	if !shopDictResult.Success {
		return &empty, errors.New(shopDictResult.Message)
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

	empty := sanzhizhouComponentConfig.ShopDetail{}

	if err != nil {
		fmt.Println("Shop Detail err:", string(bytesResult), err)
		shopDetailResult.Message = err.Error()
		return &empty, &shopDetailResult
	}

	if !shopDetailResult.Success {
		return &empty, &shopDetailResult
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
	empty := make([]int64, 0)
	if err != nil {
		fmt.Println("Shop Detail err:", string(bytesResult), err)
		return &empty, err
	}

	if !shopDetailResult.Success {
		return &empty, errors.New(shopDetailResult.Message)
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

	empty := make([]int64, 0)

	if err != nil {
		fmt.Println("Shop GetPublish-ShopIdList err:", string(bytesResult), err)
		return &empty, err
	}

	if !saveResult.Success {
		return &empty, errors.New(saveResult.Message)
	}

	return &saveResult.Data, nil
}

// 单纯的获取 api info
func (this *Shop) GetApiInfoDict(
	userToken string,
	shopIdList *[]int64,
) (*sanzhizhouComponentConfig.ShopGetApiInfoDictResultData, *sanzhizhouComponentConfig.ShopGetApiInfoDictResult) {
	saveResult := sanzhizhouComponentConfig.ShopGetApiInfoDictResult{}
	empty := sanzhizhouComponentConfig.ShopGetApiInfoDictResultData{}

	if len(*shopIdList) == 0 {
		saveResult.Message = "shopIdList 不能为空"
		return &empty, &saveResult
	}

	query := sanzhizhouComponentConfig.ShopApiInfoDictQuery{
		ShopIdList: *shopIdList,
	}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "shop/getApiInfoDict", &query, &saveResult)

	if err != nil {
		//fmt.Println("Shop ApiInfoDict err:", string(bytesResult), err)
		return &empty, &saveResult
	}

	if !saveResult.Success {
		return &empty, &saveResult
	}

	return &saveResult.Data, &saveResult
}

func (this *Shop) ApiInfoList(
	staffToken string,
	commercePlatformId int64,
	current,
	pageSize int,
) (
	*sanzhizhouComponentConfig.ShopApiInfoListResultData,
	*sanzhizhouComponentConfig.ShopApiInfoListResult,
) {
	saveResult := sanzhizhouComponentConfig.ShopApiInfoListResult{}
	empty := sanzhizhouComponentConfig.ShopApiInfoListResultData{}

	if current <= 0 {
		current = 1
	}

	if pageSize <= 0 {
		pageSize = 1
	}

	query := sanzhizhouComponentConfig.ShopApiInfoListQuery{
		CommercePlatformId: commercePlatformId,
		Current:            current,
		PageSize:           pageSize,
	}
	_, err := sanzhizhouComponentLib.MainSystem(staffToken, "shop/apiInfoList", &query, &saveResult)

	if err != nil {
		//fmt.Println("Shop ApiInfoDict err:", string(bytesResult), err)
		return &empty, &saveResult
	}

	if !saveResult.Success {
		return &empty, &saveResult
	}

	return &saveResult.Data, &saveResult
}
