package sanzhizhouComponentConfig

type UserPermissionDetailResult struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Code    string               `json:"code"`
	Data    UserPermissionDetail `json:"data"`
}

type UserPermissionDetail struct {
	Id                  int64 `json:"id"`
	HasShareUrl         bool  `json:"has-share-url"`
	ImageTranslate      int   `json:"image-translate"`
	GoodsImageTranslate int   `json:"goods-image-translate"` // 电商质量翻译
	AiCoin              int   `json:"ai-coin"`
}

type UserPermissionDeductResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type UserPermissionKindNameDict struct {
	HasShareUrl         UserPermissionKindNameDictItem `json:"has-share-url"`
	ImageTranslate      UserPermissionKindNameDictItem `json:"image-translate"`
	GoodsImageTranslate UserPermissionKindNameDictItem `json:"goods-image-translate"` // 电商质量翻译
	AiCoin              UserPermissionKindNameDictItem `json:"ai-coin"`
}

type UserPermissionKindNameDictItem struct {
	Kind      string `json:"kind"`
	ValueType string `json:"valueType"` //  float64 | int64 | string | bool
}
