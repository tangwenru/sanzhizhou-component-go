package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//UserPermissionDetail()
	//UserPermissionDeduct()
	//UserPermissionDeductRefund()
	UserPermissionGetValueType()
}

func UserPermissionDetail() {
	userPermission := sanzhizhouComponent.UserPermission{}
	err, result := userPermission.Detail(
		mainConfig.UserToken,
	)

	fmt.Println(fmt.Sprintf("UserPermissionSaveStatus result: %+v", result))
	fmt.Println(fmt.Sprintf("UserPermissionSaveStatus err: %+v", err))
}

func UserPermissionDeduct() {
	userPermission := sanzhizhouComponent.UserPermission{}
	result := userPermission.Deduct(
		mainConfig.UserToken,
		1, //  int64,
		"",
		"image-translate-amount", // permissionKind string,
		"",
		1,
		true, // isDeduct bool,
	)

	fmt.Println(fmt.Sprintf("UserPermission Deduct result: %+v", result))
}

func UserPermissionDeductRefund() {
	userPermission := sanzhizhouComponent.UserPermission{}
	result := userPermission.Deduct(
		mainConfig.UserToken,
		1, //  int64,
		"",
		"image-translate-amount", // permissionKind string,
		"",
		1,
		false, // isDeduct bool,
	)

	fmt.Println(fmt.Sprintf("UserPermission Deduct result: %+v", result))
}

func UserPermissionGetValueType() {
	userPermission := sanzhizhouComponent.UserPermission{}
	result := userPermission.GetValueType(
		sanzhizhouComponent.UserPermissionKindNameDict.AiCoinAmount.Kind,
	)

	fmt.Println(fmt.Sprintf("UserPermission GetValueType: %+v", result))
}
