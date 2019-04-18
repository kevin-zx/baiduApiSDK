// sms_service_AdgroupService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllAdgroup(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllAdgroup start ...")
	fmt.Println("############################################")
	c := NewAdgroupService()
	c.AuthHeader = authHeader
	a := NewAdgroupGetRequest()
	fields := []string{
		"adgroupId",
		"campaignId",
		"adgroupName",
		"maxPrice",
		"negativeWords",
		"exactNegativeWords",
		"pause",
		"status",
		"priceRatio",
		"accuPriceFactor",
		"wordPriceFactor",
		"widePriceFactor",
		"matchPriceStatus",
	}
	a.AdgroupFields = fields
	a.IdType = 3
	a.Ids = []int64{28213748}

	body, err := c.GetAdgroup(a)
	if err != nil {
		t.Error(err)
	}
	var car AdgroupGetResponse
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
	if len(car.Body.AdgroupTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	for _, v := range car.Body.AdgroupTypes {
		fmt.Println("Adgroup is :", v)
	}

	fmt.Println("############################################")

}

func TestGetAdgroup(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAdgroup start ...")
	fmt.Println("############################################")
	c := NewAdgroupService()
	c.AuthHeader = authHeader
	a := NewAdgroupGetRequest()
	fields := []string{
		"adgroupId",
		"campaignId",
		"adgroupName",
		"maxPrice",
		"negativeWords",
		"exactNegativeWords",
		"pause",
		"status",
		"priceRatio",
		"accuPriceFactor",
		"wordPriceFactor",
		"widePriceFactor",
		"matchPriceStatus",
	}
	a.AdgroupFields = fields
	a.IdType = 5
	a.Ids = []int64{1720779401}

	body, err := c.GetAdgroup(a)
	if err != nil {
		t.Error(err)
	}
	var car AdgroupGetResponse
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
	fmt.Println("Adgroup is :", car.Body.AdgroupTypes[0])

	fmt.Println("############################################")

}
