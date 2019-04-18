// CommonService.go
package baiduSDK

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://api.baidu.com"
)

//通用结构，包含验证信息 token ，username ，password
type CommonService struct {
	Productline string
	Version     string
	ServiceName string
	Serviceurl  string
	AuthHeader  *AuthHeader
}

func NewCommonService(productline, version, serviceName string) *CommonService {
	c := new(CommonService)
	c.Productline = productline
	c.Version = version
	c.ServiceName = serviceName
	c.Serviceurl = URL
	c.AuthHeader = NewAuthHeader()
	c.AuthHeader.Action = "API-SDK"
	return c
}

//执行post的请求，json 交互 获取 返回数据
func (c *CommonService) do(method string, request interface{}) (result []byte, err error) {
	//构建URL
	url := c.Serviceurl + "/json/" + c.Productline + "/" + c.Version + "/" +
		c.ServiceName + "/" + method
	//https x509解决
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	datas := NewReQuestData(c.AuthHeader, request)

	b, err := json.Marshal(datas)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer([]byte(b))

	res, err := client.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	result, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return
}

//设置通用结构值

func (c *CommonService) SetProductline(productline string) {
	c.Productline = productline
}
func (c *CommonService) GetProductline() string {
	return c.Productline
}
func (c *CommonService) SetVersion(version string) {
	c.Version = version
}
func (c *CommonService) GetVersion() string {
	return c.Version
}
func (c *CommonService) SetServiceName(serviceName string) {
	c.ServiceName = serviceName
}
func (c *CommonService) GetServiceName() string {
	return c.ServiceName
}
func (c *CommonService) SetServiceurl(Serviceurl string) {
	c.Serviceurl = Serviceurl
}
func (c *CommonService) GetServiceurl() string {
	return c.Serviceurl
}
func (c *CommonService) SetAuthHeader(authHeader *AuthHeader) {
	c.AuthHeader = authHeader
}
func (c *CommonService) GetAuthHeader() *AuthHeader {
	return c.AuthHeader
}

//request authheader struct
type AuthHeader struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Target      string `json:"target"`
	Token       string `json:"token"`
	Action      string `json:"action"`
	AccessToken string `json:"accesstoken"`
	AccountType string `json:"account_type"`
}

func NewAuthHeader() *AuthHeader {
	return new(AuthHeader)
}
func (a *AuthHeader) SetUsername(username string) {
	a.Username = username
}
func (a *AuthHeader) GetUsername() string {
	return a.Username
}
func (a *AuthHeader) SetPassword(password string) {
	a.Password = password
}
func (a *AuthHeader) GetPassword() string {
	return a.Password
}
func (a *AuthHeader) SetTarget(target string) {
	a.Target = target
}
func (a *AuthHeader) GetTarget() string {
	return a.Target
}
func (a *AuthHeader) SetToken(token string) {
	a.Token = token
}
func (a *AuthHeader) GetToken() string {
	return a.Token
}
func (a *AuthHeader) SetAction(action string) {
	a.Action = action
}
func (a *AuthHeader) GetAction() string {
	return a.Action
}
func (a *AuthHeader) SetAccessToken(accessToken string) {
	a.AccessToken = accessToken
}
func (a *AuthHeader) GetAccessToken() string {
	return a.AccessToken
}
func (a *AuthHeader) SetAccountType(accountType string) {
	a.AccountType = accountType
}
func (a *AuthHeader) GetAccountType() string {
	return a.AccountType
}

//提交通用request数据格式
type RequestData struct {
	Header *AuthHeader `json:"header"`
	Body   interface{} `json:"body"`
}

func NewReQuestData(header *AuthHeader, body interface{}) *RequestData {
	r := new(RequestData)
	r.Header = header
	r.Body = body
	return r
}

//返回通用 response
type ResponseHeader struct {
	Desc     string      `json:"desc"`
	Failures []*Failures `json:"failures"`
	Oprs     int         `json:"oprs"`
	Oprtime  int         `json:"oprtime"`
	Quota    int         `json:"quota"`
	Rquota   int         `json:"rquota"`
	Succ     int         `json:"succ"`
	Status   int         `json:"status"`
}
type Failures struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Position string `json:"position"`
}

func (r *ResponseHeader) GetResponseStatus() (bool, string) {

	if r.Status > 0 {

		var message string
		if len(r.Failures) > 0 {
			message += "code:" + GetIntStr(r.Failures[0].Code) + ",message:" +
				r.Failures[0].Message + "\r\n"
		}

		return false, message
	}
	return true, ""

}

// type service interface {
// 	Execute(method string, request interface{}) ([]byte, error)
// }
// type Instance func() service

// var adapters = make(map[string]Instance)

// func Register(name string, adapter Instance) {
// 	if adapter == nil {
// 		panic("service: Register adapter is nil")
// 	}
// 	if _, ok := adapters[name]; ok {
// 		panic("service: Register called twice for adapter " + name)
// 	}
// 	adapters[name] = adapter
// }

// func NewService(adapterName string) (adapter service, err error) {
// 	instanceFunc, ok := adapters[adapterName]
// 	if !ok {
// 		err = fmt.Errorf("service: unknown adapter name %q (forgot to import?)", adapterName)
// 		return
// 	}
// 	adapter = instanceFunc()
// 	if err != nil {
// 		adapter = nil
// 	}
// 	return
// }
