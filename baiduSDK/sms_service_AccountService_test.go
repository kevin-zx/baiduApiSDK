// sms_service_AccountService_Test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

var authHeader = &AuthHeader{
	Username: "",
	Password: "",
	Token:    "",
}

func TestGetAccountInfo(t *testing.T) {

	s := NewAccountService()
	s.AuthHeader = authHeader

	// 构建获取的FIlelds

	fields := []string{
		"userId", "balance", "cost", "payment", "budgetType", "budget", "regionTarget", "excludeIp",
		"openDomains", "regDomain", "budgetOfflineTime", "weeklyBudget", "userStat",
		"isDynamicCreative", "dynamicCreativeParam", "pcBalance",
		"mobileBalance", "isDynamicTagSublink", "isDynamicTitle", "isDynamicHotRedirect"}

	a := NewAccountInfoRequest()
	//去除不符合规则 的 fields
	var validFileds = []string{}
	for _, f := range fields {
		if !a.CheckFiled(f) {
			t.Errorf("%s is not valid Field", f)
			continue
		}
		validFileds = append(validFileds, f)
	}

	fmt.Println("############################################")
	fmt.Println("TestGetAccountInfo start ...")
	fmt.Println("############################################")
	fmt.Println("validFileds is:", validFileds)
	fmt.Println("############################################")

	a.SetAccountInfoFiled(validFileds)
	body, err := s.GetAccountInfo(a)
	if err != nil {
		t.Error(err)
	}

	var air AccountInfoResponse
	err = json.Unmarshal(body, &air)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}
	//获取返回数据状态
	if ok, message := air.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	//获取返回数据状态
	if len(air.Body.Data) == 0 {
		t.Error("account response body is null")
		return
	}

	fmt.Println("GetAccountInfo:", air.Body.Data[0])
	fmt.Println("GetAccountInfoRegion:", air.Body.Data[0].GetRegionTarget())
	fmt.Println("GetAccountInfoUserStat:", air.Body.Data[0].GetUserStat())
	fmt.Println("############################################")
}

// func TestUpdateAccountInfo(t *testing.T) {

// 	s := NewAccountService()
// 	s.AuthHeader = authHeader

// 	// 构建获取的FIlelds

// 	a := NewUpdateAccountInfo()

// 	//修改用户信息

// 	fmt.Println("############################################")
// 	fmt.Println("TestUpdateAccountInfo start ...")
// 	fmt.Println("############################################")
// 	body, err := s.UpdateAccountInfo(a)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	var air UpdateAccountInfo
// 	err = json.Unmarshal(body, &air)
// 	if err != nil {
// 		fmt.Println("json Unmarshal err:", err)
// 	}
// 	//获取返回数据状态
// 	if ok, message := air.AuthHeader.GetResponseStatus(); !ok {
// 		fmt.Println("failures Message:", message)
// 		return
// 	}

// 	fmt.Println("UpdateAccountInfo:", air.Body.AccountInfo)
// 	fmt.Println("############################################")
// }

// func Test"updateAccountInfo", request)
