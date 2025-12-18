package sanzhizhouComponentConfig

type FinanceTransferQuery struct {
	ToUserId            int64   `json:"toUserId"`
	Money               float64 `json:"money"`
	Type                string  `json:"type"`
	Name                string  `json:"name"`
	ExtInfo             string  `json:"extInfo"`
	TargetId            int64   `json:"targetId"`
	ToUserDownUserId    int64   `json:"toUserDownUserId"`
	ToUserDownUserLevel int     `json:"toUserDownUserLevel"`
}

type FinanceTransferResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type FinanceDetailByTargetQuery struct {
	FinanceType string `json:"financeType"`
	TargetId    int64  `json:"targetId"`
}
type FinanceDetailByTargetResult struct {
	Success bool                  `json:"success"`
	Message string                `json:"message,omitempty"`
	Code    string                `json:"code,omitempty"`
	Data    FinanceDetailByTarget `json:"data"`
}

type FinanceDetailByTarget struct {
}
