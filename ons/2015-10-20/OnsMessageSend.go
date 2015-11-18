package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessageSendRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ProducerId   string
	Topic        string
	Tag          string
	Key          string
	Message      string
}

func (r *OnsMessageSendRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessageSendRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessageSendRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessageSendRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessageSendRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessageSendRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessageSendRequest) SetProducerId(value string) {
	r.ProducerId = value
	r.QueryParams.Set("ProducerId", value)
}
func (r *OnsMessageSendRequest) GetProducerId() string {
	return r.ProducerId
}
func (r *OnsMessageSendRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsMessageSendRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsMessageSendRequest) SetTag(value string) {
	r.Tag = value
	r.QueryParams.Set("Tag", value)
}
func (r *OnsMessageSendRequest) GetTag() string {
	return r.Tag
}
func (r *OnsMessageSendRequest) SetKey(value string) {
	r.Key = value
	r.QueryParams.Set("Key", value)
}
func (r *OnsMessageSendRequest) GetKey() string {
	return r.Key
}
func (r *OnsMessageSendRequest) SetMessage(value string) {
	r.Message = value
	r.QueryParams.Set("Message", value)
}
func (r *OnsMessageSendRequest) GetMessage() string {
	return r.Message
}

func (r *OnsMessageSendRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessageSend")
	r.SetProduct(Product)
}

type OnsMessageSendResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    string `xml:"Data" json:"Data"`
}

func OnsMessageSend(req *OnsMessageSendRequest, accessId, accessSecret string) (*OnsMessageSendResponse, error) {
	var pResponse OnsMessageSendResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
