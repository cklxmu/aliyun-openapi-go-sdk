package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsEmpowerCreateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	EmpowerUser  string
	Topic        string
	Relation     int
}

func (r *OnsEmpowerCreateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsEmpowerCreateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsEmpowerCreateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsEmpowerCreateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsEmpowerCreateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsEmpowerCreateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsEmpowerCreateRequest) SetEmpowerUser(value string) {
	r.EmpowerUser = value
	r.QueryParams.Set("EmpowerUser", value)
}
func (r *OnsEmpowerCreateRequest) GetEmpowerUser() string {
	return r.EmpowerUser
}
func (r *OnsEmpowerCreateRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsEmpowerCreateRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsEmpowerCreateRequest) SetRelation(value int) {
	r.Relation = value
	r.QueryParams.Set("Relation", strconv.Itoa(value))
}
func (r *OnsEmpowerCreateRequest) GetRelation() int {
	return r.Relation
}

func (r *OnsEmpowerCreateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsEmpowerCreate")
	r.SetProduct(Product)
}

type OnsEmpowerCreateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsEmpowerCreate(req *OnsEmpowerCreateRequest, accessId, accessSecret string) (*OnsEmpowerCreateResponse, error) {
	var pResponse OnsEmpowerCreateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
