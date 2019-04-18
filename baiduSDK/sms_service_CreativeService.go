//sms_service_CreativeService.go
package baiduSDK

// 创意服务CreativeService：

// 通过增加、删除、修改、查询创意的信息，用于管理账户下任意的创意； 例如，您可以批量获取和修改某些创意的标题、描述、PC访问与显示URL、移动访问与显示URL、暂停/启用等属性。
import ()

const (
	GET_Creative    = "getCreative"
	Add_Creative    = "addCreative"
	UPDATE_Creative = "updateCreative"
	DELETE_Creative = "deleteCreative"
)

type CreativeService struct {
	*CommonService
}

func NewCreativeService() *CreativeService {
	a := new(CreativeService)
	a.CommonService = NewCommonService("sms", "service", "CreativeService")
	return a
}
func (a *CreativeService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}

func (a *CreativeService) GetCreative(request interface{}) ([]byte, error) {
	return a.Execute(GET_Creative, request)
}
func (a *CreativeService) AddCreative(request interface{}) ([]byte, error) {
	return a.Execute(Add_Creative, request)
}
func (a *CreativeService) UpdateCreative(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_Creative, request)
}
func (a *CreativeService) DeleteCreative(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_Creative, request)
}

//定义获取创意请求结构
type CreativeGetRequest struct {
	Ids            []int64  `json:"ids"`
	IdType         int      `json:"idType"`
	GetTemp        int      `json:"getTemp"`
	CreativeFields []string `json:"creativeFields"`
}

func NewCreativeGetRequest() *CreativeGetRequest {
	c := new(CreativeGetRequest)
	return c
}

//定义获取创意响应结构
type CreativeGetResponse struct {
	Header *ResponseHeader      `json:"header"`
	Body   CreativeResponseBody `json:"body"`
}
type CreativeResponseBody struct {
	CreativeTypes []*CreativeType `json:"data"`
}

//定义添加创意请求结构
type CreativeAddRequest struct {
	CreativeTypes []*CreativeType `json:"creativeTypes"`
}

func NewCreativeAddRequest() *CreativeAddRequest {
	c := new(CreativeAddRequest)
	return c
}

//定义添加创意响应结构
type CreativeAddResponse struct {
	CreativeTypes []*CreativeType `json:"data"`
}

//定义更新创意请求结构
type CreativeUpdateRequest struct {
	CreativeTypes []*CreativeType `json:"creativeTypes"`
}

func NewCreativeUpdateRequest() *CreativeUpdateRequest {
	c := new(CreativeUpdateRequest)
	return c
}

//定义更新创意响应结构
type CreativeUpdateResponse struct {
	CreativeTypes []*CreativeType `json:"data"`
}

//定义删除创意请求结构
type CreativeDeleteRequest struct {
	CreativeIds []int64 `json:"creativeIds"`
}

func NewCreativeDeleteRequest() *CreativeDeleteRequest {
	c := new(CreativeDeleteRequest)
	return c
}

//定义删除创意响应结构
type CreativeDeleteResponse struct {
	CreativeTypes []*CreativeType `json:"data"`
}

//创意服务
type CreativeType struct {
	CreativeId           int64  `json:"creativeId"`
	AdgroupId            int64  `json:"adgroupId"`
	Title                string `json:"title"`
	Description1         string `json:"description1"`
	Description2         string `json:"description2"`
	PcDestinationUrl     string `json:"pcDestinationUrl"`
	PcDisplayUrl         string `json:"pcDisplayUrl"`
	MobileDestinationUrl string `json:"mobileDestinationUrl"`
	MobileDisplayUrl     string `json:"mobileDisplayUrl"`
	Pause                bool   `json:"pause"`
	Status               int    `json:"status"`
	DevicePreference     int    `json:"devicePreference"`
}

var CreativeFields = map[string]string{
	"creativeId":           "创意 id",
	"adgroupId":            "推广单元 id",
	"title":                "创意标题",
	"description1":         "创意描述第一行",
	"description2":         "创意描述第二行",
	"pcDestinationUrl":     "访问 url",
	"pcDisplayUrl":         "显示 url",
	"mobileDestinationUrl": "访问 url",
	"mobileDisplayUrl":     "显示 url",
	"pause":                "暂停/启用创意",
	"status":               "创意状态",
	"devicePreference":     "设备偏好",
}
