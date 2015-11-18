package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsPublishSearchRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Search       string
}

func (r *OnsPublishSearchRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsPublishSearchRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsPublishSearchRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsPublishSearchRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsPublishSearchRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsPublishSearchRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsPublishSearchRequest) SetSearch(value string) {
	r.Search = value
	r.QueryParams.Set("Search", value)
}
func (r *OnsPublishSearchRequest) GetSearch() string {
	return r.Search
}

func (r *OnsPublishSearchRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsPublishSearch")
	r.SetProduct(Product)
}

type OnsPublishSearchResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		PublishInfoDo []struct {
			Id          int    `xml:"Id" json:"Id"`
			ChannelId   int    `xml:"ChannelId" json:"ChannelId"`
			ChannelName string `xml:"ChannelName" json:"ChannelName"`
			RegionId    string `xml:"RegionId" json:"RegionId"`
			RegionName  string `xml:"RegionName" json:"RegionName"`
			Owner       string `xml:"Owner" json:"Owner"`
			ProducerId  string `xml:"ProducerId" json:"ProducerId"`
			Topic       string `xml:"Topic" json:"Topic"`
			Status      int    `xml:"Status" json:"Status"`
			StatusName  string `xml:"StatusName" json:"StatusName"`
			CreateTime  int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime  int    `xml:"UpdateTime" json:"UpdateTime"`
		} `xml:"PublishInfoDo" json:"PublishInfoDo"`
	} `xml:"Data" json:"Data"`
}

func OnsPublishSearch(req *OnsPublishSearchRequest, accessId, accessSecret string) (*OnsPublishSearchResponse, error) {
	var pResponse OnsPublishSearchResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
