package sanzhizhouComponentConfig

type UserPermissionDetailResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Code    string               `json:"code"`
	Data    UserPermissionDetail `json:"data"`
}

type UserPermissionDetail struct {
	Id                        int64 `json:"id"`
	HasShareUrl               bool  `json:"hasShareUrl"`
	ImageTranslateAmount      int   `json:"imageTranslateAmount"`
	GoodsImageTranslateAmount int   `json:"goodsImageTranslateAmount"` // 电商质量翻译
	AiCoinAmount              int   `json:"aiCoinAmount"`
}

type UserPermissionDeductResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type UserPermissionKindNameDict struct {
	HasShareUrl               UserPermissionKindNameDictItem
	ImageTranslateAmount      UserPermissionKindNameDictItem
	GoodsImageTranslateAmount UserPermissionKindNameDictItem // 电商质量翻译
	AiCoinAmount              UserPermissionKindNameDictItem
}

type UserPermissionKindNameDictItem struct {
	Kind      string `json:"kind"`
	ValueType string `json:"valueType"` //  float64 | int64 | string | bool
}
