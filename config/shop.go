package sanzhizhouComponentConfig

type ShopGetLastSyncTimeQuery struct {
	ShopId int64 `json:"shopId"`
}

type GetLastSyncTimeResult struct {
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	Data    GetLastSyncTime `json:"data"`
}

type GetLastSyncTime struct {
	OrderLastSyncTime   int64 `json:"orderLastSyncTime"`
	ProductLastSyncTime int64 `json:"productLastSyncTime"`
}

type ShopBaseInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	NameRemark     string `json:"nameRemark"`
	GroupName      string `json:"groupName"`
	Currency       string `json:"currency"`
	PublishSetting string `json:"publishSetting"` // 可能为空，依赖接口参数
	ApiInfo        string `json:"apiInfo"`        // 可能为空，依赖接口参数
}

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
	ShowAll bool `json:"showExtend"`
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
}
type ShopDetailQuery struct {
	Id int64 `json:"id"`
}

type ShopDetailResult struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    ShopDetail `json:"data"`
}
