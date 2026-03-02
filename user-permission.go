package sanzhizhouComponent

import (
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserPermission struct {
}

var UserPermissionKindNameDict = sanzhizhouComponentConfig.UserPermissionKindNameDict{
	HasShareUrl: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "has-share-url",
		ValueType: "bool",
	},
	ImageTranslateAmount: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "image-translate-amount",
		ValueType: "int64",
	},
	GoodsImageTranslateAmount: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "goods-image-translate-amount",
		ValueType: "int64",
	},
	AiCoinAmount: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "ai-coin-amount",
		ValueType: "int64",
	},
}

func (this *UserPermission) Detail(userToken string) (*sanzhizhouComponentConfig.UserPermissionDetailResult, *sanzhizhouComponentConfig.UserPermissionDetail) {
	userDetailResult := sanzhizhouComponentConfig.UserPermissionDetailResult{}
	query := map[string]string{}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "userPermission/detail", &query, &userDetailResult)

	userDetail := sanzhizhouComponentConfig.UserPermissionDetail{}
	if err != nil {
		fmt.Println("userPermission detail err:", err)
		userDetailResult.Message = err.Error()
		return &userDetailResult, &userDetail
	}

	if !userDetailResult.Success {
		return &userDetailResult, &userDetail
	}

	return &userDetailResult, &userDetailResult.Data
}

// 扣费
func (this *UserPermission) Deduct(
	userToken string,
	taskId int64,
	aiTaskId string,
	permissionKind,
	taskKindName string,
	speedTrueAiCoin int,
	isDeduct bool, // 扣除，还是返还
) *sanzhizhouComponentConfig.UserPermissionDeductResult {
	userDeductResult := sanzhizhouComponentConfig.UserPermissionDeductResult{}
	query := map[string]interface{}{
		"taskId":          taskId,
		"aiTaskId":        aiTaskId,
		"permissionKind":  permissionKind,
		"taskKindName":    taskKindName,
		"speedTrueAiCoin": speedTrueAiCoin,
		"isDeduct":        isDeduct,
	}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "userPermission/deduct", &query, &userDeductResult)

	//userDetail := UserPermissionDetail{}
	if err != nil {
		fmt.Println("userPermission ToDeduction err:", err)
		userDeductResult.Message = err.Error()
		return &userDeductResult
	}

	return &userDeductResult
}
