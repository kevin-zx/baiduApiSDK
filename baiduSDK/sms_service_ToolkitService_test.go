// sms_service_ToolkitService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetOperationRecord(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetOperationRecord start ...")
	fmt.Println("############################################")
	c := NewToolkitService()
	c.AuthHeader = authHeader
	a := NewOperationRecordGetRequest()
	a.StartDate = "2016-02-28"
	a.EndDate = "2016-03-23"
	a.OptLevel = 10
	a.OptContents = OptLevelFields[10]
	body, err := c.GetOperationRecord(a)
	if err != nil {
		t.Error(err)
	}

	var car OperationRecordGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}

	for _, v := range car.Body.OptLogTypes {
		fmt.Println("OperationRecord  is :", v)
	}
	fmt.Println("############################################")

}
