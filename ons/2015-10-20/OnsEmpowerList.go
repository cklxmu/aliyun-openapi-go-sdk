package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsEmpowerListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	EmpowerUser  string
	Topic        string
}

func (r *OnsEmpowerListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsEmpowerListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsEmpowerListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsEmpowerListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsEmpowerListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsEmpowerListRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsEmpowerListRequest) SetEmpowerUser(value string) {
	r.EmpowerUser = value
	r.QueryParams.Set("EmpowerUser", value)
}
func (r *OnsEmpowerListRequest) GetEmpowerUser() string {
	return r.EmpowerUser
}
func (r *OnsEmpowerListRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsEmpowerListRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsEmpowerListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsEmpowerList")
	r.SetProduct(Product)
}

type OnsEmpowerListResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		AuthOwnerInfoDo []struct {
			Id           int    `xml:"Id" json:"Id"`
			Topic        string `xml:"Topic" json:"Topic"`
			Owner        int    `xml:"Owner" json:"Owner"`
			Relation     int    `xml:"Relation" json:"Relation"`
			RelationName string `xml:"RelationName" json:"RelationName"`
			CreateTime   int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime   int    `xml:"UpdateTime" json:"UpdateTime"`
		} `xml:"AuthOwnerInfoDo" json:"AuthOwnerInfoDo"`
	} `xml:"Data" json:"Data"`
}

func OnsEmpowerList(req *OnsEmpowerListRequest, accessId, accessSecret string) (*OnsEmpowerListResponse, error) {
	var pResponse OnsEmpowerListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
