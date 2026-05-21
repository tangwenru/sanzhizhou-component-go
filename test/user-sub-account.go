package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	UserSubAccountCheckPowerUserIdCode()
}

func UserSubAccountCheckPowerUserIdCode() {
	userSubAccount := sanzhizhouComponent.UserSubAccount{}
	result, id := userSubAccount.CheckPowerUserIdCode(

		mainConfig.UserToken,
		mainConfig.SubUserIdCode,
	)

	fmt.Println(fmt.Sprintf("UserSubAccount GetValueType: %+v, %d", result, id))
}
