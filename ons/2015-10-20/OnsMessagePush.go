package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessagePushRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	ClientId     string
	MsgId        string
	Topic        string
}

func (r *OnsMessagePushRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessagePushRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessagePushRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessagePushRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessagePushRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessagePushRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessagePushRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsMessagePushRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsMessagePushRequest) SetClientId(value string) {
	r.ClientId = value
	r.QueryParams.Set("ClientId", value)
}
func (r *OnsMessagePushRequest) GetClientId() string {
	return r.ClientId
}
func (r *OnsMessagePushRequest) SetMsgId(value string) {
	r.MsgId = value
	r.QueryParams.Set("MsgId", value)
}
func (r *OnsMessagePushRequest) GetMsgId() string {
	return r.MsgId
}
func (r *OnsMessagePushRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsMessagePushRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsMessagePushRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessagePush")
	r.SetProduct(Product)
}

type OnsMessagePushResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsMessagePush(req *OnsMessagePushRequest, accessId, accessSecret string) (*OnsMessagePushResponse, error) {
	var pResponse OnsMessagePushResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
