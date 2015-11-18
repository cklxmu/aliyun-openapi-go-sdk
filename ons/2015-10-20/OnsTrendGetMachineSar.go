package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTrendGetMachineSarRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	HostIp       string
	AppId        string
	Metric       string
	BeginTime    int
	EndTime      int
	Period       int
}

func (r *OnsTrendGetMachineSarRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTrendGetMachineSarRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTrendGetMachineSarRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTrendGetMachineSarRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTrendGetMachineSarRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTrendGetMachineSarRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTrendGetMachineSarRequest) SetHostIp(value string) {
	r.HostIp = value
	r.QueryParams.Set("HostIp", value)
}
func (r *OnsTrendGetMachineSarRequest) GetHostIp() string {
	return r.HostIp
}
func (r *OnsTrendGetMachineSarRequest) SetAppId(value string) {
	r.AppId = value
	r.QueryParams.Set("AppId", value)
}
func (r *OnsTrendGetMachineSarRequest) GetAppId() string {
	return r.AppId
}
func (r *OnsTrendGetMachineSarRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *OnsTrendGetMachineSarRequest) GetMetric() string {
	return r.Metric
}
func (r *OnsTrendGetMachineSarRequest) SetBeginTime(value int) {
	r.BeginTime = value
	r.QueryParams.Set("BeginTime", strconv.Itoa(value))
}
func (r *OnsTrendGetMachineSarRequest) GetBeginTime() int {
	return r.BeginTime
}
func (r *OnsTrendGetMachineSarRequest) SetEndTime(value int) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", strconv.Itoa(value))
}
func (r *OnsTrendGetMachineSarRequest) GetEndTime() int {
	return r.EndTime
}
func (r *OnsTrendGetMachineSarRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *OnsTrendGetMachineSarRequest) GetPeriod() int {
	return r.Period
}

func (r *OnsTrendGetMachineSarRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTrendGetMachineSar")
	r.SetProduct(Product)
}

type OnsTrendGetMachineSarResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		Title   string `xml:"Title" json:"Title"`
		XUnit   string `xml:"XUnit" json:"XUnit"`
		YUnit   string `xml:"YUnit" json:"YUnit"`
		Records struct {
			StatsDataDo []struct {
				X int     `xml:"X" json:"X"`
				Y float32 `xml:"Y" json:"Y"`
			} `xml:"StatsDataDo" json:"StatsDataDo"`
		} `xml:"Records" json:"Records"`
	} `xml:"Data" json:"Data"`
}

func OnsTrendGetMachineSar(req *OnsTrendGetMachineSarRequest, accessId, accessSecret string) (*OnsTrendGetMachineSarResponse, error) {
	var pResponse OnsTrendGetMachineSarResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
