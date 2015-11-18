package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsPublishCreateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ProducerId   string
	Topic        string
	AppName      string
}

func (r *OnsPublishCreateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsPublishCreateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsPublishCreateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsPublishCreateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsPublishCreateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsPublishCreateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsPublishCreateRequest) SetProducerId(value string) {
	r.ProducerId = value
	r.QueryParams.Set("ProducerId", value)
}
func (r *OnsPublishCreateRequest) GetProducerId() string {
	return r.ProducerId
}
func (r *OnsPublishCreateRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsPublishCreateRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsPublishCreateRequest) SetAppName(value string) {
	r.AppName = value
	r.QueryParams.Set("AppName", value)
}
func (r *OnsPublishCreateRequest) GetAppName() string {
	return r.AppName
}

func (r *OnsPublishCreateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsPublishCreate")
	r.SetProduct(Product)
}

type OnsPublishCreateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsPublishCreate(req *OnsPublishCreateRequest, accessId, accessSecret string) (*OnsPublishCreateResponse, error) {
	var pResponse OnsPublishCreateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
