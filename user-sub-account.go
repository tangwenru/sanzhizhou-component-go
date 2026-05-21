package sanzhizhouComponent

import (
	"fmt"
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type UserSubAccount struct {
}

func (this *UserSubAccount) CheckPowerUserIdCode(
	userToken,
	subUserIdCode string,
) (
	*sanzhizhouComponentConfig.UserSubAccountCheckPowerUserIdCodeResult,
	*sanzhizhouComponentConfig.UserSubAccountCheckPowerUserIdCode,
) {
	userDetailResult := sanzhizhouComponentConfig.UserSubAccountCheckPowerUserIdCodeResult{}
	query := map[string]string{
		"subUserIdCode": subUserIdCode,
	}
	_, err := sanzhizhouComponentLib.MainSystem(
		userToken,
		"userSubAccount/checkPowerUserIdCode",
		&query,
		&userDetailResult,
	)

	userDetail := sanzhizhouComponentConfig.UserSubAccountCheckPowerUserIdCode{}
	if err != nil {
		fmt.Println("userSubAccount-detail err:", err)
		userDetailResult.Message = err.Error()
		return &userDetailResult, &userDetail
	}

	if !userDetailResult.Success {
		return &userDetailResult, &userDetail
	}

	return &userDetailResult, &userDetailResult.Data
}
