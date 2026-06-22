package sanzhizhouComponent

import (
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserDealer struct {
}

type UserDealerDetailResult struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Code    string          `json:"code"`
	Data    []AllUserIdList `json:"data"`
}

type AllUserIdList []int64

func (this *UserDealer) AllUserIdList(dealerToken string) (
	*sanzhizhouComponentConfig.UserDealerAllUserIdListResult,
	*[]sanzhizhouComponentConfig.UserDealerAllUserIdList,
) {
	userDetailResult := sanzhizhouComponentConfig.UserDealerAllUserIdListResult{}

	query := map[string]string{}
	_, err := sanzhizhouComponentLib.MainSystem(
		dealerToken,
		"userDealer/allUserIdList",
		&query,
		&userDetailResult,
	)

	if err != nil {
		fmt.Println("userDealer-AllUserIdList err:", err)
		return &userDetailResult, &userDetailResult.Data
	}

	return &userDetailResult, &userDetailResult.Data
}
