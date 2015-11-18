package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTopicDeleteRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	Cluster      string
}

func (r *OnsTopicDeleteRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTopicDeleteRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTopicDeleteRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTopicDeleteRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTopicDeleteRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTopicDeleteRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTopicDeleteRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTopicDeleteRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsTopicDeleteRequest) SetCluster(value string) {
	r.Cluster = value
	r.QueryParams.Set("Cluster", value)
}
func (r *OnsTopicDeleteRequest) GetCluster() string {
	return r.Cluster
}

func (r *OnsTopicDeleteRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTopicDelete")
	r.SetProduct(Product)
}

type OnsTopicDeleteResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsTopicDelete(req *OnsTopicDeleteRequest, accessId, accessSecret string) (*OnsTopicDeleteResponse, error) {
	var pResponse OnsTopicDeleteResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
