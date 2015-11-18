package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type WafLogRequest struct {
	RpcRequest
	JstOwnerId   int
	PageNumber   int
	PageSize     int
	InstanceId   string
	InstanceType string
}

func (r *WafLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *WafLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *WafLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *WafLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *WafLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *WafLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *WafLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *WafLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *WafLogRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *WafLogRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *WafLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("WafLog")
	r.SetProduct(Product)
}

type WafLogResponse struct {
	WebAttack   int    `xml:"WebAttack" json:"WebAttack"`
	NewWafUser  bool   `xml:"NewWafUser" json:"NewWafUser"`
	WafOpened   bool   `xml:"WafOpened" json:"WafOpened"`
	InWhiteList bool   `xml:"InWhiteList" json:"InWhiteList"`
	DomainCount int    `xml:"DomainCount" json:"DomainCount"`
	StartTime   string `xml:"StartTime" json:"StartTime"`
	EndTime     string `xml:"EndTime" json:"EndTime"`
	PageNumber  int    `xml:"PageNumber" json:"PageNumber"`
	PageSize    int    `xml:"PageSize" json:"PageSize"`
	TotalCount  int    `xml:"TotalCount" json:"TotalCount"`
	LogList     struct {
		WafLog []struct {
			SourceIp string `xml:"SourceIp" json:"SourceIp"`
			Time     string `xml:"Time" json:"Time"`
			Url      string `xml:"Url" json:"Url"`
			Type     string `xml:"Type" json:"Type"`
			Status   int    `xml:"Status" json:"Status"`
		} `xml:"WafLog" json:"WafLog"`
	} `xml:"LogList" json:"LogList"`
}

func WafLog(req *WafLogRequest, accessId, accessSecret string) (*WafLogResponse, error) {
	var pResponse WafLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
