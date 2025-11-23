package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//GetLastSyncTime()
	//ShopList()
	//ShopDetail()
	//ShopDict()
	//ShopSaveStatus()
	//SaveLastSyncTime()
	GetPublishShopIdList()
}

func GetLastSyncTime() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.GetLastSyncTime(
		"a12176 44f8",
		20,
	)

	fmt.Println("GetLastSyncTime result:", result)
	fmt.Println("GetLastSyncTime err:", err)
}

func ShopList() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.List(
		"a1217689d4cb 4f8",
		7,
		true,
	)

	fmt.Println(fmt.Sprintf("ShopList result: %+v", result))
	fmt.Println("ShopList err:", err)
}

func ShopDict() {
	shop := sanzhizhouComponent.Shop{}
	//去重

	result, err := shop.Dict(
		mainConfig.UserToken,
		&[]int64{5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 8, 9, 19, 20},
		true,
	)

	fmt.Println(fmt.Sprintf("ShopDict result: %+v", result))
	fmt.Println("ShopDict err:", err)
}

func ShopDetail() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.Detail(
		"a1217689d4c f44f8",
		20,
	)

	fmt.Println(fmt.Sprintf("ShopDetail result: %+v", result))
	fmt.Println("ShopDetail err:", err)
}

func ShopSaveStatus() {
	shop := sanzhizhouComponent.Shop{}
	result := shop.SaveStatus(
		mainConfig.UserToken,
		&sanzhizhouComponentConfig.ShopSaveStatusQuery{
			ShopId:        19,
			IsBan:         "1",
			Enabled:       "0",
			StatusMessage: "营业执照被冻结",
		},
	)

	fmt.Println(fmt.Sprintf("ShopSave Status result: %+v", result))
}

func SaveLastSyncTime() {
	shop := sanzhizhouComponent.Shop{}
	result := shop.SaveLastSyncTime(
		mainConfig.UserToken,
		&sanzhizhouComponentConfig.ShopSaveLastSyncTimeQuery{
			ShopId: 24,
		},
	)

	fmt.Println(fmt.Sprintf("SaveLastSyncTime result: %+v", result))
}

func GetPublishShopIdList() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.GetPublishShopIdList(
		mainConfig.UserToken,
		7,
	)

	fmt.Println(fmt.Sprintf("SaveLastSyncTime result: %+v", result))
	fmt.Println("err:", err)
}
