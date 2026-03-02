package sanzhizhouComponentConfig

type UserPermissionDetailResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Code    string               `json:"code"`
	Data    UserPermissionDetail `json:"data"`
}

type UserPermissionDetail struct {
	Id                  int64 `json:"id"`
	HasShareUrl         bool  `json:"hasShareUrl"`
	ImageTranslate      int   `json:"ImageTranslate"`
	GoodsImageTranslate int   `json:"goodsImageTranslate"` // 电商质量翻译
	AiCoin              int   `json:"AiCoin"`
}

type UserPermissionDeductResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type UserPermissionKindNameDict struct {
	HasShareUrl         UserPermissionKindNameDictItem
	ImageTranslate      UserPermissionKindNameDictItem
	GoodsImageTranslate UserPermissionKindNameDictItem // 电商质量翻译
	AiCoin              UserPermissionKindNameDictItem
}

type UserPermissionKindNameDictItem struct {
	Kind      string `json:"kind"`
	ValueType string `json:"valueType"` //  float64 | int64 | string | bool
}
