package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsSubscriptionListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
}

func (r *OnsSubscriptionListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsSubscriptionListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsSubscriptionListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsSubscriptionListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsSubscriptionListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsSubscriptionListRequest) GetPreventCache() int {
	return r.PreventCache
}

func (r *OnsSubscriptionListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsSubscriptionList")
	r.SetProduct(Product)
}

type OnsSubscriptionListResponse struct {
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

func OnsSubscriptionList(req *OnsSubscriptionListRequest, accessId, accessSecret string) (*OnsSubscriptionListResponse, error) {
	var pResponse OnsSubscriptionListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
