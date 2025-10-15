package main

import (
	"fmt"

	sanzhizhouComponent "github.com/tangwenru/sanzhizhou-component-go"
	sanzhizhouComponentConfig "github.com/tangwenru/sanzhizhou-component-go/config"
	mainConfig "github.com/tangwenru/sanzhizhou-component-go/test/config"
)

func main() {
	Transfer()
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
