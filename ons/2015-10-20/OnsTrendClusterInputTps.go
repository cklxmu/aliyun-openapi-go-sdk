package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTrendClusterInputTpsRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Cluster      string
	BeginTime    int
	EndTime      int
	Period       int
}

func (r *OnsTrendClusterInputTpsRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTrendClusterInputTpsRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTrendClusterInputTpsRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTrendClusterInputTpsRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTrendClusterInputTpsRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTrendClusterInputTpsRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTrendClusterInputTpsRequest) SetCluster(value string) {
	r.Cluster = value
	r.QueryParams.Set("Cluster", value)
}
func (r *OnsTrendClusterInputTpsRequest) GetCluster() string {
	return r.Cluster
}
func (r *OnsTrendClusterInputTpsRequest) SetBeginTime(value int) {
	r.BeginTime = value
	r.QueryParams.Set("BeginTime", strconv.Itoa(value))
}
func (r *OnsTrendClusterInputTpsRequest) GetBeginTime() int {
	return r.BeginTime
}
func (r *OnsTrendClusterInputTpsRequest) SetEndTime(value int) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", strconv.Itoa(value))
}
func (r *OnsTrendClusterInputTpsRequest) GetEndTime() int {
	return r.EndTime
}
func (r *OnsTrendClusterInputTpsRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *OnsTrendClusterInputTpsRequest) GetPeriod() int {
	return r.Period
}

func (r *OnsTrendClusterInputTpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTrendClusterInputTps")
	r.SetProduct(Product)
}

type OnsTrendClusterInputTpsResponse struct {
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

func OnsTrendClusterInputTps(req *OnsTrendClusterInputTpsRequest, accessId, accessSecret string) (*OnsTrendClusterInputTpsResponse, error) {
	var pResponse OnsTrendClusterInputTpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
