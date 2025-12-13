package sanzhizhouComponentConfig

import "time"

type ShopGetLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type GetLastSyncTimeResult struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Data    GetLastSyncTime `json:"data"`
	Code    string          `json:"code"`
}

type GetLastSyncTime struct {
	OrderLastSyncTime   time.Time `json:"orderLastSyncTime"`
	ProductLastSyncTime time.Time `json:"productLastSyncTime"`
}

type ShopBaseInfo struct {
	Id             int64          `json:"id"`
	Name           string         `json:"name"`
	NameRemark     string         `json:"nameRemark"`
	GroupName      string         `json:"groupName"`
	Currency       string         `json:"currency"`
	PublishSetting string         `json:"publishSetting"` // 可能为空，依赖接口参数
	ApiInfo        ShopApiKeyInfo `json:"apiInfo"`        // 可能为空，依赖接口参数
	IsBan          bool           `json:"isBan"`
	Enabled        bool           `json:"enabled"`
	StatusMessage  string         `json:"statusMessage"`
}

type ShopApiKeyInfo map[string]string

type ShopDictQuery struct {
	ShopIdList []int64 `json:"shopIdList"`
	ShowExtend bool    `json:"showExtend"`
}

type ShopDictPostQuery struct {
}

type ShopDictResult struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message,omitempty"`
	Data    map[int64]ShopBaseInfo `json:"data"`
	Code    string                 `json:"code"`
}

type ShopListResultQuery struct {
	CommercePlatformId int64 `json:"commercePlatformId"`
	ShowAll            bool  `json:"showExtend"`
}
type ShopListResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Data    ShopListData `json:"data"`
	Code    string       `json:"code"`
}

type ShopListData struct {
	List       []ShopList `json:"list"`
	Pagination struct {
		Current  int   `json:"current"`
		PageSize int   `json:"pageSize"`
		Amount   int64 `json:"amount"`
	} `json:"pagination"`
}

type ShopList struct {
	Id                 int64  `json:"id"`
	UserId             int64  `json:"userId"`
	PlatformShopId     string `json:"platformShopId"`
	ShopName           string `json:"shopName"`
	ShopAvatarUrl      string `json:"shopAvatarUrl"`
	ShopNameRemark     string `json:"shopNameRemark"`
	IsLoginValid       bool   `json:"isLoginValid"`
	ShopOrganizationId int64  `json:"shopOrganizationId"`
	Enabled            bool   `json:"enabled"`
	IsBan              bool   `json:"isBan"`
	IsPublishSelect    bool   `json:"isPublishSelect"`
	IsCloud            bool   `json:"isCloud"`
	Created            int64  `json:"created"`
	Updated            int64  `json:"updated"`
	StatusMessage      string `json:"statusMessage"`

	ChromeProfileName string `json:"chromeProfileName"`

	CommercePlatformInfo CommercePlatformBaseInfo `json:"commercePlatformInfo"`
	GroupInfo            ShopGroupBaseInfo        `json:"groupInfo"`

	Currency          string            `json:"currency"`
	OrderLastSyncTime time.Time         `json:"orderLastSyncTime"`
	ApiInfo           map[string]string `json:"apiInfo"`
}

type ShopDetail struct {
	Id                 int64  `json:"id"`
	PlatformShopId     string `json:"platformShopId"`
	ShopName           string `json:"shopName"`
	ShopAvatarUrl      string `json:"shopAvatarUrl"`
	ShopNameRemark     string `json:"shopNameRemark"`
	IsLoginValid       bool   `json:"isLoginValid"`
	ShopOrganizationId int64  `json:"shopOrganizationId"`
	Enabled            bool   `json:"enabled"`
	IsBan              bool   `json:"isBan"`
	IsPublishSelect    bool   `json:"isPublishSelect"`
	IsCloud            bool   `json:"isCloud"`
	Created            int64  `json:"created"`
	Updated            int64  `json:"updated"`
	StatusMessage      string `json:"statusMessage"`

	ChromeProfileName string `json:"chromeProfileName"`

	CommercePlatformInfo CommercePlatformBaseInfo `json:"commercePlatformInfo"`
	GroupInfo            ShopGroupBaseInfo        `json:"groupInfo"`

	Currency          string            `json:"currency"`
	OrderLastSyncTime time.Time         `json:"orderLastSyncTime"`
	ApiInfo           map[string]string `json:"apiInfo"`
	PublishSetting    string            `json:"publishSetting"`
}
type ShopDetailQuery struct {
	Id int64 `json:"id"`
}

type ShopDetailResult struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    ShopDetail `json:"data"`
	Code    string     `json:"code"`
}

// //
type ShopByDomesticWarehouseIdQuery struct {
	DomesticWarehouseId int64 `json:"domesticWarehouseId"`
}

type ListByDomesticWarehouseIdResult struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []int64 `json:"data"`
	Code    string  `json:"code"`
}

type ShopSaveStatusQuery struct {
	ShopId        int64  `json:"shopId"`
	IsBan         string `json:"isBan"`   // '' | 1 | 0
	Enabled       string `json:"enabled"` // '' | 1 | 0
	StatusMessage string `json:"statusMessage"`
}
type ShopSaveStatusQueryResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	//Data    []int64 `json:"data"`
	Code string `json:"code"`
}

type ShopSaveLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type ShopSaveLastSyncTimeResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code"`
}

type ShopGetPublishShopIdListQuery struct {
	CommercePlatformId int64
}
type ShopGetPublishShopIdListResult struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []int64 `json:"data"`
	Code    string  `json:"code"`
}
