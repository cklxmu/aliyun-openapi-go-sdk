package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTrendGroupOutputTpsRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
	BeginTime    int
	EndTime      int
	Period       int
}

func (r *OnsTrendGroupOutputTpsRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTrendGroupOutputTpsRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTrendGroupOutputTpsRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTrendGroupOutputTpsRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTrendGroupOutputTpsRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTrendGroupOutputTpsRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTrendGroupOutputTpsRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsTrendGroupOutputTpsRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsTrendGroupOutputTpsRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTrendGroupOutputTpsRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsTrendGroupOutputTpsRequest) SetBeginTime(value int) {
	r.BeginTime = value
	r.QueryParams.Set("BeginTime", strconv.Itoa(value))
}
func (r *OnsTrendGroupOutputTpsRequest) GetBeginTime() int {
	return r.BeginTime
}
func (r *OnsTrendGroupOutputTpsRequest) SetEndTime(value int) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", strconv.Itoa(value))
}
func (r *OnsTrendGroupOutputTpsRequest) GetEndTime() int {
	return r.EndTime
}
func (r *OnsTrendGroupOutputTpsRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *OnsTrendGroupOutputTpsRequest) GetPeriod() int {
	return r.Period
}

func (r *OnsTrendGroupOutputTpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTrendGroupOutputTps")
	r.SetProduct(Product)
}

type OnsTrendGroupOutputTpsResponse struct {
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

func OnsTrendGroupOutputTps(req *OnsTrendGroupOutputTpsRequest, accessId, accessSecret string) (*OnsTrendGroupOutputTpsResponse, error) {
	var pResponse OnsTrendGroupOutputTpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
