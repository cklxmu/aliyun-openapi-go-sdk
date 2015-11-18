package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsRegionListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
}

func (r *OnsRegionListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsRegionListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsRegionListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsRegionListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsRegionListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsRegionListRequest) GetPreventCache() int {
	return r.PreventCache
}

func (r *OnsRegionListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsRegionList")
	r.SetProduct(Product)
}

type OnsRegionListResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		RegionDo []struct {
			Id          int    `xml:"Id" json:"Id"`
			RegionId    string `xml:"RegionId" json:"RegionId"`
			RegionName  string `xml:"RegionName" json:"RegionName"`
			ChannelId   int    `xml:"ChannelId" json:"ChannelId"`
			ChannelName string `xml:"ChannelName" json:"ChannelName"`
			Owner       string `xml:"Owner" json:"Owner"`
			Cluster     string `xml:"Cluster" json:"Cluster"`
			Status      int    `xml:"Status" json:"Status"`
			IsShare     int    `xml:"IsShare" json:"IsShare"`
			CreateTime  int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime  int    `xml:"UpdateTime" json:"UpdateTime"`
		} `xml:"RegionDo" json:"RegionDo"`
	} `xml:"Data" json:"Data"`
}

func OnsRegionList(req *OnsRegionListRequest, accessId, accessSecret string) (*OnsRegionListResponse, error) {
	var pResponse OnsRegionListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
