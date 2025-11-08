package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	MessageCreate()
}

func MessageCreate() {
	message := sanzhizhouComponent.Message{}
	err := message.Create(
		mainConfig.UserToken,
		&sanzhizhouComponentConfig.MessageCreateQuery{
			CommercePlatformId: 7,
			Title:              "标题 2",
			Introduction:       "Introduction 1",
			Content:            "Content 1",
			TypeStatus:         "error",
		},
	)

	fmt.Println("MessageCreate err:", err)
}
