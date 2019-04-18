// sms_service_KeywordService.go
package baiduSDK

// 关键词服务KeywordService：

// 通过增加、删除、修改、查询关键词的信息，用于管理账户下任意的关键词；
// 例如，您可以批量获取和修改某些关键词的出价、PC访问URL、移动访问URL、匹配模式、暂停/启用、激活状态等属性。
import ()

const (
	GET_Keyword    = "getWord"
	Add_Keyword    = "addWord"
	UPDATE_Keyword = "updateWord"
	DELETE_Keyword = "deleteWord"
)

type KeywordService struct {
	*CommonService
}

func NewKeywordService() *KeywordService {
	a := new(KeywordService)
	a.CommonService = NewCommonService("sms", "service", "KeywordService")
	return a
}
func (a *KeywordService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)
}

func (a *KeywordService) GetKeyword(request interface{}) ([]byte, error) {
	return a.Execute(GET_Keyword, request)
}
func (a *KeywordService) AddKeyword(request interface{}) ([]byte, error) {
	return a.Execute(Add_Keyword, request)
}
func (a *KeywordService) UpdateKeyword(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_Keyword, request)
}
func (a *KeywordService) DeleteKeyword(request interface{}) ([]byte, error) {
	return a.Execute(DELETE_Keyword, request)
}

//定义获取关键字请求结构
type KeywordGetRequest struct {
	Ids        []int64  `json:"ids"`
	IdType     int      `json:"idType"`
	GetTemp    int      `json:"getTemp"`
	WordFields []string `json:"wordFields"`
}

func NewKeywordGetRequest() *KeywordGetRequest {
	c := new(KeywordGetRequest)
	return c
}

//定义获取关键字响应结构
type KeywordGetResponse struct {
	Header *ResponseHeader     `json:"header"`
	Body   KeywordResponseBody `json:"body"`
}
type KeywordResponseBody struct {
	KeywordTypes []*KeywordType `json:"data"`
}

//定义添加关键字请求结构
type KeywordAddRequest struct {
	KeywordTypes []*KeywordType `json:"keywordTypes"`
}

func NewKeywordAddRequest() *KeywordAddRequest {
	c := new(KeywordAddRequest)
	return c
}

//定义添加关键字响应结构
type KeywordAddResponse struct {
	KeywordTypes []*KeywordType `json:"data"`
}

//定义更新关键字请求结构
type KeywordUpdateRequest struct {
	KeywordTypes []*KeywordType `json:"keywordTypes"`
}

func NewKeywordUpdateRequest() *KeywordUpdateRequest {
	c := new(KeywordUpdateRequest)
	return c
}

//定义更新关键字响应结构
type KeywordUpdateResponse struct {
	KeywordTypes []*KeywordType `json:"data"`
}

//定义删除关键字请求结构
type KeywordDeleteRequest struct {
	KeywordIds []int64 `json:"keywordIds"`
}

func NewKeywordDeleteRequest() *KeywordDeleteRequest {
	c := new(KeywordDeleteRequest)
	return c
}

//定义删除关键字响应结构
type KeywordDeleteResponse struct {
	KeywordTypes []*KeywordType `json:"data"`
}

//关键词服务
type KeywordType struct {
	KeywordId            int64   `json:"keywordId"`
	AdgroupId            int64   `json:"AdgroupId"`
	CampaignId           int64   `json:"campaignId"`
	Keyword              string  `json:"keyword"`
	Price                float64 `json:"price"`
	PcDestinationUrl     string  `json:"pcDestinationUrl"`
	MobileDestinationUrl string  `json:"mobileDestinationUrl"`
	MatchType            int     `json:"matchType"`
	Pause                bool    `json:"pause"`
	Status               int     `json:"status"`
	PhraseType           int     `json:"phraseType"`
	wmatchprefer         int     `json:"wmatchprefer"`
	PcQuality            int64   `json:"pcQuality"`
	PcReliable           int     `json:"pcReliable"`
	PcReason             int     `json:"pcReason"`
	PcScale              []int   `json:"pcScale"`
	MobileQuality        int64   `json:"mobileQuality"`
	MobileReliable       int     `json:"mobileReliable"`
	MobileReason         int     `json:"mobileReason"`
	MobileScale          []int   `json:"mobileScale"`
}

var KeywordFields = map[string]string{
	"keywordId":            "关键词 id",
	"AdgroupId":            "推广单元 id",
	"campaignId":           "推广计划 id",
	"keyword":              "关键词字面",
	"price":                "关键词竞价价格",
	"pcDestinationUrl":     "目标 url",
	"mobileDestinationUrl": "移动访问 URL",
	"matchType":            "匹配模式",
	"pause":                "暂停/启用关键词",
	"status":               "关键词状态",
	"phraseType":           "高级短语细分匹 配模式",
	"wmatchprefer":         "是否接受单元的 分匹配出价比例",
	"pcQuality":            "PC 上该关键词的 10 分质量度",
	"pcReliable":           "PC 上新广告维度 下是否为临时质 量度",
	"pcReason":             "质量度原因",
	"pcScale":              "PC 竞争关系分布",
	"mobileQuality":        "无线上该关键词 的 10 分质量度",
	"mobileReliable":       "无线上新广告维 度下是否为临时 质量度",
	"mobileReason":         "质量度原因",
	"mobileScale":          "无线上竞争关系 分布",
}
