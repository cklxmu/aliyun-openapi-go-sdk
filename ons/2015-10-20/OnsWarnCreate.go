package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnCreateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
	Threshold    string
	Contacts     string
}

func (r *OnsWarnCreateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnCreateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnCreateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnCreateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnCreateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnCreateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnCreateRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnCreateRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnCreateRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnCreateRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsWarnCreateRequest) SetThreshold(value string) {
	r.Threshold = value
	r.QueryParams.Set("Threshold", value)
}
func (r *OnsWarnCreateRequest) GetThreshold() string {
	return r.Threshold
}
func (r *OnsWarnCreateRequest) SetContacts(value string) {
	r.Contacts = value
	r.QueryParams.Set("Contacts", value)
}
func (r *OnsWarnCreateRequest) GetContacts() string {
	return r.Contacts
}

func (r *OnsWarnCreateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnCreate")
	r.SetProduct(Product)
}

type OnsWarnCreateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnCreate(req *OnsWarnCreateRequest, accessId, accessSecret string) (*OnsWarnCreateResponse, error) {
	var pResponse OnsWarnCreateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
