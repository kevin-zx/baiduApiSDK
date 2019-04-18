//sms_service_AdgroupService.go
package baiduSDK

// 单元服务AdgroupService：

// 通过增加、删除、修改、查询单元的信息，用于管理计划下的任意的推广单元，包括单元的名称、预算、投放地域、IP排除、否定关键词、暂停时段、暂停/启动，或者设置投放设备、设置移动出价比例等。
const (
	GET_Adgroup    = "getAdgroup"
	Add_Adgroup    = "addAdgroup"
	UPDATE_Adgroup = "updateAdgroup"
	DELETE_Adgroup = "deleteAdgroup"
)

type AdgroupService struct {
	*CommonService
}

func NewAdgroupService() *AdgroupService {
	a := new(AdgroupService)
	a.CommonService = NewCommonService("sms", "service", "AdgroupService")
	return a
}
func (a *AdgroupService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}

func (a *AdgroupService) GetAdgroup(request interface{}) ([]byte, error) {
	return a.Execute(GET_Adgroup, request)
}
func (a *AdgroupService) AddAdgroup(request interface{}) ([]byte, error) {
	return a.Execute(Add_Adgroup, request)
}
func (a *AdgroupService) UpdateAdgroup(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_Adgroup, request)
}
func (a *AdgroupService) DeleteAdgroup(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_Adgroup, request)
}

//定义获取单元请求结构
type AdgroupGetRequest struct {
	Ids           []int64  `json:"ids"`
	IdType        int      `json:"idType"`
	AdgroupFields []string `json:"adgroupFields"`
}

func NewAdgroupGetRequest() *AdgroupGetRequest {
	c := new(AdgroupGetRequest)
	return c
}

//定义获取单元响应结构
type AdgroupGetResponse struct {
	Header *ResponseHeader     `json:"header"`
	Body   AdgroupResponseBody `json:"body"`
}
type AdgroupResponseBody struct {
	AdgroupTypes []*AdgroupType `json:"data"`
}

//定义添加单元请求结构
type AdgroupAddRequest struct {
	AdgroupTypes []*AdgroupType `json:"adgroupTypes"`
}

func NewAdgroupAddRequest() *AdgroupAddRequest {
	c := new(AdgroupAddRequest)
	return c
}

//定义添加单元响应结构
type AdgroupAddResponse struct {
	AdgroupTypes []*AdgroupType `json:"data"`
}

//定义更新单元请求结构
type AdgroupUpdateRequest struct {
	AdgroupTypes []*AdgroupType `json:"adgroupTypes"`
}

func NewAdgroupUpdateRequest() *AdgroupUpdateRequest {
	c := new(AdgroupUpdateRequest)
	return c
}

//定义更新单元响应结构
type AdgroupUpdateResponse struct {
	AdgroupTypes []*AdgroupType `json:"data"`
}

//定义删除单元请求结构
type AdgroupDeleteRequest struct {
	AdgroupIds []int64 `json:"adgroupIds"`
}

func NewAdgroupDeleteRequest() *AdgroupDeleteRequest {
	c := new(AdgroupDeleteRequest)
	return c
}

//定义删除单元响应结构
type AdgroupDeleteResponse struct {
	AdgroupTypes []*AdgroupType `json:"data"`
}

//单元基础对象
type AdgroupType struct {
	AdgroupId          int64    `json:"adgroupId"`
	CampaignId         int64    `json:"campaignId"`
	AdgroupName        string   `json:"adgroupName"`
	MaxPrice           float64  `json:"maxPrice"`
	NegativeWords      []string `json:"negativeWords"`
	ExactNegativeWords []string `json:"exactNegativeWords"`
	Pause              bool     `json:"pause"`
	Status             int      `json:"status"`
	PriceRatio         float64  `json:"priceRatio"`
	AccuPriceFactor    float64  `json:"accuPriceFactor"`
	WordPriceFactor    float64  `json:"wordPriceFactor"`
	WidePriceFactor    float64  `json:"widePriceFactor"`
	MatchPriceStatus   int      `json:"matchPriceStatus"`
}

var AdgroupFields = map[string]string{
	"adgroupId":          "推广单元 id", //基本字段
	"campaignId":         "推广计划 id", //基本字段
	"adgroupName":        "推广单元名称",
	"maxPrice":           "推广单元最高出价",
	"negativeWords":      "单元否定关键词",
	"exactNegativeWords": "单元精确否定关键词",
	"pause":              "暂停/启用推广单元",
	"status":             "推广单元状态",
	"priceRatio":         "单元移动出价比例",
	"accuPriceFactor":    "精确出价比例",
	"wordPriceFactor":    "短语出价比例",
	"widePriceFactor":    "广泛出价比例",
	"matchPriceStatus":   "分匹配状态",
}
