package sanzhizhouComponentConfig

type MessageCreateQuery struct {
	CommercePlatformId int64  `json:"commercePlatformId"`
	Title              string `json:"title"`
	Introduction       string `json:"introduction"`
	Content            string `json:"content"`
}

type MessageCreateResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    bool   `json:"data"`
}
