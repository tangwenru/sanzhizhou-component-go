package sanzhizhouComponentLib

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

func MainSystem(userToken string, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Sanzhizhou.ApiUrl")
	return ApiGet(userToken, apiUrl, apiPath, data, result)
}
func MainSystemPost(userToken string, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Sanzhizhou.ApiUrl")
	return ApiPost(userToken, apiUrl, apiPath, data, result)
}

func SubsystemVideoPublish(userId string, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	apiUrl, _ := beego.AppConfig.String("Subsystem.VideoPublishAPIUrl")
	if len(apiUrl) < 5 {
		return nil, errors.New("请配置 app.conf 文件的 Subsystem.VideoPublishAPIUrl")
	}
	return ApiGet(userId, apiUrl, apiPath, data, result)
}

func ApiGet(userToken string, apiUrl, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	url := apiUrl + strings.ToLower(apiPath)
	regSelfUrl := regexp.MustCompile("^https?:")
	// 强行是用自己的 domain
	if regSelfUrl.MatchString(apiPath) {
		url = apiPath
	}

	req := httplib.Get(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	dataByte, err := json.Marshal(data)

	postData := Encrypt(string(dataByte))
	req.Param("data", postData)
	req.Param("userToken", userToken)

	// 子系统名称
	appName, _ := beego.AppConfig.String("appName")
	reg := regexp.MustCompile("^api-")
	subSystem := reg.ReplaceAllString(appName, "")
	req.Param("subSystem", subSystem)

	//fmt.Println("ApiGet:", global.IdEncrypt(userId), url, postData)

	bytesResult, err := req.Bytes()
	if err != nil {
		fmt.Println("get api :", err)
		return []byte(""), err
	}

	//query := ThirdIpResult{}
	//errJson := json.Unmarshal(bytesResult, &query)
	//if errJson != nil {
	//	fmt.Println("err 21:", errJson)
	//	return ""
	//}

	errJson := json.Unmarshal(bytesResult, result)
	if errJson != nil {
		fmt.Println("err 1:", url)
		fmt.Println("err api get json:", errJson, string(bytesResult))
		return []byte(""), errJson
	}

	return bytesResult, nil
}

func ApiPost(userToken string, apiUrl, apiPath string, data interface{}, result interface{}) ([]byte, error) {
	url := apiUrl + strings.ToLower(apiPath)
	regSelfUrl := regexp.MustCompile("^https?:")
	// 强行是用自己的 domain
	if regSelfUrl.MatchString(apiPath) {
		url = apiPath
	}
	//
	if strings.Index(url, "?") == -1 {
		url += "?" // _method=post
	} else {
		url += "&" // =post
	}

	// 子系统名称
	appName, _ := beego.AppConfig.String("appName")
	reg := regexp.MustCompile("^api-")
	subSystem := reg.ReplaceAllString(appName, "")

	url += "&subSystem=" + subSystem
	url += "&userToken=" + userToken
	url += "&_method=post"
	//req.Param("subSystem", subSystem)
	//req.Param("userToken", userToken)
	//req.Param("_method", "post")

	req := httplib.Post(url)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	dataByte, err := json.Marshal(data)

	postData := Encrypt(string(dataByte))

	req.Header("Content-Type", "application/json")

	//fmt.Println("ApiGet:", global.IdEncrypt(userId), url, postData)

	bodyData := map[string]string{
		"data": postData,
		//"userToken": userToken,
		//"subSystem": subSystem,
	}

	bodyByte, _ := json.Marshal(bodyData)

	req.Body(bodyByte)

	bytesResult, err := req.Bytes()
	if err != nil {
		fmt.Println("post api:", err)
		return []byte(""), err
	}

	//fmt.Println(fmt.Sprintf("reqreqreq: %+v", req.GetRequest()))

	//query := ThirdIpResult{}
	//errJson := json.Unmarshal(bytesResult, &query)
	//if errJson != nil {
	//	fmt.Println("err 21:", errJson)
	//	return ""
	//}

	errJson := json.Unmarshal(bytesResult, result)
	if errJson != nil {
		fmt.Println("err post 1:", url)
		fmt.Println("err api post err:", errJson)
		fmt.Println("err api post json:", string(bytesResult))
		return []byte(""), errJson
	}

	return bytesResult, nil
}
