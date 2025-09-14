package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
)

func main() {
	//GetOrderLastSyncTime()
	ShopDict()
}

func GetOrderLastSyncTime() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.GetOrderLastSyncTime(
		"a1217689d4cb45390feb02d4125f44f8",
		20,
	)

	fmt.Println("GetOrderLastSyncTime result:", result)
	fmt.Println("GetOrderLastSyncTime err:", err)
}

func ShopDict() {
	shop := sanzhizhouComponent.Shop{}
	result, err := shop.Dict(
		"a1217689d4cb45390feb02d4125f44f8",
		&[]int64{5, 6, 7, 8, 9, 19, 20},
		true,
	)

	fmt.Println(fmt.Sprintf("ShopDict result: %+v", result))
	fmt.Println("ShopDict err:", err)
}
