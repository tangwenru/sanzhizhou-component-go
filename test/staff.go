package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//GetLastSyncTime()
	//ShopList()
	StaffInfo()
}

func StaffInfo() {
	staff := sanzhizhouComponent.Staff{}
	result, err := staff.Info(
		mainConfig.StaffToken,
	)

	fmt.Println(fmt.Sprintf("staffDetail result: %+v", result))
	fmt.Println("staffDetail err:", err)
}
