package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnEditorRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
	Threshold    string
	Contacts     string
}

func (r *OnsWarnEditorRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnEditorRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnEditorRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnEditorRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnEditorRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnEditorRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnEditorRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnEditorRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnEditorRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnEditorRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsWarnEditorRequest) SetThreshold(value string) {
	r.Threshold = value
	r.QueryParams.Set("Threshold", value)
}
func (r *OnsWarnEditorRequest) GetThreshold() string {
	return r.Threshold
}
func (r *OnsWarnEditorRequest) SetContacts(value string) {
	r.Contacts = value
	r.QueryParams.Set("Contacts", value)
}
func (r *OnsWarnEditorRequest) GetContacts() string {
	return r.Contacts
}

func (r *OnsWarnEditorRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnEditor")
	r.SetProduct(Product)
}

type OnsWarnEditorResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsWarnEditor(req *OnsWarnEditorRequest, accessId, accessSecret string) (*OnsWarnEditorResponse, error) {
	var pResponse OnsWarnEditorResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
