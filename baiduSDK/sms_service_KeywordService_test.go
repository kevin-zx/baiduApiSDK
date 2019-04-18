// sms_service_KeywordService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllKeyword(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllKeyword start ...")
	fmt.Println("############################################")
	c := NewKeywordService()
	c.AuthHeader = authHeader
	a := NewKeywordGetRequest()
	fields := []string{
		"keywordId",
		"adgroupId",
		"campaignId",
		"keyword",
		"price",
		"pcDestinationUrl",
		"mobileDestinationUrl",
		"matchType",
		"pause",
		"status",
		"phraseType",
		"wmatchprefer",
		"pcQuality",
		"pcScale",
		"mobileQuality",
		"mobileScale"}

	a.WordFields = fields
	a.IdType = 5
	a.Ids = []int64{1231822026}

	body, err := c.GetKeyword(a)
	if err != nil {
		t.Error(err)
	}
	var car KeywordGetResponse
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
	if len(car.Body.KeywordTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	for _, v := range car.Body.KeywordTypes {
		fmt.Println("Keyword is :", v)
	}

	fmt.Println("############################################")

}

func TestGetKeyword(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetKeyword start ...")
	fmt.Println("############################################")
	c := NewKeywordService()
	c.AuthHeader = authHeader
	a := NewKeywordGetRequest()
	fields := []string{
		"keywordId",
		"adgroupId",
		"campaignId",
		"keyword",
		"price",
		"pcDestinationUrl",
		"mobileDestinationUrl",
		"matchType",
		"pause",
		"status",
		"phraseType",
		"wmatchprefer",
		"pcQuality",
		"pcScale",
		"mobileQuality",
		"mobileScale"}

	a.WordFields = fields
	a.IdType = 11
	a.Ids = []int64{37355424201}

	body, err := c.GetKeyword(a)
	if err != nil {
		t.Error(err)
	}
	var car KeywordGetResponse
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
	if len(car.Body.KeywordTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	fmt.Println("Keyword is :", car.Body.KeywordTypes[0])

	fmt.Println("############################################")

}
