package apiUtil

import (
	"encoding/json"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
)

type QueryExpandService struct {
	AuthHeader *baiduSDK.AuthHeader `json:"-"`
}
type QueryExpandTask struct {
	QueryType  int        `json:"queryType"`
	Query      string     `json:"query"`
	SeedFilter SeedFilter `json:"seedFilter"`
}

type SeedFilter struct {
	Device        int   `json:"device"`
	CompeteLow    int   `json:"competeLow"`
	CompeteHigh   int   `json:"competeHigh"`
	SearchRegions []int `json:"searchRegions"`
}

type QueryResult struct {
	PcPV        int      `json:"pcPV"`
	ShowReasons []string `json:"showReasons"`
	MobilePV    int      `json:"mobilePV"`
	RecBid      float64  `json:"recBid"`
	//BusinessPoints []interface{} `json:"businessPoints"`
	Pv          int     `json:"pv"`
	Word        string  `json:"word"`
	Similar     float64 `json:"similar"`
	Competition int     `json:"competition"`
}

// 0 是全部 1pc 2移动
func (qes *QueryExpandService) ExpandWordsByQuery(word string, device int) (qrs *[]QueryResult, err error) {
	krs := baiduSDK.NewKRService()
	krs.AuthHeader = qes.AuthHeader
	r, err := krs.GetKRByQuery(makeKeywordQueryExpandTask(word, device))
	if err != nil {
		return
	}
	d, err := responseIsSuccess(r)
	if err != nil {
		return
	}
	qrs = &[]QueryResult{}
	err = json.Unmarshal([]byte(d), qrs)
	return
}

func NewQueryExpandService(header *baiduSDK.AuthHeader) *QueryExpandService {
	return &QueryExpandService{AuthHeader: header}
}

func makeKeywordQueryExpandTask(word string, device int) (qet *QueryExpandTask) {
	qet = &QueryExpandTask{
		Query:     word,
		QueryType: 1,
		SeedFilter: SeedFilter{
			Device:      device,
			CompeteLow:  0,
			CompeteHigh: 100,
		},
	}
	return
}
