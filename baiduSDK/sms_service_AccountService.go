// sms_service_AccountService.go
package baiduSDK

import ()

const (
	//获取帐户信息
	GET_AccountInfo = "getAccountInfo"
	//更新账户信息
	UPDATE_AccountInfo = "updateAccountInfo"
)

// 账户服务AccountService：

// 获取账户ID，投资、消费、余额等财务数据、推广域名和开放域名； 获取和设置账户日预算和周预算、推广地域、IP排除等
type AccountService struct {
	*CommonService
}

func NewAccountService() *AccountService {
	a := new(AccountService)
	a.CommonService = NewCommonService("sms", "service", "AccountService")
	return a
}

func (a *AccountService) Execute(method string, request interface{}) ([]byte, error) {
	return a.do(method, request)

}

func (a *AccountService) GetAccountInfo(request interface{}) ([]byte, error) {
	return a.Execute(GET_AccountInfo, request)
}

func (a *AccountService) UpdateAccountInfo(request interface{}) ([]byte, error) {
	return a.Execute(UPDATE_AccountInfo, request)
}

//accountInfoType
type AccountInfo struct {
	Userid               int64              `json:"userid"`
	Balance              float64            `json:"balance"`
	PcBalance            float64            `json:"pcBalance"`
	MobileBalance        float64            `json:"MobileBalance"`
	Cost                 float64            `json:"Cost"`
	Payment              float64            `json:"Payment"`
	BudgetType           int                `json:"budgetType"`
	Budget               float64            `json:"budget"`
	RegionTarget         []int              `json:"regionTarget"`
	ExcludeIp            []string           `json:"excludeIp"`
	OpenDomains          []string           `json:"openDomains"`
	RegDomain            string             `json:"regDomain"`
	BudgetOfflineTime    []*OfflineTimeType `json:"budgetOfflineTime"`
	WeeklyBudget         []float64          `json:"weeklyBudget"`
	UserStat             int                `json:"userStat"`
	IsDynamicCreative    bool               `json:"isDynamicCreative"`
	IsDynamicTagSublink  bool               `json:"isDynamicTagSublink"`
	IsDynamicTitle       bool               `json:"isDynamicTitle"`
	IsDynamicHotRedirect bool               `json:"isDynamicHotRedirect"`
	DynamicCreativeParam string             `json:"dynamicCreativeParam"`
}

type OfflineTimeType struct {
	Flag int    `json:"flag"`
	Time string `json:"time"`
}

//定义更新用户信息结果
type UpdateAccountInfo struct {
	Header *ResponseHeader               `json:"header"`
	Body   UpdateAccountInfoResponseBody `json:"body"`
}
type UpdateAccountInfoResponseBody struct {
	AccountInfo *AccountInfo `json:"accountInfo"`
}

func NewUpdateAccountInfo() *UpdateAccountInfo {
	a := new(UpdateAccountInfo)
	return a
}

//定义获取用户返回信息结构
type AccountInfoResponse struct {
	Header *ResponseHeader         `json:"header"`
	Body   AccountInfoResponseBody `json:"body"`
}
type AccountInfoResponseBody struct {
	Data []*AccountInfo `json:"data"`
}

//定义用户请求信息字段
type AccountInfoRequest struct {
	AccountFields []string `json:"accountFields"`
}

func NewAccountInfoRequest() *AccountInfoRequest {
	a := new(AccountInfoRequest)
	return a
}
func (a *AccountInfoRequest) SetAccountInfoFiled(fields []string) {
	a.AccountFields = fields
}

//检查填写 field 是否合理
func (a *AccountInfoRequest) CheckFiled(f string) bool {
	if _, ok := AccountInfoFileds[f]; ok {
		return true
	}
	return false
}

var AccountInfoFileds = map[string]string{
	"userId":               "账户 ID",
	"balance":              "推广共享包金额",
	"pcBalance":            "基准资金包金额",
	"mobileBalance":        "无线优惠资金包金额",
	"cost":                 "账户累积消费",
	"payment":              "账户投资",
	"budgetType":           "账户预算类型",
	"budget":               "账户预算",
	"regionTarget":         "推广地域列表",
	"excludeIp":            "ip 排除列表",
	"openDomains":          "账户开放域名列表",
	"regDomain":            "账户注册域名",
	"budgetOfflineTime":    "到达预算下线时段",
	"weeklyBudget":         "返回本周的每日预算值",
	"userStat":             "账户状态",
	"isDynamicCreative":    "子链开关",
	"isDynamicTagSublink":  "标签子链开关",
	"isDynamicTitle":       "动态标题开关",
	"isDynamicHotRedirect": "热点直达开关",
	"dynamicCreativeParam": "动态创意统计参数"}

//获取账户预算类型
func (a *AccountInfo) GetBudgetType() string {
	if a.BudgetType == 0 {
		return "不设置预算"
	} else if a.BudgetType == 1 {
		return "日预算"
	} else if a.BudgetType == 1 {
		return "周预算"
	}
	return "非法值"

}

//根据账户类型 验证账户预算范围是否正常
func (a *AccountInfo) CheckBudget() (bool, string) {
	if a.BudgetType == 0 {
		return true, ""
	} else if a.BudgetType == 1 {
		if a.Budget >= 50 && a.Budget <= 10000000 {
			return true, ""
		}
		return false, "日预算取值范围:[50,10000000]"
	} else if a.BudgetType == 1 {
		if a.Budget >= 388 && a.Budget <= 70000000 {
			return true, ""
		}
		return false, "周预算取值范围:[388,70000000]"
	}
	return false, "非法值"

}

//获取推广地域列表
func (a *AccountInfo) GetRegionTarget() (s []string) {
	for _, v := range a.RegionTarget {
		s = append(s, regionMap[v])
	}
	return s

}

//检查IP排除列表 数组元素个数最大值:203
func (a *AccountInfo) CheckExcludeIp() bool {
	if len(a.ExcludeIp) > 203 {
		return false
	}
	return true
}

var userStatMap = map[int]string{
	1:  "开户金未到",
	2:  "正常生效",
	3:  "余额为零",
	4:  "未通过审核",
	6:  "审核中",
	7:  "被禁用",
	11: "预算不足",
}

//获取账户状态(用于一站式平台显示账户状态)
func (a *AccountInfo) GetUserStat() string {
	return userStatMap[a.UserStat]
}

//检查动态创意统计参数  以”?”或者”#”开头,长度不超 过 256 字节
func (a *AccountInfo) CheckDynamicCreativeParam() bool {
	if len([]byte(a.DynamicCreativeParam)) > 256 {
		return false
	}
	return true

}
