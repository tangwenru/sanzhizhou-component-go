package sanzhizhouComponent

import (
	"fmt"

	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserPermission struct {
}

type UserPermissionDetailResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Code    string               `json:"code"`
	Data    UserPermissionDetail `json:"data"`
}

type UserPermissionDetail struct {
	Id                        int64 `json:"id"`
	HasShareUrl               bool  `json:"hasShareUrl"`
	ImageTranslateAmount      int   `json:"imageTranslateAmount"`
	GoodsImageTranslateAmount int   `json:"goodsImageTranslateAmount"` // 电商质量翻译
	TextTranslateAmount       int   `json:"textTranslateAmount"`
}

type UserPermissionDeductResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

var UserPermissionKindNameDict = struct {
	HasShareUrl               string
	ImageTranslateAmount      string
	GoodsImageTranslateAmount string // 电商质量翻译
	AiCoinAmount              string
}{
	HasShareUrl:               "has-share-url",
	ImageTranslateAmount:      "image-translate-amount",
	GoodsImageTranslateAmount: "goods-image-translate-amount",
	AiCoinAmount:              "ai-coin-amount",
}

func (this *UserPermission) Detail(userToken string) (*UserPermissionDetailResult, *UserPermissionDetail) {
	userDetailResult := UserPermissionDetailResult{}
	query := map[string]string{}
	_, err := sanzhizhouComponentLib.MainSystem(userToken, "userPermission/detail", &query, &userDetailResult)

	userDetail := UserPermissionDetail{}
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
	permissionKind string,
	isDeduct bool, // 扣除，还是返还
) *UserPermissionDeductResult {
	userDeductResult := UserPermissionDeductResult{}
	query := map[string]interface{}{
		"taskId":         taskId,
		"aiTaskId":       aiTaskId,
		"permissionKind": permissionKind,
		"isDeduct":       isDeduct,
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
