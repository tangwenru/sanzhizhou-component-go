package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//UserPermissionDetail()
	//UserPermissionDeduct()
	UserPermissionDeductRefund()
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
		1,                        //  int64,
		"image_translate_amount", // permissionKind string,
		true,                     // isDeduct bool,
	)

	fmt.Println(fmt.Sprintf("UserPermission Deduct result: %+v", result))
}

func UserPermissionDeductRefund() {
	userPermission := sanzhizhouComponent.UserPermission{}
	result := userPermission.Deduct(
		mainConfig.UserToken,
		1,                        //  int64,
		"image_translate_amount", // permissionKind string,
		false,                    // isDeduct bool,
	)

	fmt.Println(fmt.Sprintf("UserPermission Deduct result: %+v", result))
}
