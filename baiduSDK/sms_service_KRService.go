// sms_service_KRService.go
package baiduSDK

//关键词规划师

const (
	GET_KRService_EstimatedDataByBid = "getEstimatedDataByBid"
	GET_KRService_EstimatedData      = "getEstimatedData"
	GET_KRService_KRFileIdByWords    = "getKRFileIdByWords"
	GET_KRService_FilePath           = "getFilePath"
	GET_KRService_FileStatus         = "getFileStatus"
	GET_KRService_KRByQuery          = "getKRByQuery"
	GET_KRService_KRCustom           = "getKRCustom"
	GET_KRService_BidByWords         = "getBidByWords"
)

type KRService struct {
	*CommonService
}

func NewKRService() *KRService {
	a := new(KRService)
	a.CommonService = NewCommonService("sms", "service", "KRService")
	return a
}

func (a *KRService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)

}

func (a *KRService) GetEstimatedDataByBid(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_EstimatedDataByBid, request)
}
func (a *KRService) GetEstimatedData(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_EstimatedData, request)
}

func (a *KRService) GetKRFileIdByWords(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_KRFileIdByWords, request)
}

func (a *KRService) GetFilePath(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_FilePath, request)
}

func (a *KRService) GetFileStatus(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_FileStatus, request)
}

func (a *KRService) GetKRByQuery(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_KRByQuery, request)
}

func (a *KRService) GetKRCustom(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_KRCustom, request)
}

func (a *KRService) GetBidByWords(request interface{}) ([]byte, error) {
	return a.Execute(GET_KRService_BidByWords, request)
}

type SeedFilter struct {
}
