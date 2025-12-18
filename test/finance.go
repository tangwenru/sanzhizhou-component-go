package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	//Transfer()
	DetailByTarget()
}

func Transfer() {
	finance := sanzhizhouComponent.Finance{}
	err := finance.Transfer(
		mainConfig.StaffToken,
		&sanzhizhouComponentConfig.FinanceTransferQuery{
			ToUserId:            1,
			Money:               -3,
			Type:                "domestic-warehouse-fee",
			Name:                "转运费",
			ExtInfo:             "",
			TargetId:            1,
			ToUserDownUserId:    0,
			ToUserDownUserLevel: 0,
		},
	)

	fmt.Println("finance err:", err)
}

func DetailByTarget() {
	finance := sanzhizhouComponent.Finance{}
	_, result := finance.DetailByTarget(
		mainConfig.UserToken,
		"domestic-warehouse-fee",
		3083,
	)

	fmt.Println("finance DetailByTarget:", result)
}
