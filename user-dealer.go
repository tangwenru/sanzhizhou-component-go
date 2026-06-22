package sanzhizhouComponent

import (
	"fmt"

	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserDealer struct {
}

func (this *UserDealer) AllUserIdList(dealerToken string) (
	*sanzhizhouComponentConfig.UserDealerAllUserIdListResult,
	*[]int64,
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
