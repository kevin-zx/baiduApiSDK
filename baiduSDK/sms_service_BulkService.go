// sms_service_BulkService.go
package baiduSDK

import (
	"encoding/json"
	"errors"
	"fmt"
	// "time"
)

//使用此服务可以将对您的大批量请求安排为异步批量请求,并可以获取您近期所执行请求的状态、 结果等信息。
const (
	GET_AllChangedObjects = "getAllChangedObjects"
	CANCEL_Download       = "cancelDownload"
	GET_ChangedItemId     = "getChangedItemId"
	GET_ChangedScale      = "getChangedScale"
	GET_AllObjects        = "getAllObjects"
	GET_UserCache         = "getUserCache"
	GET_FilePath          = "getFilePath"
	GET_FileStatus        = "getFileStatus"
	GET_ChangedId         = "getChangedId"
)

type BulkService struct {
	*CommonService
}

func NewBulkService() *BulkService {
	a := new(BulkService)
	a.CommonService = NewCommonService("sms", "service", "BulkJobService")
	return a
}

func (a *BulkService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)

}

func (a *BulkService) GetAllChangedObjects(request interface{}) ([]byte, error) {
	return a.Execute(GET_AllChangedObjects, request)
}

func (a *BulkService) CancelDownload(request interface{}) ([]byte, error) {
	return a.Execute(CANCEL_Download, request)
}
func (a *BulkService) GetChangedItemId(request interface{}) ([]byte, error) {
	return a.Execute(GET_ChangedItemId, request)
}
func (a *BulkService) GetChangedScale(request interface{}) ([]byte, error) {
	return a.Execute(GET_ChangedScale, request)
}
func (a *BulkService) GetAllObjects(request interface{}) ([]byte, error) {
	return a.Execute(GET_AllObjects, request)
}
func (a *BulkService) GetUserCache(request interface{}) ([]byte, error) {
	return a.Execute(GET_UserCache, request)
}
func (a *BulkService) GetFilePath(request interface{}) ([]byte, error) {
	return a.Execute(GET_FilePath, request)
}
func (a *BulkService) GetFileStatus(request interface{}) ([]byte, error) {
	return a.Execute(GET_FileStatus, request)
}
func (a *BulkService) GetChangedId(request interface{}) ([]byte, error) {
	return a.Execute(GET_ChangedId, request)
}

//getAllObjects start
type AllObjectGetRequest struct {
	CampaignIds           []int64  `json:"campaignIds"`
	IncludeTemp           bool     `json:"includeTemp"`
	Format                int      `json:"format"`
	AccountFields         []string `json:"accountFields"`
	CampaignFields        []string `json:"campaignFields"`
	AdgroupFields         []string `json:"adgroupFields"`
	KeywordFields         []string `json:"keywordFields"`
	CreativeFields        []string `json:"creativeFields"`
	SublinkFields         []string `json:"sublinkFields"`
	MobileSublinkFields   []string `json:"mobileSublinkFields"`
	PhoneFields           []string `json:"phoneFields"`
	BridgeFields          []string `json:"bridgeFields"`
	DynamicCreativeFields []string `json:"dynamicCreativeFields"`
}

func NewAllObjectGetRequest() *AllObjectGetRequest {
	return &AllObjectGetRequest{
		IncludeTemp:           true,
		Format:                1,
		AccountFields:         []string{"all"},
		CampaignFields:        []string{"all"},
		AdgroupFields:         []string{"all"},
		KeywordFields:         []string{"all"},
		CreativeFields:        []string{"all"},
		SublinkFields:         []string{"all"},
		MobileSublinkFields:   []string{"all"},
		PhoneFields:           []string{"all"},
		BridgeFields:          []string{"all"},
		DynamicCreativeFields: []string{"all"},
	}
}

type AllObjectsGetResponse struct {
	Header *ResponseHeader            `json:"header"`
	Body   *AllObjectsGetResponseBody `json:"body"`
}
type AllObjectsGetResponseBody struct {
	AllObjectsFileIdGetResponse []*AllObjectsFileIdGetResponse `json:"data"` //整账户下载文件 id
}
type AllObjectsFileIdGetResponse struct {
	FileId string `json:"fileId"` //整账户下载文件 idfileId
}

//getAllObjects end

//getAllChangedObjects start
type AllChangedObjectsGetRequest struct {
	StartTime             string   `json:"startTime"`
	CampaignIds           []int64  `json:"campaignIds"`
	IncludeTemp           bool     `json:"includeTemp"`
	Format                int      `json:"format"`
	CampaignFields        []string `json:"campaignFields"`
	AdgroupFields         []string `json:"adgroupFields"`
	KeywordFields         []string `json:"keywordFields"`
	CreativeFields        []string `json:"creativeFields"`
	SublinkFields         []string `json:"sublinkFields"`
	MobileSublinkFields   []string `json:"mobileSublinkFields"`
	PhoneFields           []string `json:"phoneFields"`
	BridgeFields          []string `json:"bridgeFields"`
	DynamicCreativeFields []string `json:"dynamicCreativeFields"`
}

func NewAllChangedObjectsGetRequest() *AllChangedObjectsGetRequest {
	return &AllChangedObjectsGetRequest{
		CampaignIds: []int64{},
		IncludeTemp: true,
		Format:      1,
	}
}

type AllChangedObjectsGetResponse struct {
	Header *ResponseHeader                   `json:"header"`
	Body   *AllChangedObjectsGetResponseBody `json:"body"`
}
type AllChangedObjectsGetResponseBody struct {
	AllChangedObjectsFileIdGetResponse []*AllChangedObjectsFileIdGetResponse `json:"data"`
}
type AllChangedObjectsFileIdGetResponse struct {
	FileId string `json:"fileId"`
}

//getAllChangedObjects end

//getFileStatus start
type FileStateGetRequest struct {
	FileId string `json:"fileId"`
}

func NewFileStateGetRequest() *FileStateGetRequest {
	return new(FileStateGetRequest)
}

type FileStateGetResponse struct {
	Header *ResponseHeader           `json:"header"`
	Body   *FileStateGetResponseBody `json:"body"`
}
type FileStateGetResponseBody struct {
	FileStateIsGeneratedGetResponse []*FileStateIsGeneratedGetResponse `json:"data"`
}
type FileStateIsGeneratedGetResponse struct {
	IsGenerated int `json:"isGenerated"`
}

//getFileStatus end

//getFilePath start
type FilePathGetRequest struct {
	FileId string `json:"fileId"`
}

func NewFilePathGetRequest() *FilePathGetRequest {
	return new(FilePathGetRequest)
}

type FilePathGetResponse struct {
	Header *ResponseHeader          `json:"header"`
	Body   *FilePathGetResponseBody `json:"body"`
}
type FilePathGetResponseBody struct {
	FilePathsGetResponse []*FilePathType `json:"data"`
}

//getFilePath end
// cancelDownload start
type DownloadCancelRequest struct {
	FileId string `json:"fileId"`
}

func NewDownloadCancelRequest() *DownloadCancelRequest {
	return new(DownloadCancelRequest)
}

type DownloadCancelResponse struct {
	Header *ResponseHeader             `json:"header"`
	Body   *DownloadCancelResponseBody `json:"body"`
}
type DownloadCancelResponseBody struct {
	DownloadIsCanceledCancelResponse []*DownloadIsCanceledCancelResponse `json:"data"`
}
type DownloadIsCanceledCancelResponse struct {
	IsCanceled int `json:"isCanceled"`
}

// cancelDownload end

// getChangedId start
type ChangedIdGetRequest struct {
	StartTime            string `json:"startTime"`
	CampaignLevel        bool   `json:"campaignLevel"`
	AdgroupLevel         bool   `json:"adgroupLevel"`
	KeywordLevel         bool   `json:"keywordLevel"`
	CreativeLevel        bool   `json:"creativeLevel"`
	SublinkLevel         bool   `json:"sublinkLevel"`
	MobileSublinkLevel   bool   `json:"mobileSublinkLevel"`
	PhoneLevel           bool   `json:"phoneLevel"`
	BridgeLevel          bool   `json:"bridgeLevel"`
	DynamicCreativeLevel bool   `json:"dynamicCreativeLevel"`
}

func NewChangedIdGetRequest() *ChangedIdGetRequest {
	return &ChangedIdGetRequest{
		CampaignLevel:        true,
		AdgroupLevel:         true,
		KeywordLevel:         true,
		CreativeLevel:        true,
		SublinkLevel:         true,
		MobileSublinkLevel:   true,
		PhoneLevel:           true,
		BridgeLevel:          true,
		DynamicCreativeLevel: true,
	}
}

type ChangedIdGetResponse struct {
	Header *ResponseHeader           `json:"header"`
	Body   *ChangedIdGetResponseBody `json:"body"`
}

type ChangedIdGetResponseBody struct {
	ChangedIdType []*ChangedIdType `json:"data"`
}
type ChangedIdType struct {
	EndTime                   string               `json:"endTime"`
	ChangedCampaignIds        []*ChangedItemIdType `json:"changedCampaignIds"`
	ChangedAdgroupIds         []*ChangedItemIdType `json:"changedAdgroupIds"`
	ChangedKeywordIds         []*ChangedItemIdType `json:"changedKeywordIds"`
	ChangedCreativeIds        []*ChangedItemIdType `json:"changedCreativeIds"`
	ChangedSublinkIds         []*ChangedItemIdType `json:"changedSublinkIds"`
	ChangedMobileSublinkIds   []*ChangedItemIdType `json:"changedMobileSublinkIds"`
	ChangedPhoneIds           []*ChangedItemIdType `json:"changedPhoneIds"`
	ChangedBridgeIds          []*ChangedItemIdType `json:"changedBridgeIds"`
	ChangedDynamicCreativeIds []*ChangedItemIdType `json:"changedDynamicCreativeIds"`
}

// getChangedId end
// getChangedItemId start
//获取从指定时间到当前时间段内有变化的物料 id。各层级除计划层级返回的数据限制不超 过两万条。
type ChangedItemIdGetRequest struct {
	StartTime            string  `json:"startTime"`
	ItemType             int     `json:"itemType"`
	Type                 int     `json:"type"`
	Ids                  []int64 `json:"Ids"`
	CampaignLevel        bool    `json:"campaignLevel"`
	AdgroupLevel         bool    `json:"adgroupLevel"`
	KeywordLevel         bool    `json:"keywordLevel"`
	CreativeLevel        bool    `json:"creativeLevel"`
	SublinkLevel         bool    `json:"sublinkLevel"`
	MobileSublinkLevel   bool    `json:"mobileSublinkLevel"`
	PhoneLevel           bool    `json:"phoneLevel"`
	BridgeLevel          bool    `json:"bridgeLevel"`
	DynamicCreativeLevel bool    `json:"dynamicCreativeLevel"`
}

func NewChangedItemIdGetRequest() *ChangedItemIdGetRequest {
	return &ChangedItemIdGetRequest{
		CampaignLevel:        true,
		AdgroupLevel:         true,
		KeywordLevel:         true,
		CreativeLevel:        true,
		SublinkLevel:         true,
		MobileSublinkLevel:   true,
		PhoneLevel:           true,
		BridgeLevel:          true,
		DynamicCreativeLevel: true,
	}
}

type ChangedItemIdGetResponse struct {
	Header *ResponseHeader               `json:"header"`
	Body   *ChangedItemIdGetResponseBody `json:"body"`
}
type ChangedItemIdGetResponseBody struct {
	ChangedItemIdDataTypes []*ChangedItemIdDataType `json:"data"`
}

type ChangedItemIdDataType struct {
	EndTime        string               `json:"endTime"`
	ChangedItemIds []*ChangedItemIdType `json:"changedItemIds"`
}

// getChangedItemId end

// getChangedScale start
//获取完整账户下,或者指定计划 ID 下的变化物料规模,从而帮助用户决定后续的最近更新
// 策略。当有变化的物料规模大于一定比例时,用户不妨选择整账户下载。
// 注:变化物料仅指因用户操作发生的变化,即在历史操作记录中可以查询到的操作。而对于
// 质量度,状态这些在系统内自动发生的改变不在统计范围内。
type ChangedScaleGetRequest struct {
	StartTime                   string  `json:"startTime"`
	CampaignIds                 []int64 `json:"campaignIds"`
	ChangedCampaignScale        bool    `json:"changedCampaignScale"`
	ChangedAdgroupScale         bool    `json:"changedAdgroupScale"`
	ChangedKeywordScale         bool    `json:"changedKeywordScale"`
	ChangedCreativeScale        bool    `json:"changedCreativeScale"`
	ChangedSublinkScale         bool    `json:"changedSublinkScale"`
	ChangedMoblieSublinkScale   bool    `json:"changedMoblieSublinkScale"`
	ChangedPhoneScale           bool    `json:"changedPhoneScale"`
	ChangedBridgeScale          bool    `json:"changedBridgeScale"`
	ChangedDynamicCreativeScale bool    `json:"changedDynamicCreativeScale"`
}

func NewChangedScaleGetRequest() *ChangedScaleGetRequest {
	return &ChangedScaleGetRequest{
		ChangedCampaignScale:        true,
		ChangedAdgroupScale:         true,
		ChangedKeywordScale:         true,
		ChangedCreativeScale:        true,
		ChangedSublinkScale:         true,
		ChangedMoblieSublinkScale:   true,
		ChangedPhoneScale:           true,
		ChangedBridgeScale:          true,
		ChangedDynamicCreativeScale: true,
	}
}

type ChangedScaleGetResponse struct {
	Header *ResponseHeader              `json:"header"`
	Body   *ChangedScaleGetResponseBody `json:"body"`
}
type ChangedScaleGetResponseBody struct {
	ChangedScaleDataTypes []*ChangedScaleDataType `json:"data"`
}

type ChangedScaleDataType struct {
	ChangedCampaignScale        []int64 `json:"changedCampaignScale"`
	ChangedAdgroupScale         []int64 `json:"changedAdgroupScale"`
	ChangedKeywordScale         []int64 `json:"changedKeywordScale"`
	ChangedCreativeScale        []int64 `json:"changedCreativeScale"`
	ChangedSublinkScale         []int64 `json:"changedSublinkScale"`
	ChangedMoblieSublinkScale   []int64 `json:"changedMoblieSublinkScale"`
	ChangedPhoneScale           []int64 `json:"changedPhoneScale"`
	ChangedBridgeScale          []int64 `json:"changedBridgeScale"`
	ChangedDynamicCreativeScale []int64 `json:"changedDynamicCreativeScale"`
}

// getChangedScale end

type FilePathType struct {
	AccountFilePath         string `json:"accountFilePath"`
	AccountFileMd5          string `json:"accountFileMd5"`
	CampaignFilePath        string `json:"campaignFilePath"`
	CampaignFileMd5         string `json:"campaignFileMd5"`
	AdgroupFilePath         string `json:"adgroupFilePath"`
	AdgroupFileMd5          string `json:"adgroupFileMd5"`
	KeywordFilePath         string `json:"keywordFilePath"`
	KeywordFileMd5          string `json:"keywordFileMd5"`
	CreativeFilePath        string `json:"creativeFilePath"`
	CreativeFileMd5         string `json:"creativeFileMd5"`
	SublinkFilePath         string `json:"sublinkFilePath"`
	SublinkFileMd5          string `json:"sublinkFileMd5"`
	MobileSublinkFilePath   string `json:"mobileSublinkFilePath"`
	MobileSublinkFileMd5    string `json:"mobileSublinkFileMd5"`
	PhoneFilePath           string `json:"phoneFilePath"`
	PhoneFileMd5            string `json:"phoneFileMd5"`
	BridgeFilePath          string "bridgeFilePath"
	BridgeFileMd5           string "bridgeFileMd5"
	DynamicCreativeFilePath string "dynamicCreativeFilePath"
	DynamicCreativeFileMd5  string "dynamicCreativeFileMd5"
}
type ChangedItemIdType struct {
	Operator      int   `json:"operator"`
	CreativeId    int64 `json:"creativeId"`
	DynCreativeId int64 `json:"dynCreativeId"`
	KeywordId     int64 `json:"keywordId"`
	AdgroupId     int64 `json:"adgroupId"`
	CampaignId    int64 `json:"campaignId"`
}

//获取异步ID GetAllObjects
func (a *BulkService) GetAllObjectsId() (string, error) {
	aog := NewAllObjectGetRequest()
	var (
		body []byte
		err  error
	)
	body, err = a.GetAllObjects(aog)
	if err != nil {
		return "", err

	}
	var aogid AllObjectsGetResponse
	err = json.Unmarshal(body, &aogid)
	if err != nil {
		return "", err
	}
	fmt.Println(aogid.Body)
	fmt.Println(aogid.Header)
	for _, f := range aogid.Header.Failures {
		fmt.Println(f)
	}
	if aogid.Header.Desc != "success" {
		return "", errors.New("BulkService GetAllObjectsId AuthHeader Desc failed")
	}
	if len(aogid.Body.AllObjectsFileIdGetResponse) == 0 {
		return "", errors.New("BulkService GetAllObjectsId AllObjectsFileIdGetResponse Is null")
	}
	return aogid.Body.AllObjectsFileIdGetResponse[0].FileId, nil
}

//通过异步ID，获取状态

func (a *BulkService) GetFileStatusById(buikId string) (bool, error) {
	s := NewFileStateGetRequest()
	s.FileId = buikId
	body, err := a.GetFileStatus(s)
	if err != nil {
		return false, err
	}
	var fsg FileStateGetResponse
	err = json.Unmarshal(body, &fsg)
	if err != nil {
		return false, err
	}
	// fmt.Println(fsg)
	if fsg.Header.Desc != "success" {
		return false, errors.New("BulkService GetFileStatusById  AuthHeader Desc Failed")
	}
	if len(fsg.Body.FileStateIsGeneratedGetResponse) == 0 {
		return false, errors.New("BulkService GetFileStatusById  FileStateIsGeneratedGetResponse Is null")
	}
	if fsg.Body.FileStateIsGeneratedGetResponse[0].IsGenerated == 3 {
		return true, nil
	}
	return false, nil
}

func (a *BulkService) GetFilePathById(buikId string) (*FilePathType, error) {
	rfu := NewFilePathGetRequest()
	rfu.FileId = buikId
	body, err := a.GetFilePath(rfu)
	if err != nil {
		return nil, err
	}
	var rfur FilePathGetResponse
	err = json.Unmarshal(body, &rfur)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	if rfur.Header.Desc != "success" {
		return nil, errors.New("BulkService GetFilePathById  AuthHeader Desc Failed")
	}
	if len(rfur.Body.FilePathsGetResponse) == 0 {
		return nil, errors.New("BulkService GetFilePathById  FilePathsGetResponse Is null")
	}
	return rfur.Body.FilePathsGetResponse[0], nil
}
