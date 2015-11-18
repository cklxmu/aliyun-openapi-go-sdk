package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsSubscriptionDeleteRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsSubscriptionDeleteRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsSubscriptionDeleteRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsSubscriptionDeleteRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsSubscriptionDeleteRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsSubscriptionDeleteRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsSubscriptionDeleteRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsSubscriptionDeleteRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsSubscriptionDeleteRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsSubscriptionDeleteRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsSubscriptionDeleteRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsSubscriptionDeleteRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsSubscriptionDelete")
	r.SetProduct(Product)
}

type OnsSubscriptionDeleteResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsSubscriptionDelete(req *OnsSubscriptionDeleteRequest, accessId, accessSecret string) (*OnsSubscriptionDeleteResponse, error) {
	var pResponse OnsSubscriptionDeleteResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
