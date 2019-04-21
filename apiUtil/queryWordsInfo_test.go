package apiUtil

import (
	"fmt"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
	"testing"
)

var krAuthHeader = &baiduSDK.AuthHeader{

	//Username: "",
	//Password: "",
	//Token:    "",
	//Action: "API-SDK",

}

func TestQueryTasks_Query(t *testing.T) {
	qs := NewQueryService(krAuthHeader)
	kis, err := qs.Query([]string{"斜塘二手房"})
	if err != nil {
		panic(err)
	}
	for _, ki := range *kis {
		fmt.Printf("%v\n", ki)
	}
}
