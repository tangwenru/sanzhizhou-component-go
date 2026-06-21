package sanzhizhouComponent

import (
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

func (this *Shop) TransferShop(
	userToken string,
	query *sanzhizhouComponentConfig.ShopTransferShopQuery,
) (*sanzhizhouComponentConfig.ShopTransferShopResult, *sanzhizhouComponentConfig.ShopTransferShopResultData) {
	saveResult := sanzhizhouComponentConfig.ShopTransferShopResult{}
	empty := sanzhizhouComponentConfig.ShopTransferShopResultData{}

	_, err := sanzhizhouComponentLib.MainSystemPost(
		userToken,
		"shop/transferShop",
		&query,
		&saveResult,
	)

	if err != nil {
		return &saveResult, &empty
	}

	if !saveResult.Success {
		return &saveResult, &empty
	}

	return &saveResult, &saveResult.Data
}
