package sanzhizhouComponent

type Huoshan struct {
	ApiId  string
	ApiKey string
}

func init() {

}

func NewHuoshan(apiId, apiKey string) *Huoshan {
	huoshan := &Huoshan{}
	huoshan.ApiId = apiId
	huoshan.ApiKey = apiKey

	return huoshan
}
