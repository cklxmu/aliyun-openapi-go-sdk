package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsSubscriptionSearchRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Search       string
}

func (r *OnsSubscriptionSearchRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsSubscriptionSearchRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsSubscriptionSearchRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsSubscriptionSearchRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsSubscriptionSearchRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsSubscriptionSearchRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsSubscriptionSearchRequest) SetSearch(value string) {
	r.Search = value
	r.QueryParams.Set("Search", value)
}
func (r *OnsSubscriptionSearchRequest) GetSearch() string {
	return r.Search
}

func (r *OnsSubscriptionSearchRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsSubscriptionSearch")
	r.SetProduct(Product)
}

type OnsSubscriptionSearchResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		SubscribeInfoDo []struct {
			Id          int    `xml:"Id" json:"Id"`
			ChannelId   int    `xml:"ChannelId" json:"ChannelId"`
			ChannelName string `xml:"ChannelName" json:"ChannelName"`
			RegionId    string `xml:"RegionId" json:"RegionId"`
			RegionName  string `xml:"RegionName" json:"RegionName"`
			Owner       string `xml:"Owner" json:"Owner"`
			ConsumerId  string `xml:"ConsumerId" json:"ConsumerId"`
			Topic       string `xml:"Topic" json:"Topic"`
			Status      int    `xml:"Status" json:"Status"`
			StatusName  string `xml:"StatusName" json:"StatusName"`
			CreateTime  int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime  int    `xml:"UpdateTime" json:"UpdateTime"`
		} `xml:"SubscribeInfoDo" json:"SubscribeInfoDo"`
	} `xml:"Data" json:"Data"`
}

func OnsSubscriptionSearch(req *OnsSubscriptionSearchRequest, accessId, accessSecret string) (*OnsSubscriptionSearchResponse, error) {
	var pResponse OnsSubscriptionSearchResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
