package sanzhizhouComponentConfig

type UserDealerAllUserIdListResult struct {
	Success bool                      `json:"success"`
	Code    string                    `json:"code"`
	Message string                    `json:"message"`
	Data    []UserDealerAllUserIdList `json:"data"`
}

type UserDealerAllUserIdList int64
