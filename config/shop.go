package sanzhizhouComponentConfig

import "time"

type ShopGetLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type GetLastSyncTimeResult struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Data    GetLastSyncTime `json:"data"`
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
}

type ShopListResultQuery struct {
	CommercePlatformId int64 `json:"commercePlatformId"`
	ShowAll            bool  `json:"showExtend"`
}
type ShopListResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Data    ShopListData `json:"data"`
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
	IsBan              bool   `json:"isBan"`
	IsPublishSelect    bool   `json:"isPublishSelect"`
	IsCloud            bool   `json:"isCloud"`
	Created            int64  `json:"created"`
	Updated            int64  `json:"updated"`

	ChromeProfileName string `json:"chromeProfileName"`

	CommercePlatformInfo CommercePlatformBaseInfo `json:"commercePlatformInfo"`
	GroupInfo            ShopGroupBaseInfo        `json:"groupInfo"`

	Currency          string            `json:"currency"`
	OrderLastSyncTime int64             `json:"orderLastSyncTime"`
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
	IsBan              bool   `json:"isBan"`
	IsPublishSelect    bool   `json:"isPublishSelect"`
	IsCloud            bool   `json:"isCloud"`
	Created            int64  `json:"created"`
	Updated            int64  `json:"updated"`

	ChromeProfileName string `json:"chromeProfileName"`

	CommercePlatformInfo CommercePlatformBaseInfo `json:"commercePlatformInfo"`
	GroupInfo            ShopGroupBaseInfo        `json:"groupInfo"`

	Currency          string            `json:"currency"`
	OrderLastSyncTime int64             `json:"orderLastSyncTime"`
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
}

// //
type ShopByDomesticWarehouseIdQuery struct {
	DomesticWarehouseId int64 `json:"domesticWarehouseId"`
}

type ListByDomesticWarehouseIdResult struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []int64 `json:"data"`
}

type ShopSaveStatusQuery struct {
	ShopId  int64  `json:"shopId"`
	IsBan   string `json:"isBan"`   // '' | 1 | 0
	Enabled string `json:"enabled"` // '' | 1 | 0
}
type ShopSaveStatusQueryResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	//Data    []int64 `json:"data"`
}

type ShopSaveLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type ShopSaveLastSyncTimeResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
