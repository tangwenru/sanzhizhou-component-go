package sanzhizhouComponent

import (
	"errors"
	"fmt"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Vip struct {
}

type VipList struct {
	Id                    int64   `json:"id"`
	DealerId              int64   `json:"dealerId"`
	Title                 string  `json:"title"`
	Price                 float64 `json:"price"`
	Month                 float64 `json:"month"`
	ProductType           string  `json:"productType"`
	ProductTypeOrderIndex int64   `json:"productTypeOrderIndex"`
	MonthTitle            string  `json:"monthTitle"`
	MarketPrice           float64 `json:"marketPrice"`
	RatioDownPercent1     float64 `json:"ratioDownPercent1"`
	RatioDownPercent2     float64 `json:"ratioDownPercent2"`
	Alt                   string  `json:"alt"`
	VipLevel              int64   `json:"vipLevel"`
	PropertyAmount        string  `json:"propertyAmount"` // 最大值
	Recommend             int64   `json:"recommend"`
	ShowContent           string  `json:"showContent"`
	CanBuy                bool    `json:"canBuy"`
	Enabled               bool    `json:"enabled"`
}

type VipDetailQuery struct {
	ProductType string `json:"productType"`
	Domain      string `json:"domain"`
}

type VipListResult struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []VipList `json:"data"`
}

func init() {

}

func (this *Vip) ListByProductType(userToken string, productType string, domain string) (error, *[]VipList) {
	query := VipDetailQuery{
		ProductType: productType,
		Domain:      domain,
	}
	vipListResult := VipListResult{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "vip/listByProductType", &query, &vipListResult)

	empty := make([]VipList, 0)

	if err != nil {
		fmt.Println("Vip List By Product Type err:", string(bytesResult), err)
		return err, &empty
	}

	if !vipListResult.Success {
		return errors.New(vipListResult.Message), &empty
	}

	return err, &vipListResult.Data
}
