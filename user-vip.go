package sanzhizhouComponent

import (
	"fmt"

	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserVip struct {
}

type UserVipDetail struct {
	Id                int64
	UserId            int64
	ProductType       string
	RatioDownPercent1 float64
	RatioDownPercent2 float64
	VipLevel          int64
	IsVip             bool
	//VipDetail         typeChatgpt.VipDetail `json:"vipDetail"`
	Expired int64
	Created int64
	Updated int64
}

type UserVipDetailQuery struct {
	ProductType string `json:"productType"`
}

func init() {

}

func (this *UserVip) Detail(userToken, productType string) (error, *UserVipDetail) {
	query := UserVipDetailQuery{
		ProductType: productType,
	}
	userVipDetail := UserVipDetail{}

	bytesResult, err := sanzhizhouComponentLib.MainSystem(userToken, "userVip/detail", &query, &userVipDetail)

	if err != nil {
		fmt.Println("UserVip Detail:", string(bytesResult))
		//json.Unmarshal(bytesResult, &userVipDetail)
	}

	return err, &userVipDetail
}
