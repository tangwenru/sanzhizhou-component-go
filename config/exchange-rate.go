package sanzhizhouComponentConfig

type ExchangeRateListResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message,omitempty"`
	Data    ExchangeRateListData `json:"data"`
}

type ExchangeRateListData struct {
	List       []ExchangeRateList `json:"list"`
	Pagination struct {
		Current  int   `json:"current"`
		PageSize int   `json:"pageSize"`
		Amount   int64 `json:"amount"`
	} `json:"pagination"`
}

type ExchangeRateList struct {
	Id       int64   `json:"id"`
	Currency string  `json:"currency"`
	Label    string  `json:"label"`
	IconUrl  string  `json:"iconUrl"`
	Rate     float32 `json:"rate"`
	Updated  int64   `json:"updated"`
}

type ExchangeRateListQuery struct {
	CurrencyList []string `json:"currencyList"`
}
