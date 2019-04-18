//sms_service_NewNewCreativeService.go
package baiduSDK

// 新创意（蹊径）服务
// NewNewCreativeService： 通过增加、删除、修改、查询蹊径的信息，用于管理账户下任意的蹊径创意；
// 例如，您可以批量获取和修改某些蹊径创意的标题、描述、PC访问与显示URL、移动访问与显示URL、暂停/启用等属性。
import ()

const (
	//获取
	GET_Sublink = "getSublink"
	GET_Bridge  = "getBridge"
	GET_Phone   = "getPhone"
	//添加
	ADD_Sublink = "addSublink"
	ADD_Bridge  = "addBridge"
	ADD_Phone   = "addPhone"
	//更新
	UPDATE_Sublink = "updateSublink"
	UPDATE_Phone   = "updatePhone"
	UPDATE_Bridge  = "updateBridge"
	//删除
	DELETE_Sublink = "deleteSublink"
)

type New_CreativeService struct {
	*CommonService
}

func NewNewCreativeService() *New_CreativeService {
	a := new(New_CreativeService)
	a.CommonService = NewCommonService("sms", "service", "NewCreativeService")
	return a
}
func (a *New_CreativeService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}

func (a *New_CreativeService) GetSublink(request interface{}) ([]byte, error) {
	return a.Execute(GET_Sublink, request)
}
func (a *New_CreativeService) GetBridge(request interface{}) ([]byte, error) {
	return a.Execute(GET_Bridge, request)
}
func (a *New_CreativeService) GetPhone(request interface{}) ([]byte, error) {
	return a.Execute(GET_Phone, request)
}

//定义获取推广子链请求结构
type SublinkGetRequest struct {
	Ids           []int64  `json:"ids"`
	IdType        int      `json:"idType"`
	GetTemp       int      `json:"getTemp"`
	SublinkFields []string `json:"sublinkFields"`
	Device        int      `json:"device"`
}

func NewSublinkGetRequest() *SublinkGetRequest {
	c := new(SublinkGetRequest)
	return c
}

//定义推广子链响应结构
type SublinkGetResponse struct {
	Header *ResponseHeader         `json:"header"`
	Body   SublinkTypeResponseBody `json:"body"`
}
type SublinkTypeResponseBody struct {
	SublinkTypes []*SublinkType `json:"data"`
}

//定义添加推广子链请求结构
type SublinkAddRequest struct {
	SublinkTypes []*SublinkType `json:"sublinkTypes"`
}

func NewKSublinkAddRequest() *SublinkAddRequest {
	c := new(SublinkAddRequest)
	return c
}

//定义添加推广子链响应结构
type SublinkAddResponse struct {
	SublinkTypes []*SublinkType `json:"data"`
}

//定义更新推广子链请求结构
type SublinkUpdateRequest struct {
	SublinkTypes []*SublinkType `json:"sublinkTypes"`
}

func NewKSublinkUpdateRequest() *SublinkUpdateRequest {
	c := new(SublinkUpdateRequest)
	return c
}

//定义更新推广子链响应结构
type SublinkUpdateResponse struct {
	SublinkTypes []*SublinkType `json:"data"`
}

//定义删除推广子链请求结构
type SublinkDeleteRequest struct {
	SublinkIds []int64 `json:"sublinkIds"`
}

func NewSublinkDeleteRequest() *SublinkDeleteRequest {
	c := new(SublinkDeleteRequest)
	return c
}

//定义删除推广子链响应结构
type SublinkDeleteResponse struct {
	SublinkTypes []*SublinkType `json:"data"`
}

//推广子链的数据类型
type SublinkType struct {
	SublinkId    int64          `json:"sublinkId"`
	SublinkInfos []*SublinkInfo `json:"sublinkInfos"`
	AdgroupId    int64          `json:"adgroupId"`
	Pause        bool           `json:"pause"`
	Status       int            `json:"status"`
	Device       int            `json:"device"`
}
type SublinkInfo struct {
	Description    string `json:"description"`
	DestinationUrl string `json:"destinationUrl"`
}

//定义电话请求结构
type PhoneGetRequest struct {
	Ids         []int64  `json:"ids"`
	IdType      int      `json:"idType"`
	PhoneFields []string `json:"phoneFields"`
}

func NewPhoneGetRequest() *PhoneGetRequest {
	c := new(PhoneGetRequest)
	return c
}

//定义电话响应结构
type PhoneGetResponse struct {
	Header *ResponseHeader       `json:"header"`
	Body   PhoneTypeResponseBody `json:"body"`
}
type PhoneTypeResponseBody struct {
	PhoneTypes []*PhoneType `json:"data"`
}

//定义更新电话请求结构
type PhoneUpdateRequest struct {
	PhoneTypes []*PhoneType `json:"phoneTypes"`
}

func NewKPhoneUpdateRequest() *PhoneUpdateRequest {
	c := new(PhoneUpdateRequest)
	return c
}

//定义更新电话响应结构
type PhoneUpdateResponse struct {
	PhoneTypes []*PhoneType `json:"data"`
}

type PhoneType struct {
	PhoneId   int64  `json:"phoneId"`
	PhoneNum  string `json:"phoneNum"`
	AdgroupId int64  `json:"adgroupId"`
	Pause     bool   `json:"pause"`
	Status    int    `json:"status"`
}
type BridgeGetRequest struct {
	Ids          []int64  `json:"ids"`
	IdType       int      `json:"idType"`
	BridgeFields []string `json:"bridgeFields"`
}

func NewBridgeGetRequest() *BridgeGetRequest {
	c := new(BridgeGetRequest)
	return c
}

type BridgeGetResponse struct {
	Header *ResponseHeader        `json:"header"`
	Body   BridgeTypeResponseBody `json:"body"`
}
type BridgeTypeResponseBody struct {
	BridgeTypes []*BridgeType `json:"data"`
}

//定义更新商桥请求结构
type BridgeUpdateRequest struct {
	BridgeTypes []*BridgeType `json:"bridgeTypes"`
}

func NewKBridgeUpdateRequest() *BridgeUpdateRequest {
	c := new(BridgeUpdateRequest)
	return c
}

//定义更新商桥响应结构
type BridgeUpdateResponse struct {
	BridgeTypes []*BridgeType `json:"data"`
}

//商桥
type BridgeType struct {
	BridgeId  int64 `json:"bridgeId"`
	AdgroupId int64 `json:"adgroupId"`
	Pause     bool  `json:"pause"`
	Status    int   `json:"status"`
}

var SublinkFields = map[string]string{
	"sublinkId":    "推广子链的 id",
	"sublinkInfos": "推广子链描述集合",
	"adgroupId":    "推广单元 id",
	"pause":        "暂停/启用创意",
	"status":       "推广子链状态",
}
var BridgeFields = map[string]string{
	"bridgeId":  "商桥的 id",
	"adgroupId": "推广单元 id",
	"pause":     "暂停/启用商桥",
	"status":    "商桥状态",
}
var PhoneFields = map[string]string{
	"phoneId":   "电话 id",
	"phoneNum":  "电话号码",
	"adgroupId": "推广单元 id",
	"pause":     "暂停/启用",
	"status":    "状态",
}
