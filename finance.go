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

func (this *Finance) DetailByTarget(
	userToken,
	financeType string,
	targetId int64,
) (*sanzhizhouComponentConfig.FinanceDetailByTarget, *sanzhizhouComponentConfig.FinanceDetailByTargetResult) {
	result := sanzhizhouComponentConfig.FinanceDetailByTargetResult{}

	query := sanzhizhouComponentConfig.FinanceDetailByTargetQuery{
		FinanceType: financeType,
		TargetId:    targetId,
	}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "finance/detail", &query, &result)

	if err != nil {
		fmt.Println("Finance-Detail-err:", string(bytesResult), err)
	}

	if !result.Success {
		return nil, &result
	}

	return &result.Data, &result
}
