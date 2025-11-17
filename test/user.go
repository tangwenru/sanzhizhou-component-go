package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	UserDict()
}

func UserDict() {
	user := sanzhizhouComponent.User{}
	err, result := user.Dict(
		mainConfig.StaffToken,
		&[]int64{
			1, 2, 3, 4, 5, 20, 21,
		},
	)

	fmt.Println(fmt.Sprintf("UserSaveStatus result: %+v", result))
	fmt.Println(fmt.Sprintf("UserSaveStatus err: %+v", err))
}
