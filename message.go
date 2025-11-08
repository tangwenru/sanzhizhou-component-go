package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Message struct {
}

func init() {

}

func (this *Message) Create(
	userToken string,
	query *sanzhizhouComponentConfig.MessageCreateQuery,
) error {

	messageCreateResult := sanzhizhouComponentConfig.MessageCreateResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(
		userToken,
		"message/create",
		&query,
		&messageCreateResult,
	)

	if err != nil {
		fmt.Println("messageCreateResult err:", string(bytesResult), err)
	}

	if !messageCreateResult.Success {
		return errors.New(messageCreateResult.Message)
	}

	return nil
}
