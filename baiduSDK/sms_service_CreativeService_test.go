// sms_service_CreativeService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllCreative(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllCreative start ...")
	fmt.Println("############################################")
	c := NewCreativeService()
	c.AuthHeader = authHeader
	a := NewCreativeGetRequest()
	fields := []string{
		"creativeId",
		"adgroupId",
		"title",
		"description1",
		"description2",
		"pcDestinationUrl",
		"pcDisplayUrl",
		"mobileDestinationUrl",
		"mobileDisplayUrl",
		"pause",
		"status",
		"devicePreference"}

	a.CreativeFields = fields
	a.IdType = 5
	a.Ids = []int64{1696828587}

	body, err := c.GetCreative(a)
	if err != nil {
		t.Error(err)
	}
	var car CreativeGetResponse
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
	if len(car.Body.CreativeTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	for _, v := range car.Body.CreativeTypes {
		fmt.Println("Creative is :", v)
	}

	fmt.Println("############################################")

}
func TestGetCreative(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetCreative start ...")
	fmt.Println("############################################")
	c := NewCreativeService()
	c.AuthHeader = authHeader
	a := NewCreativeGetRequest()
	fields := []string{
		"creativeId",
		"adgroupId",
		"title",
		"description1",
		"description2",
		"pcDestinationUrl",
		"pcDisplayUrl",
		"mobileDestinationUrl",
		"mobileDisplayUrl",
		"pause",
		"status",
		"devicePreference"}

	a.CreativeFields = fields
	a.IdType = 7
	a.Ids = []int64{10431859733}

	body, err := c.GetCreative(a)
	if err != nil {
		t.Error(err)
	}
	var car CreativeGetResponse
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
	if len(car.Body.CreativeTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	fmt.Println("Creative is :", car.Body.CreativeTypes[0])

	fmt.Println("############################################")

}
