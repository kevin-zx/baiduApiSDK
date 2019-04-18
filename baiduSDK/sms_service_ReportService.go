// sms_service_ReportService.go
package baiduSDK

import (
	"encoding/json"
	// "time"
	"errors"
	"fmt"
)

// 通过 ReportService,您可以获取推广效果的详细报表。
// 说明: 目前报告总体分为两种类型:实时报告和异步报告。实时报告适用于小量、移动推广客户端等等
// 来使用,返回条数有限制,但是实时响应性强。异步报告需要从获取的链接进行报告下载。
// 目前分小时报告提供到“账户”和“计划”粒度。 搜索词报告数据非完全实时,系统无法为您提供 3 个小时以内的结果。
const (
	GET_RealTimeQueryData    = "getRealTimeQueryData"
	GET_RealTimePairData     = "getRealTimePairData"
	GET_ReportState          = "getReportState"
	GET_ReportFileUrl        = "getReportFileUrl"
	GET_RealTimeData         = "getRealTimeData"
	GET_ProfessionalReportId = "getProfessionalReportId"
	//入库报表
	AsyncReportAccount     = "account"
	AsyncReportCampaingn   = "campaign"
	AsyncReportAdgroup     = "adgroup"
	AsyncReportKeyword     = "keyword"
	AsyncReportCreative    = "creative"
	AsyncReportWordid      = "word"
	AsyncReportLocation    = "location"
	AsyncReportSearchkey   = "search"
	AsyncReportpair        = "pair"
	AsyncReportNewCreative = "xijing"
	AsyncReportLocation2   = "location_l2"
	AsyncReportHistory     = "history"
	//时间类型
	UnitOfTimeYear    = 1
	UnitOfTimeMonth   = 3
	UnitOfTimeWeekday = 4
	UnitOfTimeDay     = 5
	UnitOfTimeHour    = 7
	UnitOfTimeRange   = 8
	// 1,3,4,5,7,8 (年报,月报,周报,日报,小时报,请求时间段汇总)
)

var Device = map[string]int{
	"pc":     1,
	"mobile": 2,
}

func GetUnitTime(ut string) int {
	switch ut {
	case "year":
		return UnitOfTimeYear
	case "month":
		return UnitOfTimeMonth
	case "weekday":
		return UnitOfTimeWeekday
	case "day":
		return UnitOfTimeDay
	case "hour":
		return UnitOfTimeHour
	case "range":
		return UnitOfTimeRange
	default:
		return UnitOfTimeDay
	}
}

type ReportService struct {
	*CommonService
}

func NewReportService() *ReportService {
	a := new(ReportService)
	a.CommonService = NewCommonService("sms", "service", "ReportService")
	return a
}
func (a *ReportService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}
func (a *ReportService) GetRealTimeQueryData(request interface{}) ([]byte, error) {
	return a.Execute(GET_RealTimeQueryData, request)
}
func (a *ReportService) GetRealTimePairData(request interface{}) ([]byte, error) {
	return a.Execute(GET_RealTimePairData, request)
}
func (a *ReportService) GetReportState(request interface{}) ([]byte, error) {
	return a.Execute(GET_ReportState, request)
}
func (a *ReportService) GetReportFileUrl(request interface{}) ([]byte, error) {
	return a.Execute(GET_ReportFileUrl, request)
}
func (a *ReportService) GetRealTimeData(request interface{}) ([]byte, error) {
	return a.Execute(GET_RealTimeData, request)
}
func (a *ReportService) GetProfessionalReportId(request interface{}) ([]byte, error) {
	return a.Execute(GET_ProfessionalReportId, request)
}

type AsyncReportGetStruct struct {
	ReportType string
	StartDate  string
	EndDate    string
	Device     int
	UnitOfTime int
}

func NewAsyncReportGetStruct() *AsyncReportGetStruct {
	return &AsyncReportGetStruct{
		UnitOfTime: UnitOfTimeDay,
	}
}

//获取异步ID
func (a *ReportService) GetProfessionalReportIdByType(asyncReportGetStruct *AsyncReportGetStruct) (string, error) {

	prg := NewProfessionalReportIdGetRequest()
	reportRequestType := a.GetAsyncReportType(asyncReportGetStruct)
	prg.ReportRequestType = reportRequestType
	var (
		body []byte
		err  error
	)
	body, err = a.GetProfessionalReportId(prg)

	if err != nil {
		return "", err
	}

	var prid ProfessionalReportIdGetResponse
	err = json.Unmarshal(body, &prid)
	if err != nil {
		return "", err
	}
	if prid.Header.Desc != "success" {
		errmsg := fmt.Sprintf("ReportType:%s Code:%d Message:%s Position:%s", asyncReportGetStruct.ReportType, prid.Header.Failures[0].Code,
			prid.Header.Failures[0].Message, prid.Header.Failures[0].Position)
		return "", errors.New(errmsg)
	}
	if len(prid.Body.Data) == 0 {
		return "", errors.New("ReportService GetProfessionalReportIdByType data is null")
	}
	return prid.Body.Data[0].ReportId, nil
}

//通过异步报告ID，获取状态
func (a *ReportService) GetReportStateById(reportid string) (bool, error) {
	rsg := NewReportStateGetRequest()
	rsg.ReportId = reportid
	body, err := a.GetReportState(rsg)
	if err != nil {
		return false, err
	}
	var rsgr ReportStateGetResponse
	err = json.Unmarshal(body, &rsgr)
	if err != nil {
		return false, err
	}
	if rsgr.Header.Desc != "success" {
		errmsg := fmt.Sprintf("Code:%d Message:%s Position:%s", rsgr.Header.Failures[0].Code,
			rsgr.Header.Failures[0].Message, rsgr.Header.Failures[0].Position)
		return false, errors.New(errmsg)
	}

	if len(rsgr.Body.Data) == 0 {
		return false, errors.New("ReportService GetReportStateById data is null")
	}

	if rsgr.Body.Data[0].IsGenerated == 3 {
		return true, nil
	}
	return false, nil
}
func (a *ReportService) GetReportFileUrlById(reportid string) (string, error) {
	rfu := NewReportFileUrlGetRequest()
	rfu.ReportId = reportid
	body, err := a.GetReportFileUrl(rfu)
	if err != nil {
		return "", err
	}
	var rfur ReportFileUrlGetResponse
	err = json.Unmarshal(body, &rfur)
	if err != nil {
		return "", err
	}
	if rfur.Header.Desc != "success" {
		errmsg := fmt.Sprintf("Code:%d Message:%s Position:%s", rfur.Header.Failures[0].Code,
			rfur.Header.Failures[0].Message, rfur.Header.Failures[0].Position)
		return "", errors.New(errmsg)
	}
	if len(rfur.Body.Data) == 0 {
		return "", errors.New("ReportService GetReportFileUrlById data is null")
	}
	return rfur.Body.Data[0].ReportFilePath, nil
}

//getRealTimeData start
type RealTimeDataGetRequest struct {
	RealTimeRequestType *RealTimeRequestType `json:"realTimeRequestType"`
}

func NewRealTimeDataGetRequest() *RealTimeDataGetRequest {
	return new(RealTimeDataGetRequest)
}
func NewRealTimeRequestType() *RealTimeRequestType {
	return new(RealTimeRequestType)
}

type RealTimeDataGetResponse struct {
	Header *ResponseHeader             `json:"header"`
	Body   RealTimeDataGetResponseBody `json:"body"`
}
type RealTimeDataGetResponseBody struct {
	RealTimeResultTypes []*RealTimeResultType `json:"realTimeResultTypes"`
}

//getRealTimeData end

//实时报告请求-包含多类报告
type RealTimeRequestType struct {
	PerformanceData []string       `json:"performanceData"`
	Order           bool           `json:"order"`
	StartDate       string         `json:"startDate"`
	EndDate         string         `json:"endDate"`
	LevelOfDetails  int            `json:"levelOfDetails"`
	Attributes      *AttributeType `json:"attributes"`
	ReportType      int            `json:"reportType"`
	StatIds         []int64        `json:"statIds"`
	StatRange       int            `json:"statRange"`
	UnitOfTime      int            `json:"unitOfTime"`
	Number          int            `json:"number"`
	Device          int            `json:"device"`
	Platform        int            `json:"platform"`
}

//getRealTimeQueryData(搜索词报告) start
type RealTimeQueryDataGetRequest struct {
	RealTimeQueryRequestType *RealTimeQueryRequestType `json:"realTimeQueryRequestType"`
}

func NewRealTimeQueryDataGetRequest() *RealTimeQueryDataGetRequest {
	return new(RealTimeQueryDataGetRequest)
}
func NewRealTimeQueryRequestType() *RealTimeQueryRequestType {
	return new(RealTimeQueryRequestType)
}

type RealTimeQueryDataGetResponse struct {
	Header *ResponseHeader                  `json:"header"`
	Body   RealTimeQueryDataGetResponseBody `json:"body"`
}
type RealTimeQueryDataGetResponseBody struct {
	RealTimeQueryRequestTypes []*RealTimeQueryRequestType `json:"realTimeResultTypes"`
}

//getRealTimeQueryData(搜索词报告) end

//搜索词报告请求
type RealTimeQueryRequestType struct {
	PerformanceData []string       `json:"performanceData"`
	StartDate       string         `json:"startDate"`
	EndDate         string         `json:"endDate"`
	LevelOfDetails  int            `json:"levelOfDetails"`
	Attributes      *AttributeType `json:"attributes"`
	ReportType      int            `json:"reportType"`
	StatIds         []int64        `json:"statIds"`
	StatRange       int            `json:"statRange"`
	UnitOfTime      int            `json:"unitOfTime"`
	Number          int            `json:"number"`
	Device          int            `json:"device"`
}

//getRealTimePairData(配对报告) start
type RealTimePairDataGetRequest struct {
	RealTimePairRequestType *RealTimePairRequestType `json:"realTimePairRequestType"`
}

func NewRealTimePairDataGetRequest() *RealTimePairDataGetRequest {
	return new(RealTimePairDataGetRequest)
}
func NewRealTimePairRequestType() *RealTimePairRequestType {
	return new(RealTimePairRequestType)
}

type RealTimePairDataGetResponse struct {
	Header *ResponseHeader                  `json:"header"`
	Body   RealTimeQueryDataGetResponseBody `json:"body"`
}
type RealTimePairDataGetResponseBody struct {
	RealTimePairRequestTypes []*RealTimePairRequestType `json:"realTimePairRequestTypes"`
}

//getRealTimePairData(配对报告) end

//配对报告请求
type RealTimePairRequestType struct {
	PerformanceData []string       `json:"performanceData"`
	Order           bool           `json:"order"`
	StartDate       string         `json:"startDate"`
	EndDate         string         `json:"endDate"`
	LevelOfDetails  int            `json:"levelOfDetails"`
	Attributes      *AttributeType `json:"attributes"`
	ReportType      int            `json:"reportType"`
	StatIds         []int64        `json:"statIds"`
	StatRange       int            `json:"statRange"`
	UnitOfTime      int            `json:"unitOfTime"`
	Number          int            `json:"number"`
	Device          int            `json:"device"`
}

//getProfessionalReportId(异步报告) start 获取报告ID
type ProfessionalReportIdGetRequest struct {
	ReportRequestType *ReportRequestType `json:"reportRequestType"`
}

func NewProfessionalReportIdGetRequest() *ProfessionalReportIdGetRequest {
	return new(ProfessionalReportIdGetRequest)
}

type ProfessionalReportIdGetResponse struct {
	Header *ResponseHeader                     `json:"header"`
	Body   ProfessionalReportIdGetResponseBody `json:"body"`
}
type ProfessionalReportIdGetResponseBody struct {
	Data []*ProfessionalReportIdType `json:"data"`
}
type ProfessionalReportIdType struct {
	ReportId string `json:"reportId"`
}

//getProfessionalReportId(异步报告) end

// getReportState(异步报告) start 获取报告声称状态
type ReportStateGetRequest struct {
	ReportId string `json:"reportId"`
}

func NewReportStateGetRequest() *ReportStateGetRequest {
	return new(ReportStateGetRequest)
}

type ReportStateGetResponse struct {
	Header *ResponseHeader            `json:"header"`
	Body   ReportStateGetResponseBody `json:"body"`
}
type ReportStateGetResponseBody struct {
	Data []*ReportStateGetResponseType `json:"data"`
}
type ReportStateGetResponseType struct {
	IsGenerated int `json:"isGenerated"`
}

// getReportState(异步报告) end

// getReportFileUrl(异步报告) start 获取报告生成URL
type ReportFileUrlGetRequest struct {
	ReportId string `json:"reportId"`
}

func NewReportFileUrlGetRequest() *ReportFileUrlGetRequest {
	return new(ReportFileUrlGetRequest)
}

type ReportFileUrlGetResponse struct {
	Header *ResponseHeader              `json:"header"`
	Body   ReportFileUrlGetResponseBody `json:"body"`
}
type ReportFileUrlGetResponseBody struct {
	Data []*ReportFileUrlGetResponseType `json:"data"`
}
type ReportFileUrlGetResponseType struct {
	ReportFilePath string `json:"reportFilePath"`
}

// getReportFileUrl(异步报告) end

//异步报告请求
type ReportRequestType struct {
	PerformanceData []string       `json:"performanceData"`
	StartDate       string         `json:"startDate"`
	EndDate         string         `json:"endDate"`
	IdOnly          bool           `json:"idOnly"`
	LevelOfDetails  int            `json:"levelOfDetails"`
	Attributes      *AttributeType `json:"attributes"`
	Format          int            `json:"format"`
	ReportType      int            `json:"reportType"`
	StatIds         []int64        `json:"statIds"`
	StatRange       int            `json:"statRange"`
	UnitOfTime      int            `json:"unitOfTime"`
	Device          int            `json:"device"`
	Platform        int            `json:"platform"`
}

//异步报告
type AttributeType struct {
	Key   string `json:"key"`
	Value []int  `json:"value"`
}

//实时报告返回
type RealTimeResultType struct {
	ID        int64    `json:"ID"`
	Name      []string `json:"name"`
	RelatedId int64    `json:"relatedId"`
	Date      string   `json:"Date"`
	KPIs      []string `json:"KPIs"`
}

//搜索词报告返回
type RealTimeQueryResultType struct {
	Query      string   `json:"query"`
	KeywordId  int64    `json:"keywordId"`
	CreativeId int64    `json:"creativeId"`
	QueryInfo  []string `json:"queryInfo"`
	Date       string   `json:"Date"`
	KPIs       []string `json:"KPIs"`
}

//配对报告返回
type RealTimePairResultType struct {
	keywordId  int64    `json:"keywordId"`
	creativeId int64    `json:"creativeId"`
	PairInfo   []string `json:"pairInfo"`
	Date       string   `json:"Date"`
	KPIs       []string `json:"KPIs"`
}

//account campaingn adgroup keyword creative wordid location searchkey pair newcreative  Location2 history
func (a *ReportService) GetAsyncReportType(asyncReportGetStruct *AsyncReportGetStruct) *ReportRequestType {
	// //通用
	// UnitOfTime      5(日报)
	// IdOnly         true
	// Format          2
	// Attributes:      nil,
	// startDate   > 20010917
	// endDate
	// Device 0，1，2
	r := new(ReportRequestType)
	r.Device = asyncReportGetStruct.Device
	r.StartDate = asyncReportGetStruct.StartDate
	r.EndDate = asyncReportGetStruct.EndDate
	r.IdOnly = true
	r.UnitOfTime = asyncReportGetStruct.UnitOfTime
	r.Format = 2

	switch asyncReportGetStruct.ReportType {
	case AsyncReportAccount:
		// －－－－－账户报告
		// reportType       2
		// levelOfDetails 2
		// statRange 2
		// PerformanceData []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"},

		r.ReportType = 2
		r.LevelOfDetails = 2
		r.StatRange = 2
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"}

		// －－－－－计划报告
		// reportType       10
		// StatRange = 3 // 2,3
		// LevelOfDetails = 3
		// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"}
	case AsyncReportCampaingn:

		r.ReportType = 10
		r.StatRange = 3
		r.LevelOfDetails = 3
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"}

		// －－－－－－单元报告
		// ReportType = 11
		// StatRange = 5 // 2,3,5
		// LevelOfDetails = 5
		// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"}
	case AsyncReportAdgroup:

		r.ReportType = 11
		r.StatRange = 5
		r.LevelOfDetails = 5
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "conversion"}

	// －－－－－－关键词报告 keywordid
	// ReportType = 14
	// StatRange = 11 // 2,3,5,11
	// LevelOfDetails = 11
	// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position","conversion"}
	case AsyncReportKeyword:

		r.ReportType = 14
		r.StatRange = 11
		r.LevelOfDetails = 11
		if asyncReportGetStruct.UnitOfTime == UnitOfTimeHour {
			r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm"}
		} else {
			r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position", "conversion"}
		}
		// －－－－－－创意报告

		// ReportType = 12
		// StatRange = 7 // 2,3,5,7
		// LevelOfDetails = 7
		// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position","conversion"}
	case AsyncReportCreative:

		r.ReportType = 12
		r.StatRange = 7
		r.LevelOfDetails = 7
		if asyncReportGetStruct.UnitOfTime == UnitOfTimeHour {
			r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm"}
		} else {
			r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position", "conversion"}
		}
		// －－－－－－关键词报告（wordid）
		// ReportType = 9
		// StatRange = 6 // 2,3,5,11
		// LevelOfDetails = 6
		// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position","conversion"}
	case AsyncReportWordid:
		r.ReportType = 9
		r.StatRange = 6
		r.LevelOfDetails = 6
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position", "conversion"}

	// －－－－－－－地域报告
	// ReportType = 3
	// LevelOfDetails = 5 // 2,3,5
	// StatRange = 5      // 2,3,5
	// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position","conversion"}
	case AsyncReportLocation:

		r.ReportType = 3
		r.LevelOfDetails = 5
		r.StatRange = 5
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position", "conversion"}

	// －－－－－－－搜索词报告
	// ReportType = 6
	// LevelOfDetails = 3       // 2,3
	// StatRange = 12 // 7,12
	// PerformanceData = []string{"click", "impression"}
	case AsyncReportSearchkey:

		r.ReportType = 6
		r.LevelOfDetails = 3
		r.StatRange = 12
		r.PerformanceData = []string{"click", "impression"}

	// －－－－－－－配对报告
	// ReportType = 15
	// StatRange = 11 //2,3,5,7,11
	// LevelOfDetails = 12
	// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm"}
	case AsyncReportpair:

		r.ReportType = 15
		r.LevelOfDetails = 11
		r.StatRange = 12
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm"}

		// －－－－－－－蹊径报告
	// ReportType = 21
	// StatRange = 2
	// LevelOfDetails = 6
	// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr"}
	case AsyncReportNewCreative:

		r.ReportType = 21
		r.LevelOfDetails = 2
		r.StatRange = 6
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr"}

	// －－－－－－二级地市
	// ReportType = 5
	// LevelOfDetails = 3
	// StatRange = 3
	// PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position","conversion"}
	case AsyncReportLocation2:

		r.ReportType = 5
		r.LevelOfDetails = 3
		r.StatRange = 3
		r.PerformanceData = []string{"cost", "cpc", "click", "impression", "ctr", "cpm", "position", "conversion"}

		// －－－－－－历史排名
	// ReportType = 38
	// StatRange = 11 // 2,11
	// LevelOfDetails = 11
	// PerformanceData = []string{"rank1shows", "rank2shows", "rank3shows", "rank1to3shows", "rank1to3ratio",  "rank1to8shows", "rank1to8ratio","rank4to8shows"}
	// 附注:移动设备不能请求 rank4to8shows 绩效数据
	case AsyncReportHistory:
		var performanceData = []string{"rank1shows", "rank2shows", "rank3shows", "rank1to3shows", "rank1to3ratio",
			"rank1to8shows", "rank1to8ratio", "rank4to8shows"}
		if asyncReportGetStruct.Device == 2 {
			performanceData = performanceData[0 : len(performanceData)-2]

		}
		r.ReportType = 38
		r.LevelOfDetails = 11
		r.StatRange = 11
		r.PerformanceData = performanceData

	}
	return r

}
