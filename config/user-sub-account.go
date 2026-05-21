package sanzhizhouComponentConfig

type UserSubAccountCheckPowerUserIdCodeResult struct {
	Success bool                               `json:"success"`
	Message string                             `json:"message"`
	Code    string                             `json:"code"`
	Data    UserSubAccountCheckPowerUserIdCode `json:"data"`
}

type UserSubAccountCheckPowerUserIdCode struct {
	UserId int64 `json:"userId"`
}
