package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTrendTopicInputTpsRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	BeginTime    int
	EndTime      int
	Period       int
}

func (r *OnsTrendTopicInputTpsRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTrendTopicInputTpsRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTrendTopicInputTpsRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTrendTopicInputTpsRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTrendTopicInputTpsRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTrendTopicInputTpsRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTrendTopicInputTpsRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTrendTopicInputTpsRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsTrendTopicInputTpsRequest) SetBeginTime(value int) {
	r.BeginTime = value
	r.QueryParams.Set("BeginTime", strconv.Itoa(value))
}
func (r *OnsTrendTopicInputTpsRequest) GetBeginTime() int {
	return r.BeginTime
}
func (r *OnsTrendTopicInputTpsRequest) SetEndTime(value int) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", strconv.Itoa(value))
}
func (r *OnsTrendTopicInputTpsRequest) GetEndTime() int {
	return r.EndTime
}
func (r *OnsTrendTopicInputTpsRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *OnsTrendTopicInputTpsRequest) GetPeriod() int {
	return r.Period
}

func (r *OnsTrendTopicInputTpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTrendTopicInputTps")
	r.SetProduct(Product)
}

type OnsTrendTopicInputTpsResponse struct {
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

func OnsTrendTopicInputTps(req *OnsTrendTopicInputTpsRequest, accessId, accessSecret string) (*OnsTrendTopicInputTpsResponse, error) {
	var pResponse OnsTrendTopicInputTpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
