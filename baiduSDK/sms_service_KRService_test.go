package baiduSDK

import (
	"fmt"
	"testing"
)

var krAuthHeader = &AuthHeader{
	Username: "",
	Password: "",
	Token:    "",
	Action:   "",
}

func TestKRService_GetKRFileIdByWords(t *testing.T) {
	s := NewKRService()
	s.AuthHeader = krAuthHeader
	s.SetVersion("v4")
	result, err := s.GetKRFileIdByWords(map[string]string{"seedWords": "seedWords", "seedFilter": ""})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
