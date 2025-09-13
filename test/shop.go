package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
)

func main() {
	GetOrderLastSyncTime()
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
