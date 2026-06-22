package sanzhizhouComponentConfig

type UserDealerAllUserIdListResult struct {
	Success bool    `json:"success"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    []int64 `json:"data"`
}
