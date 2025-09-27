package sanzhizhouComponent

import (
	"errors"
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type ExchangeRate struct {
}

func init() {

}

func (this *ExchangeRate) List(
	userToken string,
	query *sanzhizhouComponentConfig.ExchangeRateListQuery,
) (*sanzhizhouComponentConfig.ExchangeRateListData, error) {
	exchangeRateDictResult := sanzhizhouComponentConfig.ExchangeRateListResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "exchangeRate/list", &query, &exchangeRateDictResult)

	if err != nil {
		fmt.Println("ExchangeRate List err:", string(bytesResult), err)
	}

	if !exchangeRateDictResult.Success {
		return &sanzhizhouComponentConfig.ExchangeRateListData{}, errors.New(exchangeRateDictResult.Message)
	}

	return &exchangeRateDictResult.Data, nil
}
