package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsEmpowerDeleteRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	EmpowerUser  string
	Topic        string
}

func (r *OnsEmpowerDeleteRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsEmpowerDeleteRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsEmpowerDeleteRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsEmpowerDeleteRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsEmpowerDeleteRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsEmpowerDeleteRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsEmpowerDeleteRequest) SetEmpowerUser(value string) {
	r.EmpowerUser = value
	r.QueryParams.Set("EmpowerUser", value)
}
func (r *OnsEmpowerDeleteRequest) GetEmpowerUser() string {
	return r.EmpowerUser
}
func (r *OnsEmpowerDeleteRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsEmpowerDeleteRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsEmpowerDeleteRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsEmpowerDelete")
	r.SetProduct(Product)
}

type OnsEmpowerDeleteResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsEmpowerDelete(req *OnsEmpowerDeleteRequest, accessId, accessSecret string) (*OnsEmpowerDeleteResponse, error) {
	var pResponse OnsEmpowerDeleteResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
