package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	UserDealerAllUserIdList()
}

func UserDealerAllUserIdList() {
	userDealer := sanzhizhouComponent.UserDealer{}
	result, _ := userDealer.AllUserIdList(
		mainConfig.UserToken,
	)

	fmt.Println(fmt.Sprintf("UserDealerAllUserIdList result: %+v", result))
}
