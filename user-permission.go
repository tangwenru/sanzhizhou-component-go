package sanzhizhouComponent

import (
	"fmt"
	"reflect"

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
		Kind:      "image-translate",
		ValueType: "int64",
	},
	GoodsImageTranslateAmount: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "goods-image-translate",
		ValueType: "int64",
	},
	AiCoinAmount: sanzhizhouComponentConfig.UserPermissionKindNameDictItem{
		Kind:      "ai-coin",
		ValueType: "int64",
	},
}

// 性能
var UserPermissionValueTypeDict = map[string]string{}

// 反射有稍微 性能问题，缓存
func (this *UserPermission) GetValueType(kind string) (_result string) {
	cacheItem := UserPermissionValueTypeDict[kind]
	if cacheItem != "" {
		return cacheItem
	}

	defer func() {
		// 存起来
		if _result != "" {
			UserPermissionValueTypeDict[kind] = _result
		}
	}()

	// 获取全局变量的反射值
	v := reflect.ValueOf(UserPermissionKindNameDict)
	// 如果传入的是指针，需要先取指针指向的元素
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// 确保是结构体类型
	if v.Kind() != reflect.Struct {
		return ""
	}

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// 检查字段类型是否为 UserPermissionKindNameDictItem
		// 假设目标类型为 sanzhizhouComponentConfig.UserPermissionKindNameDictItem
		itemType := reflect.TypeOf(sanzhizhouComponentConfig.UserPermissionKindNameDictItem{})
		if field.Type() != itemType {
			continue
		}

		// 获取 Kind 字段的值
		kindField := field.FieldByName("Kind")
		if !kindField.IsValid() || kindField.Kind() != reflect.String {
			continue
		}
		if kindField.String() == kind {
			// 获取 ValueType 字段的值
			valueTypeField := field.FieldByName("ValueType")
			if valueTypeField.IsValid() && valueTypeField.Kind() == reflect.String {
				_result = valueTypeField.String()
				return _result
			}
		}
	}
	return ""
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
