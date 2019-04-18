// sms_service_DynamicCreativeService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetDynCreative(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("动态创意－－TestGetDynCreative start ...")
	fmt.Println("############################################")
	c := NewDynamicCreativeService()
	c.AuthHeader = authHeader
	a := NewDynCreativeGetRequest()
	//单元
	// a.Ids = []int64{1715190897}
	// a.IdType = 5 //13 3
	//计划
	a.Ids = []int64{28213748}
	a.IdType = 3
	a.DynCreativeFields = []string{
		"dynCreativeId",
		"campaignId",
		"adgroupId",
		"bindingType",
		"title",
		"url",
		"murl",
		"status",
		"pause"}

	body, err := c.GetDynCreative(a)
	if err != nil {
		t.Error(err)
	}

	var car DynCreativeGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}

	for _, v := range car.Body.DynCreativeTypes {
		fmt.Println("DynCreativeType is :", v)
	}
	fmt.Println("############################################")

}
func TestGetExclusionTypeByCampaignId(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("动态创意－－GetExclusionTypeByCampaignId start ...")
	fmt.Println("############################################")
	c := NewDynamicCreativeService()
	c.AuthHeader = authHeader
	a := NewExclusionTypeByCampaignIdGetRequest()

	a.CampaignIds = []int64{28213748}

	body, err := c.GetExclusionTypeByCampaignId(a)
	if err != nil {
		t.Error(err)
	}

	var car ExclusionTypeByCampaignIdGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}

	for _, v := range car.Body.DynCreativeExclusionTypes {
		fmt.Println("DynCreativeExclusionType is :", v)
	}
	fmt.Println("############################################")

}
