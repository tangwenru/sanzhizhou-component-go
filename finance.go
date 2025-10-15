package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Finance struct {
}

func init() {

}

func (this *Finance) Transfer(
	staffToken string,
	query *sanzhizhouComponentConfig.FinanceTransferQuery,
) error {
	shopDictResult := sanzhizhouComponentConfig.FinanceTransferResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(staffToken, "finance/transfer", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Finance Transfer err:", string(bytesResult), err)
	}

	if !shopDictResult.Success {
		return errors.New(shopDictResult.Message)
	}

	return nil
}
