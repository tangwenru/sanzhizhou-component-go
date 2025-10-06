package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
)

func main() {
	//GetLastSyncTime()
	//ShopList()
	StaffInfo()
}

func StaffInfo() {
	staff := sanzhizhouComponent.Staff{}
	result, err := staff.Info(
		"e8c80437ee5fc97a",
	)

	fmt.Println(fmt.Sprintf("staffDetail result: %+v", result))
	fmt.Println("staffDetail err:", err)
}
