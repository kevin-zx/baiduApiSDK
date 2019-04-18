// sms_service_DynamicCreativeService.go
package baiduSDK

//动态创意片段服务
import ()

const (
	GET_DynCreative               = "getDynCreative"
	GET_ExclusionTypeByCampaignId = "getExclusionTypeByCampaignId"
	ADD_DynCreative               = "addDynCreative"
	DELETE_DynCreative            = "deleteDynCreative"
	UPDATE_DynCreative            = "updateDynCreative"
	ADD_ExclusionType             = "addExclusionType"
	DELETE_ExclusionType          = "delExclusionType"
)

type DynamicCreativeService struct {
	*CommonService
}

func NewDynamicCreativeService() *DynamicCreativeService {
	a := new(DynamicCreativeService)
	a.CommonService = NewCommonService("sms", "service", "DynamicCreativeService")
	return a
}
func (a *DynamicCreativeService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}
func (a *DynamicCreativeService) GetDynCreative(request interface{}) ([]byte, error) {
	return a.Execute(GET_DynCreative, request)
}
func (a *DynamicCreativeService) GetExclusionTypeByCampaignId(request interface{}) ([]byte, error) {
	return a.Execute(GET_ExclusionTypeByCampaignId, request)
}
func (a *DynamicCreativeService) AddDynCreative(request interface{}) ([]byte, error) {
	return a.Execute(ADD_DynCreative, request)
}
func (a *DynamicCreativeService) DeleteDynCreative(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_DynCreative, request)
}
func (a *DynamicCreativeService) UpdateDynCreative(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_DynCreative, request)
}
func (a *DynamicCreativeService) AddExclusionType(request interface{}) ([]byte, error) {
	return a.Execute(ADD_ExclusionType, request)
}
func (a *DynamicCreativeService) DelExclusionType(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_ExclusionType, request)
}

//getDynCreative start
type DynCreativeGetRequest struct {
	Ids               []int64  `json:"ids"`
	IdType            int      `json:"idType"`
	DynCreativeFields []string `json:"dynCreativeFields"`
}

func NewDynCreativeGetRequest() *DynCreativeGetRequest {
	return new(DynCreativeGetRequest)
}

type DynCreativeGetResponse struct {
	Header *ResponseHeader            `json:"header"`
	Body   DynCreativeGetResponseBody `json:"body"`
}
type DynCreativeGetResponseBody struct {
	DynCreativeTypes []*DynCreativeType `json:"dynCreativeTypes"`
}

//getDynCreative end

//getExclusionTypeByCampaignId start 根据指定推广计划 id 获取推广计划下所有排除对象(id 可批量)。

type ExclusionTypeByCampaignIdGetRequest struct {
	CampaignIds []int64 `json:"campaignIds"`
}

func NewExclusionTypeByCampaignIdGetRequest() *ExclusionTypeByCampaignIdGetRequest {
	return new(ExclusionTypeByCampaignIdGetRequest)
}

type ExclusionTypeByCampaignIdGetResponse struct {
	Header *ResponseHeader                          `json:"header"`
	Body   ExclusionTypeByCampaignIdGetResponseBody `json:"body"`
}
type ExclusionTypeByCampaignIdGetResponseBody struct {
	DynCreativeExclusionTypes []*DynCreativeExclusionType `json:"data"`
}

//getExclusionTypeByCampaignId end

//AddDynCreative start
type AddDynCreativeGetRequest struct {
	DynCreativeTypes []*DynCreativeType `json:"dynCreativeTypes"`
}

func NewAddDynCreativeGetRequest() *AddDynCreativeGetRequest {
	return new(AddDynCreativeGetRequest)
}

type AddDynCreativeGetResponse struct {
	Header *ResponseHeader               `json:"header"`
	Body   AddDynCreativeGetResponseBody `json:"body"`
}
type AddDynCreativeGetResponseBody struct {
	DynCreativeTypes []*DynCreativeType `json:"data"`
}

//AddDynCreative end

//DeleteDynCreative start
type DeleteDynCreativeGetRequest struct {
	DynCreativeIds []int64 `json:"dynCreativeIds"`
}

func NewDeleteDynCreativeGetRequest() *DeleteDynCreativeGetRequest {
	return new(DeleteDynCreativeGetRequest)
}

type DeleteDynCreativeGetResponse struct {
	Header *ResponseHeader                  `json:"header"`
	Body   DeleteDynCreativeGetResponseBody `json:"body"`
}
type DeleteDynCreativeGetResponseBody struct {
	DynCreativeTypes []*DynCreativeType `json:"data"`
}

//DeleteDynCreative end

//UpdateDynCreative start
type UpdateDynCreativeGetRequest struct {
	DynCreativeTypes []*DynCreativeType `json:"dynCreativeTypes"`
}

func NewUpdateDynCreativeGetRequest() *UpdateDynCreativeGetRequest {
	return new(UpdateDynCreativeGetRequest)
}

type UpdateDynCreativeGetResponse struct {
	Header *ResponseHeader                  `json:"header"`
	Body   UpdateDynCreativeGetResponseBody `json:"body"`
}
type UpdateDynCreativeGetResponseBody struct {
	DynCreativeTypes []*DynCreativeType `json:"data"`
}

//UpdateDynCreative end

type DynCreativeType struct {
	DynCreativeId int64  `json:"dynCreativeId"`
	CampaignId    int64  `json:"campaignId"`
	AdgroupId     int64  `json:"adgroupId"`
	BindingType   int    `json:"bindingType"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	Murl          string `json:"murl"`
	Pause         bool   `json:"pause"`
	Status        int    `json:"status"`
}

//AddExclusionType start
type AddExclusionTypeGetRequest struct {
	DynCreativeExclusionTypes []*DynCreativeExclusionType `json:"dynCreativeExclusionTypes"`
}

func NewAddExclusionTypeGetRequest() *AddExclusionTypeGetRequest {
	return new(AddExclusionTypeGetRequest)
}

type AddExclusionTypeGetResponse struct {
	Header *ResponseHeader                 `json:"header"`
	Body   AddExclusionTypeGetResponseBody `json:"body"`
}
type AddExclusionTypeGetResponseBody struct {
	DynCreativeExclusionTypes []*DynCreativeExclusionType `json:"data"`
}

//AddExclusionType end

//DelExclusionType start 􏰀
//方法说明按计划 id 及动态创意排除对象 id 删除动态创意排除对象
type DelExclusionTypeGetRequest struct {
	DynCreativeExclusionTypes []DynCreativeExclusionType `json:"dynCreativeExclusionTypes"`
}

func NewDelExclusionTypeGetRequest() *DelExclusionTypeGetRequest {
	return new(DelExclusionTypeGetRequest)
}

type DelExclusionTypeGetResponse struct {
	Header *ResponseHeader                 `json:"header"`
	Body   DelExclusionTypeGetResponseBody `json:"body"`
}
type DelExclusionTypeGetResponseBody struct {
	OptResult int `json:"optResult"`
}

//DelExclusionType end

type DynCreativeExclusionType struct {
	ExclusionTypes []*ExclusionType `json:"exclusionTypes"`
	CampaignId     int64            `json:"campaignId"`
}
type ExclusionType struct {
	ExclusionId      int64  `json:"exclusionId"`
	ExclusionContent string `json:"exclusionContent"`
	ExclusionType    int    `json:"exclusionType"`
}

var DynCreativeFields = map[string]string{
	"dynCreativeId": "动态创意 ",
	"campaignId":    "推广计划ID",
	"adgroupId":     "推广单元ID",
	"bindingType":   "动态创意绑定层级",
	"title":         "动态创意描述",
	"url":           "默认访问 url",
	"murl":          "移动访问 url",
	"status":        "动态创意状态",
	"pause":         "启停状态",
}
