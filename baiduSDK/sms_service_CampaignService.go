// sms_service_CampaignService.go
package baiduSDK

import ()

const (
	GET_Campaign    = "getCampaign"
	Add_Campaign    = "addCampaign"
	UPDATE_Campaign = "updateCampaign"
	DELETE_Campaign = "deleteCampaign"
)

// 计划服务CampaignService：
// 通过增加、删除、修改、查询计划的信息，用于管理账户下的任意的推广计划，包括计划的名称、预算、投放地域、IP排除、否定关键词、暂停时段、暂停/启动，或者设置投放设备、设置移动出价比例等。
type CampaignService struct {
	*CommonService
}

func NewCampaignService() *CampaignService {
	a := new(CampaignService)
	a.CommonService = NewCommonService("sms", "service", "CampaignService")
	return a
}
func (a *CampaignService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}

func (a *CampaignService) GetCampain(request interface{}) ([]byte, error) {
	return a.Execute(GET_Campaign, request)
}
func (a *CampaignService) AddCampain(request interface{}) ([]byte, error) {
	return a.Execute(Add_Campaign, request)
}
func (a *CampaignService) UpdateCampain(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_Campaign, request)
}
func (a *CampaignService) DeleteCampain(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_Campaign, request)
}

//定义获取计划请求结构
type CampaignGetRequest struct {
	CampaignIds    []int64  `json:"campaignIds"`
	CampaignFields []string `json:"campaignFields"`
}

func NewCampaignGetRequest() *CampaignGetRequest {
	c := new(CampaignGetRequest)
	return c
}

//定义获取计划响应结构
type CampaignGetResponse struct {
	Header *ResponseHeader          `json:"header"`
	Body   CampaignTypeResponseBody `json:"body"`
}
type CampaignTypeResponseBody struct {
	CampaignTypes []*CampaignType `json:"data"`
}

//定义添加计划请求结构
type CampaignAddRequest struct {
	CampaignTypes []*CampaignType `json:"campaignTypes"`
}

func NewCampaignAddRequest() *CampaignAddRequest {
	c := new(CampaignAddRequest)
	return c
}

//定义添加计划响应结构
type CampaignAddResponse struct {
	CampaignTypes []*CampaignType `json:"data"`
}

//定义更新计划请求结构
type CampaignUpdateRequest struct {
	CampaignTypes []*CampaignType `json:"campaignTypes"`
}

func NewCampaignUpdateRequest() *CampaignUpdateRequest {
	c := new(CampaignUpdateRequest)
	return c
}

//定义更新计划响应结构
type CampaignUpdateResponse struct {
	CampaignTypes []*CampaignType `json:"data"`
}

//定义删除计划请求结构
type CampaignDeleteRequest struct {
	CampaignIds []int64 `json:"campaignIds"`
}

func NewCampaignDeleteRequest() *CampaignDeleteRequest {
	c := new(CampaignDeleteRequest)
	return c
}

//定义删除计划响应结构
type CampaignDeleteResponse struct {
	CampaignTypes []*CampaignType `json:"data"`
}

//推广计划基础对象
type CampaignType struct {
	CampaignId           int64              `json:"campaignId"`
	CampaignName         string             `json:"campaignName"`
	Budget               float64            `json:"budget"`
	RegionTarget         []int              `json:"regionTarget"`
	NegativeWords        []string           `json:"negativeWords"`
	ExactNegativeWords   []string           `json:"exactNegativeWords"`
	Schedule             []*ScheduleType    `json:"schedule"`
	BudgetOfflineTime    []*OfflineTimeType `json:"budgetOfflineTime"`
	ShowProb             int                `json:"showProb"`
	Device               int                `json:"device"`
	PriceRatio           float64            `json:"priceRatio"`
	Pause                bool               `json:"pause"`
	RmktStatus           bool               `json:"rmktStatus"`
	Status               int                `json:"status"`
	CampaignType         int                `json:"campaignType"`
	IsDynamicCreative    bool               `json:"isDynamicCreative"`
	IsDynamicTagSublink  bool               `json:"isDynamicTagSublink"`
	IsDynamicTitle       bool               `json:"isDynamicTitle"`
	IsDynamicHotRedirect bool               `json:"isDynamicHotRedirect"`
	rmktPriceRatio       float64            `json:"rmktPriceRatio"`
}
type ScheduleType struct {
	WeekDay   int `json:"weekDay"`
	StartHour int `json:"startHour"`
	EndHour   int `json:"endHour"`
}

var CampaignFields = map[string]string{
	"campaignId":           "推广计划 id",
	"campaignName":         "计划名",
	"budget":               "计划预算",
	"budgetOfflineTime":    "最近 7 天预算撞线时间",
	"campaignType":         "计划类型",
	"device":               "推广设备",
	"exactNegativeWords":   "精确否定关键词",
	"isDynamicCreative":    "子链开关",
	"isDynamicTagSublink":  "标签子链开关",
	"isDynamicTitle":       "动态标题开关",
	"isDynamicHotRedirect": "热点直达开关",
	"regionTarget":         "计划推广地域",
	"negativeWords":        "计划否定关键词",
	"pause":                "计划启停",
	"rmktStatus":           "混合再营销计划启用/暂停",
	"priceRatio":           "无线出价比例",
	"schedule":             "循环暂停",
	"showProb":             "创意展现策略",
	"status":               "计划状态",
	"rmktPriceRatio":       "混合再营销出价比例",
}

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var WeekDayToString = map[int]string{
	Monday:    "星期一",
	Tuesday:   "星期二",
	Wednesday: "星期三",
	Thursday:  "星期四",
	Friday:    "星期五",
	Saturday:  "星期六",
	Sunday:    "星期日",
}

var StringToWeekDay = reverseInt(WeekDayToString)

func reverseInt(m map[int]string) map[string]int {
	n := make(map[string]int)
	for u, s := range m {
		n[s] = u
	}
	return n
}
