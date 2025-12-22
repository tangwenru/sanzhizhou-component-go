package sanzhizhouComponent

import (
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
) sanzhizhouComponentConfig.FinanceTransferResult {
	shopDictResult := sanzhizhouComponentConfig.FinanceTransferResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(staffToken, "finance/transfer", &query, &shopDictResult)

	if err != nil {
		fmt.Println("Finance Transfer err:", string(bytesResult), err)
	}

	return shopDictResult
}

func (this *Finance) DetailByTarget(
	userToken,
	financeType string,
	targetId int64,
) (*sanzhizhouComponentConfig.FinanceDetail, *sanzhizhouComponentConfig.FinanceDetailResult) {
	result := sanzhizhouComponentConfig.FinanceDetailResult{}

	query := sanzhizhouComponentConfig.FinanceDetailByTargetQuery{
		FinanceType: financeType,
		TargetId:    targetId,
	}

	empty := sanzhizhouComponentConfig.FinanceDetail{}
	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "finance/detailByTarget", &query, &result)

	if err != nil {
		fmt.Println("Finance-Detail-err:", string(bytesResult), err)
		return &empty, &result
	}

	if !result.Success {
		return &empty, &result
	}

	return &result.Data, &result
}
