// sms_service_BulkService_test.go
package baiduSDK

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetAllObjects(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllObjects start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewAllObjectGetRequest()
	// a.CampaignIds = []int64{28213748}
	a.AccountFields = []string{"all"}
	a.CreativeFields = []string{"all"}
	a.CampaignFields = []string{"all"}
	body, err := c.GetAllObjects(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car AllObjectsGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	if len(car.Body.AllObjectsFileIdGetResponse) == 0 {
		return
	}
	fmt.Println("getAllObjectsResponse fileId   is :", car.Body.AllObjectsFileIdGetResponse[0].FileId)
	fmt.Println("############################################")

}
func TestGetAllChangedObjects(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetAllChangedObjects start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewAllChangedObjectsGetRequest()
	a.StartTime = "2016-03-18"
	a.CreativeFields = []string{"all"}
	a.CampaignFields = []string{"all"}
	body, err := c.GetAllChangedObjects(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car AllChangedObjectsGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	if len(car.Body.AllChangedObjectsFileIdGetResponse) == 0 {
		return
	}
	fmt.Println("getAllChangedObjectsFileIdGetResponse fileId   is :", car.Body.AllChangedObjectsFileIdGetResponse[0].FileId)
	fmt.Println("############################################")

}

//35dc27dfa9dd876bbaeef6c27b8c030c
func TestGetFileStatus(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("查询请求的文件是否已生成--TestGetFileStatus start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewFileStateGetRequest()
	a.FileId = "35dc27dfa9dd876bbaeef6c27b8c030c"
	body, err := c.GetFileStatus(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car FileStateGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	if len(car.Body.FileStateIsGeneratedGetResponse) == 0 {
		return
	}
	fmt.Println("FileStateIsGeneratedGetResponse isGenerated   is :", car.Body.FileStateIsGeneratedGetResponse[0].IsGenerated)
	fmt.Println("############################################")
}

//35dc27dfa9dd876bbaeef6c27b8c030c
func TestGetFilePath(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("获取请求的文件下载地址--TestGetFileStatus start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewFilePathGetRequest()
	a.FileId = "35dc27dfa9dd876bbaeef6c27b8c030c"
	body, err := c.GetFilePath(a)
	if err != nil {
		t.Error(err)
	}
	var car FilePathGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	if len(car.Body.FilePathsGetResponse) == 0 {
		return
	}
	fmt.Println("FilePathGetResponse   is :", car.Body.FilePathsGetResponse[0])
	fmt.Println("############################################")
}

//35dc27dfa9dd876bbaeef6c27b8c030c
func TestCancelDownload(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("取消一个下载任务--TestCancelDownload start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewDownloadCancelRequest()
	a.FileId = "35dc27dfa9dd876bbaeef6c27b8c030c"
	body, err := c.CancelDownload(a)
	if err != nil {
		t.Error(err)
	}
	var car DownloadCancelResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	if len(car.Body.DownloadIsCanceledCancelResponse) == 0 {
		return
	}
	fmt.Println("CancelDownload isCanceled  is :", car.Body.DownloadIsCanceledCancelResponse[0])
	fmt.Println("############################################")
}

// getChangedId

func TestGetChangedId(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("获取账户内有变化的ID--TestGetChangedId start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewChangedIdGetRequest()
	a.StartTime = "2016-03-18"
	body, err := c.GetChangedId(a)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(string(body))
	var car ChangedIdGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	for _, v := range car.Body.ChangedIdType {
		for _, vv := range v.ChangedCreativeIds {
			fmt.Println("getChangedId ChangedCreativeId is :", vv)
		}

	}
	fmt.Println("############################################")
}

// getChangedItemId

func TestGetChangedItemId(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("获取从指定时间到当前时间段内有变化的物料 id--TestGetChangedItemId start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewChangedItemIdGetRequest()
	a.StartTime = "2016-03-18"
	a.ItemType = 1
	body, err := c.GetChangedItemId(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car ChangedItemIdGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}
	for _, v := range car.Body.ChangedItemIdDataTypes {
		for _, vv := range v.ChangedItemIds {
			fmt.Println("getChangedItemId changedItemId is :", vv)
		}

	}
	fmt.Println("############################################")
}

// getChangedScale

func TestGetChangedScale(t *testing.T) {

	fmt.Println("############################################")
	fmt.Println("TestGetChangedScale start ...")
	fmt.Println("############################################")
	c := NewBulkService()
	c.AuthHeader = authHeader
	a := NewChangedScaleGetRequest()
	a.StartTime = "2016-03-18"
	body, err := c.GetChangedScale(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(body))
	var car ChangedScaleGetResponse
	err = json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}

	//获取返回数据状态
	if ok, message := car.Header.GetResponseStatus(); !ok {
		fmt.Println("failures Message:", message)
		return
	}

	for _, v := range car.Body.ChangedScaleDataTypes {
		fmt.Println("getChangedScale is :", v)

	}
	fmt.Println("############################################")
}
