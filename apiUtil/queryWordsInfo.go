package apiUtil

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
	"github.com/tidwall/gjson"
)

type SearchInfo struct {
	ShowRate float64 `json:"showRate"`
	Rank     int     `json:"rank"`
	Charge   float64 `json:"charge"`
	Cpc      float64 `json:"cpc"`
	RecBid   float64 `json:"recBid"`
	Click    float64 `json:"click"`
	Pv       int     `json:"pv"`
	Show     int     `json:"show"`
	Ctr      float64 `json:"ctr"`
}

type WordInfo struct {
	MatchType   int        `json:"matchType"`
	Word        string     `json:"word"`
	Bid         float64    `json:"bid"`
	Competition int        `json:"competition"`
	Pc          SearchInfo `json:"pc"`
	All         SearchInfo `json:"all"`
	Mobile      SearchInfo `json:"mobile"`
}

type QueryTasks struct {
	Words []*QueryWord `json:"words"`
}

type QueryWord struct {
	Bid       float64 `json:"bid"`
	MatchType int     `json:"matchType"`
	Word      string  `json:"word"`
}

func (qts *QueryTasks) Query(words []string, authHeader *baiduSDK.AuthHeader) (sis *[]WordInfo, err error) {
	// 这里其实超过500会出错，但是这边先限制在100，试用几个版本后更改
	if len(words) > 100 {
		return nil, errors.New("keywords len too long, default limit:100")
	}
	s := baiduSDK.NewKRService()
	s.AuthHeader = authHeader
	qts.Words = convertWordsToQueryTasks(words)
	data, err := s.GetEstimatedDataByBid(qts)
	if err != nil {
		return
	}
	jo := gjson.ParseBytes(data)
	sucessFlag := jo.Get("header.succ").Int()
	if sucessFlag != 1 {
		return nil, errors.New(string(data))
	}
	d := jo.Get("body.data").Raw
	fmt.Println(d)
	if d == "" {
		return nil, errors.New("data is empty")
	}
	sis = &[]WordInfo{}
	err = json.Unmarshal([]byte(d), &sis)
	return
}

func convertWordsToQueryTasks(words []string) (qws []*QueryWord) {
	for _, w := range words {
		qws = append(qws, &QueryWord{Word: w, MatchType: 1, Bid: 1.0})
	}
	return
}
