package sanzhizhouComponent

import (
	"errors"
	"fmt"
	//sanzhizhouComponentLib "sanzhizhouComponent/lib"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
)

type Staff struct {
}

type StaffDetailResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    StaffDetail `json:"data"`
}

type StaffDetail struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Mobile              string `json:"mobile"`
	Gender              int    `json:"gender"`
	AvatarUrl           string `json:"avatarUrl"`
	Role                string `json:"role"`
	DomesticWarehouseId int64  `json:"domesticWarehouseId"`
}

func (this *Staff) Info(staffToken string) (error, *StaffDetail) {
	staffDetailResult := StaffDetailResult{}
	query := map[string]string{}
	_, err := sanzhizhouComponentLib.MainSystem(staffToken, "staff/info", &query, &staffDetailResult)

	staffDetail := StaffDetail{}
	if err != nil {
		fmt.Println("shiPinLv staff info err:", err)
		return err, &staffDetail
	}

	if !staffDetailResult.Success {
		return errors.New(staffDetailResult.Message), &staffDetail
	}

	return err, &staffDetailResult.Data
}
