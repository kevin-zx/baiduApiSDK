//sms_service_NewCreativeService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllSublink(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllSublink start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewSublinkGetRequest()
	fields := []string{
		"sublinkId",
		"sublinkInfos",
		"adgroupId",
		"pause",
		"status",
	}

	a.SublinkFields = fields
	a.IdType = 5
	a.Ids = []int64{1231821855}

	body, err := c.GetSublink(a)
	if err != nil {
		t.Error(err)
	}

	var car SublinkGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(car.Body.SublinkTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	for _, v := range car.Body.SublinkTypes {
		fmt.Println("SublinkType AdgroupId  is :", v.AdgroupId)
		fmt.Println("SublinkType SublinkId  is :", v.SublinkId)
		fmt.Println("SublinkType Device is :", v.Device)
		for _, s := range v.SublinkInfos {
			fmt.Println("SublinkInfo  is :", s)
		}
	}

	fmt.Println("############################################")

}

func TestGetSublink(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetSublink start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewSublinkGetRequest()
	fields := []string{
		"sublinkId",
		"sublinkInfos",
		"adgroupId",
		"pause",
		"status",
	}

	a.SublinkFields = fields
	a.IdType = 12
	a.Ids = []int64{1307750769}

	body, err := c.GetSublink(a)
	if err != nil {
		t.Error(err)
	}

	var car SublinkGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}

	fmt.Println("SublinkType AdgroupId  is :", car.Body.SublinkTypes[0].AdgroupId)
	fmt.Println("SublinkType SublinkId  is :", car.Body.SublinkTypes[0].SublinkId)
	fmt.Println("SublinkType Device is :", car.Body.SublinkTypes[0].Device)
	for _, s := range car.Body.SublinkTypes[0].SublinkInfos {
		fmt.Println("SublinkInfo  is :", s)
	}

	fmt.Println("############################################")

}

func TestGetAllBridge(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("商桥－－TestGetAllBridge start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewBridgeGetRequest()
	fields := []string{
		"bridgeId",
		"adgroupId",
		"pause",
		"status",
	}

	a.BridgeFields = fields
	a.IdType = 5
	a.Ids = []int64{
		1715190897,
	}

	body, err := c.GetBridge(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car BridgeGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(car.Body.BridgeTypes) == 0 {
		fmt.Println("account response body is null")
		return
	}
	for _, v := range car.Body.BridgeTypes {
		fmt.Println("BridgeType AdgroupId  is :", v.AdgroupId)
		fmt.Println("BridgeType BridgeId  is :", v.BridgeId)
	}

	fmt.Println("############################################")

}

func TestGetBridge(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("商桥－－TestGetBridge start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewBridgeGetRequest()
	fields := []string{
		"bridgeId",
		"adgroupId",
		"pause",
		"status",
	}

	a.BridgeFields = fields
	a.IdType = 10
	a.Ids = []int64{
		1715190897,
	}

	body, err := c.GetBridge(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car BridgeGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(car.Body.BridgeTypes) == 0 {
		fmt.Println("account response body is null")
		return
	}
	fmt.Println("BridgeType AdgroupId  is :", car.Body.BridgeTypes[0].AdgroupId)
	fmt.Println("BridgeType BridgeId  is :", car.Body.BridgeTypes[0].BridgeId)

	fmt.Println("############################################")

}

func TestGetAllPhone(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllPhone start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewPhoneGetRequest()
	fields := []string{
		"phoneId",
		"phoneNum",
		"adgroupId",
		"pause",
		"status",
	}

	a.PhoneFields = fields
	a.IdType = 5
	a.Ids = []int64{
		1715190897,
		1231821870,
		1231821990,
		1231821864,
		1715190900,
		1231821984,
		1715190903,
		1231821987,
		1231821867,
		1715782791,
		1231821996,
		1231821861,
		1231821999,
		1715190906,
		1715190909,
		1231821993,
		1231821858,
		1231821885,
		1231822005,
		1703146212,
		1231822002,
		1703146209,
		1231821882,
		1231821876,
		1231822014,
		1231821879,
		1231822008,
		1231821873,
		1231822011,
		1231821822,
		1231821957,
		1231821837,
		1231821834,
		1231821954,
		1231821819,
		1231821966,
		1231821813,
		1231821963,
		1231821810,
		1231821960,
		1231821855,
		1231821807,
		1231821975,
		1231821972,
		1231822095,
		1231821852,
		1231821804,
		1231822089,
		1231821801,
		1231821969,
		1231821798,
		1231821981,
		1231821978,
		1696828566,
		1231822050,
		1231822053,
		1696828572,
		1231822056,
		1660784725,
		1231822059,
		1696828575,
		1696828569,
		1231822062,
		1231822065,
		1231822068,
		1231822071,
		1231822074,
		1231822077,
		1231822017,
		1720779416,
		1231822023,
		1231822020,
		1720779413,
		1231822026,
		1715190930,
		1720779410,
		1231822029,
		1715190927,
		1720779404,
		1231822035,
		1720779407,
		1696828581,
		1231822032,
		1720779401,
		1231822038,
		1696828578,
		1231822041,
		1696828587,
		1231822047,
		1231822044,
		1696828584,
	}

	body, err := c.GetPhone(a)
	if err != nil {
		t.Error(err)
	}
	var car PhoneGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(car.Body.PhoneTypes) == 0 {
		fmt.Println("account response body is null")
		return
	}
	for _, v := range car.Body.PhoneTypes {
		fmt.Println("PhoneType   is :", v)
	}

	fmt.Println("############################################")

}
func TestGetPhone(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetPhone start ...")
	fmt.Println("############################################")
	c := NewNewCreativeService()
	c.AuthHeader = authHeader
	a := NewPhoneGetRequest()
	fields := []string{
		"phoneId",
		"phoneNum",
		"adgroupId",
		"pause",
		"status",
	}

	a.PhoneFields = fields
	a.IdType = 9
	a.Ids = []int64{1722221156}

	body, err := c.GetPhone(a)
	if err != nil {
		t.Error(err)
	}
	var car PhoneGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(car.Body.PhoneTypes) == 0 {
		fmt.Println("account response body is null")
		return
	}
	for _, v := range car.Body.PhoneTypes {
		fmt.Println("PhoneType is :", v)
	}

	fmt.Println("############################################")

}
