// sms_service_ReportService_test.go
package baiduSDK

import (
	// "encoding/json"
	"fmt"
	"testing"
)

func TestGetRealTimeData(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("实时报告－－－－－TestGetRealTimeData start ...")
	fmt.Println("############################################")
	c := NewReportService()
	c.AuthHeader = authHeader

	a := NewRealTimeDataGetRequest()
	r := NewRealTimeRequestType()
	r.StartDate = "2016-03-01"
	r.EndDate = "2016-03-23"
	r.PerformanceData = []string{
		"cost",
		"cpc",
		"click",
		"impression",
		"ctr",
		"cpm",
		// "position",
		"conversion"}
	r.Order = true
	r.ReportType = 11
	r.StatRange = 5
	r.LevelOfDetails = 5
	r.UnitOfTime = 3
	r.Device = 0
	r.Platform = 0
	r.Number = 5
	a.RealTimeRequestType = r

	// a.RealTimeRequestType.StartDate = "2016-03-01"
	// a.RealTimeRequestType.EndDate = "2016-03-23"

	body, err := c.GetRealTimeData(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	// var car OperationRecordGetResponse
	// err = json.Unmarshal(body, &car)
	// if err != nil {
	// 	fmt.Println("json Unmarshal err:", err)
	// }

	// //获取返回数据状态
	// if ok, message := car.AuthHeader.GetResponseStatus(); !ok {
	// 	fmt.Println("failures Message:", message)
	// 	return
	// }

	// for _, v := range car.Body.OptLogTypes {
	// 	fmt.Println("OperationRecord  is :", v)
	// }
	fmt.Println("############################################")

}
func TestGetRealTimePairData(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("配对报告－TestGetRealTimePairData start ...")
	fmt.Println("############################################")
	c := NewReportService()
	c.AuthHeader = authHeader

	a := NewRealTimePairDataGetRequest()
	r := NewRealTimePairRequestType()
	r.StartDate = "2016-03-01"
	r.EndDate = "2016-03-23"
	r.PerformanceData = []string{
		"cost",
		"cpc",
		"click",
		"impression",
		"ctr",
		"cpm",
		// "position",
		// "conversion"
	}
	// r.Order = true
	r.ReportType = 15
	r.StatRange = 3
	r.LevelOfDetails = 12
	r.UnitOfTime = 3
	r.Device = 0
	// r.Platform = 0
	r.Number = 5
	a.RealTimePairRequestType = r

	// a.RealTimeRequestType.StartDate = "2016-03-01"
	// a.RealTimeRequestType.EndDate = "2016-03-23"

	body, err := c.GetRealTimePairData(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))

	// var car OperationRecordGetResponse
	// err = json.Unmarshal(body, &car)
	// if err != nil {
	// 	fmt.Println("json Unmarshal err:", err)
	// }

	// //获取返回数据状态
	// if ok, message := car.AuthHeader.GetResponseStatus(); !ok {
	// 	fmt.Println("failures Message:", message)
	// 	return
	// }

	// for _, v := range car.Body.OptLogTypes {
	// 	fmt.Println("OperationRecord  is :", v)
	// }
	fmt.Println("############################################")

}
func TestGetAsyncReportType(t *testing.T) {
	//c := NewReportService()
	//fmt.Println(c.GetAsyncReportType(AsyncReportAccount, "", "", 0))

}
