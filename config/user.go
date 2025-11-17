package sanzhizhouComponentConfig

type UserDict map[int64]UserDictItem

type UserDictItem struct {
	Id          int64   `json:"id"`
	AccountName string  `json:"accountName"`
	Nickname    string  `json:"nickname"`
	AvatarUrl   string  `json:"avatarUrl"`
	DealerId    int64   `json:"dealerId"`
	Balance     float64 `json:"balance"`
}

type UserDictResult struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Code    string   `json:"code"`
	Data    UserDict `json:"data"`
}
