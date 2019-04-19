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
	qts := QueryTasks{}
	kis, err := qts.Query([]string{"测试"}, krAuthHeader)
	if err != nil {
		panic(err)
	}
	for _, ki := range *kis {
		fmt.Printf("%v\n", ki)
	}
}
