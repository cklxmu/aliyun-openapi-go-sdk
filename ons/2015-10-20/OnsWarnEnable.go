package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnEnableRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsWarnEnableRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnEnableRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnEnableRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnEnableRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnEnableRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnEnableRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnEnableRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnEnableRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnEnableRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnEnableRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsWarnEnableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnEnable")
	r.SetProduct(Product)
}

type OnsWarnEnableResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnEnable(req *OnsWarnEnableRequest, accessId, accessSecret string) (*OnsWarnEnableResponse, error) {
	var pResponse OnsWarnEnableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
