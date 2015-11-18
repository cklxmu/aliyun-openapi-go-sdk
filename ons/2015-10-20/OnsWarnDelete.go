package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnDeleteRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsWarnDeleteRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnDeleteRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnDeleteRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnDeleteRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnDeleteRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnDeleteRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnDeleteRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnDeleteRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnDeleteRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnDeleteRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsWarnDeleteRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnDelete")
	r.SetProduct(Product)
}

type OnsWarnDeleteResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnDelete(req *OnsWarnDeleteRequest, accessId, accessSecret string) (*OnsWarnDeleteResponse, error) {
	var pResponse OnsWarnDeleteResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
