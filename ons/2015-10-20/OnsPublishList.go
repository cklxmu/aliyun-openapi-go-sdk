package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsPublishListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
}

func (r *OnsPublishListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsPublishListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsPublishListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsPublishListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsPublishListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsPublishListRequest) GetPreventCache() int {
	return r.PreventCache
}

func (r *OnsPublishListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsPublishList")
	r.SetProduct(Product)
}

type OnsPublishListResponse struct {
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

func OnsPublishList(req *OnsPublishListRequest, accessId, accessSecret string) (*OnsPublishListResponse, error) {
	var pResponse OnsPublishListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
