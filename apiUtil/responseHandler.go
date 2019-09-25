package apiUtil

import (
	"errors"
	"github.com/tidwall/gjson"
)

//如果response 说明是成功的，则返回data ，如果不成功抛出异常
func responseIsSuccess(response []byte) (data string, err error) {
	//fmt.Println(string(response))
	jo := gjson.ParseBytes(response)
	successFlag := jo.Get("header.desc").String()
	if successFlag != "success" {
		return "", errors.New(string(response))
	}
	data = jo.Get("body.data").Raw
	if data == "" {
		return "", errors.New("response is empty")
	}
	return
}
