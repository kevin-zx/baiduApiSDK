// sms_service_CampaignService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

//如果 CampaignIds 不设置值，返回所有数据
func TestGetAllCampain(t *testing.T) {

	c := NewCampaignService()
	c.AuthHeader = authHeader
	a := NewCampaignGetRequest()
	fields := []string{
		"campaignId",
		"campaignName",
		"budget",
		"budgetOfflineTime",
		"campaignType",
		"device",
		"exactNegativeWords",
		"isDynamicCreative",
		"isDynamicTagSublink",
		"isDynamicTitle",
		"isDynamicHotRedirect",
		"regionTarget",
		"negativeWords",
		"pause",
		"rmktStatus",
		"priceRatio",
		"schedule",
		"showProb",
		"status",
		"rmktPriceRatio",
	}
	a.CampaignFields = fields
	fmt.Println("############################################")
	fmt.Println("TestGetAllCampain start ...")
	fmt.Println("############################################")
	body, err := c.GetCampain(a)
	if err != nil {
		t.Error(err)
	}
	var car CampaignGetResponse
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
	if len(car.Body.CampaignTypes) == 0 {
		t.Error("account response body is null")
		return
	}
	for _, v := range car.Body.CampaignTypes {
		fmt.Println("Campain is :", v)
	}

	fmt.Println("############################################")

}
func TestGetCampain(t *testing.T) {
	fmt.Println("############################################")
	fmt.Println("TestGetCampain start ...")
	fmt.Println("############################################")
	//试获取单个计划信息通过ID
	c := NewCampaignService()
	c.AuthHeader = authHeader
	a := NewCampaignGetRequest()
	fields := []string{
		"campaignId",
		"campaignName",
		"budget",
		"budgetOfflineTime",
		"campaignType",
		"device",
		"exactNegativeWords",
		"isDynamicCreative",
		"isDynamicTagSublink",
		"isDynamicTitle",
		"isDynamicHotRedirect",
		"regionTarget",
		"negativeWords",
		"pause",
		"rmktStatus",
		"priceRatio",
		"schedule",
		"showProb",
		"status",
		"rmktPriceRatio",
	}
	a.CampaignFields = fields
	a.CampaignIds = []int64{}

	body, err := c.GetCampain(a)
	if err != nil {
		t.Error(err)
	}
	var car CampaignGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}
	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	fmt.Println("Campain is :", car.Body.CampaignTypes[0])
	fmt.Println("############################################")

}

//测
