package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnDisableRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsWarnDisableRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnDisableRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnDisableRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnDisableRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnDisableRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnDisableRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnDisableRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnDisableRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnDisableRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnDisableRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsWarnDisableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnDisable")
	r.SetProduct(Product)
}

type OnsWarnDisableResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnDisable(req *OnsWarnDisableRequest, accessId, accessSecret string) (*OnsWarnDisableResponse, error) {
	var pResponse OnsWarnDisableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
