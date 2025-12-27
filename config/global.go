package sanzhizhouComponentConfig

type Pagination struct {
	Current  int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
	Total    int64 `json:"total"`
}
