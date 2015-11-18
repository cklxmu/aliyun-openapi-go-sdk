package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsPublishDeleteRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ProducerId   string
	Topic        string
}

func (r *OnsPublishDeleteRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsPublishDeleteRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsPublishDeleteRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsPublishDeleteRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsPublishDeleteRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsPublishDeleteRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsPublishDeleteRequest) SetProducerId(value string) {
	r.ProducerId = value
	r.QueryParams.Set("ProducerId", value)
}
func (r *OnsPublishDeleteRequest) GetProducerId() string {
	return r.ProducerId
}
func (r *OnsPublishDeleteRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsPublishDeleteRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsPublishDeleteRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsPublishDelete")
	r.SetProduct(Product)
}

type OnsPublishDeleteResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsPublishDelete(req *OnsPublishDeleteRequest, accessId, accessSecret string) (*OnsPublishDeleteResponse, error) {
	var pResponse OnsPublishDeleteResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
