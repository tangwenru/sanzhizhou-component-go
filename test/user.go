package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//UserDict()
	UserGetDomesticWarehouseId()
}

func UserDict() {
	user := sanzhizhouComponent.User{}
	err, result := user.Dict(
		mainConfig.StaffToken,
		&[]int64{
			30, 40, 41, 42, 43, 44, 45,
		},
	)

	fmt.Println(fmt.Sprintf("UserSaveStatus result: %+v", result))
	fmt.Println(fmt.Sprintf("UserSaveStatus err: %+v", err))
}

func UserGetDomesticWarehouseId() {
	user := sanzhizhouComponent.User{}
	result, err := user.GetDomesticWarehouseId(
		mainConfig.UserToken,
	)

	fmt.Println(fmt.Sprintf("UserGet DomesticWarehouseId result: %+v", result))
	fmt.Println(fmt.Sprintf("UserGet DomesticWarehouseId err: %+v", err))
}
