// sms_service_ToolkitService.go
package baiduSDK

import ()

const (
	GET_OperationRecord = "getOperationRecord"
)

type ToolkitService struct {
	*CommonService
}

func NewToolkitService() *ToolkitService {
	a := new(ToolkitService)
	a.CommonService = NewCommonService("sms", "service", "ToolkitService")
	return a
}
func (a *ToolkitService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}
func (a *ToolkitService) GetOperationRecord(request interface{}) ([]byte, error) {
	return a.Execute(GET_OperationRecord, request)
}

// type OperationRecordGetRequest struct {
// 	OperationRecordRequestType *OperationRecordRequestType `json:"operationRecordRequestType"`
// }

func NewOperationRecordGetRequest() *OperationRecordRequestType {
	return new(OperationRecordRequestType)
}

type OperationRecordGetResponseBody struct {
	OptLogTypes []*OptlogType `json:"data"`
}
type OperationRecordGetResponse struct {
	Header *ResponseHeader                `json:"header"`
	Body   OperationRecordGetResponseBody `json:"body"`
}

//操作记录
type OptlogType struct {
	UserId     int64  `json:"userId"`
	PlanId     int64  `json:"planId"`
	UnitId     int64  `json:"unitId"`
	OptTime    string `json:"optTime"`
	OptContent string `json:"optContent"`
	OptType    int    `json:"optType"`
	OptLevel   int    `json:"optLevel"`
	OldValue   string `json:"oldValue"`
	NewValue   string `json:"newValue"`
	OptObj     string `json:"optObj"`
}

type OperationRecordRequestType struct {
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	OptTypes    []int    `json:"optTypes"`
	OptLevel    int      `json:"optLevel"` // 1:推广单元 2:推广计划 3:帐户4:创意 5:关键词 7:蹊径子链 8:App 推广 9:推广电话 10:商桥 11:页面回呼 12:动态创意
	OptContents []string `json:"optContents"`
}

var OptLevelFields = map[int][]string{
	1: []string{
		"addUnit",
		"delUnit",
		"shelve",
		"bidPrice",
		"updUnitName",
		"negativeWord",
		"accurateNegativeWord",
		"mobilePriceFactor",
		"matchPriceFactorStatus",
		"matchPriceFactor"},
	2: []string{
		"addPlan",
		"delPlan",
		"shelve",
		"cycShelve",
		"dailyBudget",
		"zone",
		"ipExclude",
		"showFac",
		"negativeWord",
		"accurateNegativeWord",
		"cproJoin",
		"setCproPrice",
		"deviceCfgStat",
		"targetDevice",
		"phoneNumber",
		"bridge",
		"mobilePrice",
		"queryRegion",
		"dynamicIdeaStat"},
	3: []string{
		"budget",
		"dailyBudget",
		"budgetDay2Day",
		"budgetDay2Week",
		"budgetWeek2Day",
		"budgetWeek2Week",
		"zone",
		"materialActive",
		"ipExclude",
		"queryRegion",
		"dynamicIdeaStat",
		"updExternalFlow"},
	4: []string{
		"addIdea",
		"delIdea",
		"shelve",
		"updIdea",
		"active",
		"deviceOpt"},
	5: []string{
		"addWord",
		"delWord",
		"shelve",
		"bidPrice",
		"updWordMatch",
		"active",
		"wordTransfer",
		"updWordUrl",
		"updWordMobileUrl",
		"matchPrefer"},
	7: []string{
		"addSublink",
		"delSublink",
		"updSublink",
		"shelveSublink"},
	8: []string{
		"addAppCreative",
		"delAppCreative",
		"updAppCreative",
		"shelveAppCreative"},
	9: []string{
		"addPhoneCreative",
		"delPhoneCreative",
		"updPhoneCreative",
		"shelvePhoneCreative"},
	10: []string{
		"addBridgeCreative",
		"delBridgeCreative",
		"shelveBridgeCreative"},
	11: []string{
		"addLXBCreative",
		"delLXBCreative",
		"updLXBCreative",
		"shelveLXBCreative"},
	12: []string{
		"addDynamicIdea",
		"delDynamicIdea",
		"updDynamicIdea",
		"shelveDynamicIdea"},
}
