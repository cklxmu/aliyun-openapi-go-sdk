package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnAdminRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	UserId       string
	ConsumerId   string
	Topic        string
	Type         string
}

func (r *OnsWarnAdminRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnAdminRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnAdminRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnAdminRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnAdminRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnAdminRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnAdminRequest) SetUserId(value string) {
	r.UserId = value
	r.QueryParams.Set("UserId", value)
}
func (r *OnsWarnAdminRequest) GetUserId() string {
	return r.UserId
}
func (r *OnsWarnAdminRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnAdminRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnAdminRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnAdminRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsWarnAdminRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *OnsWarnAdminRequest) GetType() string {
	return r.Type
}

func (r *OnsWarnAdminRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnAdmin")
	r.SetProduct(Product)
}

type OnsWarnAdminResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnAdmin(req *OnsWarnAdminRequest, accessId, accessSecret string) (*OnsWarnAdminResponse, error) {
	var pResponse OnsWarnAdminResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
