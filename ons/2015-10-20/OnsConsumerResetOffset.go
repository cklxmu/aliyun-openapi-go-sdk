package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsConsumerResetOffsetRequest struct {
	RpcRequest
	OnsRegionId    string
	OnsPlatform    string
	PreventCache   int
	ConsumerId     string
	Topic          string
	ResetTimestamp int
}

func (r *OnsConsumerResetOffsetRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsConsumerResetOffsetRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsConsumerResetOffsetRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsConsumerResetOffsetRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsConsumerResetOffsetRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsConsumerResetOffsetRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsConsumerResetOffsetRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsConsumerResetOffsetRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsConsumerResetOffsetRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsConsumerResetOffsetRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsConsumerResetOffsetRequest) SetResetTimestamp(value int) {
	r.ResetTimestamp = value
	r.QueryParams.Set("ResetTimestamp", strconv.Itoa(value))
}
func (r *OnsConsumerResetOffsetRequest) GetResetTimestamp() int {
	return r.ResetTimestamp
}

func (r *OnsConsumerResetOffsetRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsConsumerResetOffset")
	r.SetProduct(Product)
}

type OnsConsumerResetOffsetResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsConsumerResetOffset(req *OnsConsumerResetOffsetRequest, accessId, accessSecret string) (*OnsConsumerResetOffsetResponse, error) {
	var pResponse OnsConsumerResetOffsetResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
