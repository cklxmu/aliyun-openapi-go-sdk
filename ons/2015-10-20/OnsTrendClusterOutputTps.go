package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTrendClusterOutputTpsRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Cluster      string
	BeginTime    int
	EndTime      int
	Period       int
}

func (r *OnsTrendClusterOutputTpsRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTrendClusterOutputTpsRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTrendClusterOutputTpsRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTrendClusterOutputTpsRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTrendClusterOutputTpsRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTrendClusterOutputTpsRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTrendClusterOutputTpsRequest) SetCluster(value string) {
	r.Cluster = value
	r.QueryParams.Set("Cluster", value)
}
func (r *OnsTrendClusterOutputTpsRequest) GetCluster() string {
	return r.Cluster
}
func (r *OnsTrendClusterOutputTpsRequest) SetBeginTime(value int) {
	r.BeginTime = value
	r.QueryParams.Set("BeginTime", strconv.Itoa(value))
}
func (r *OnsTrendClusterOutputTpsRequest) GetBeginTime() int {
	return r.BeginTime
}
func (r *OnsTrendClusterOutputTpsRequest) SetEndTime(value int) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", strconv.Itoa(value))
}
func (r *OnsTrendClusterOutputTpsRequest) GetEndTime() int {
	return r.EndTime
}
func (r *OnsTrendClusterOutputTpsRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *OnsTrendClusterOutputTpsRequest) GetPeriod() int {
	return r.Period
}

func (r *OnsTrendClusterOutputTpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTrendClusterOutputTps")
	r.SetProduct(Product)
}

type OnsTrendClusterOutputTpsResponse struct {
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

func OnsTrendClusterOutputTps(req *OnsTrendClusterOutputTpsRequest, accessId, accessSecret string) (*OnsTrendClusterOutputTpsResponse, error) {
	var pResponse OnsTrendClusterOutputTpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
