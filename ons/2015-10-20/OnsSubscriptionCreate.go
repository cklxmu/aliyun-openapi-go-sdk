package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsSubscriptionCreateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsSubscriptionCreateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsSubscriptionCreateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsSubscriptionCreateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsSubscriptionCreateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsSubscriptionCreateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsSubscriptionCreateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsSubscriptionCreateRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsSubscriptionCreateRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsSubscriptionCreateRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsSubscriptionCreateRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsSubscriptionCreateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsSubscriptionCreate")
	r.SetProduct(Product)
}

type OnsSubscriptionCreateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsSubscriptionCreate(req *OnsSubscriptionCreateRequest, accessId, accessSecret string) (*OnsSubscriptionCreateResponse, error) {
	var pResponse OnsSubscriptionCreateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
