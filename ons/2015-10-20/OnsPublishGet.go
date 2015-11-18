package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsPublishGetRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ProducerId   string
	Topic        string
}

func (r *OnsPublishGetRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsPublishGetRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsPublishGetRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsPublishGetRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsPublishGetRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsPublishGetRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsPublishGetRequest) SetProducerId(value string) {
	r.ProducerId = value
	r.QueryParams.Set("ProducerId", value)
}
func (r *OnsPublishGetRequest) GetProducerId() string {
	return r.ProducerId
}
func (r *OnsPublishGetRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsPublishGetRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsPublishGetRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsPublishGet")
	r.SetProduct(Product)
}

type OnsPublishGetResponse struct {
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

func OnsPublishGet(req *OnsPublishGetRequest, accessId, accessSecret string) (*OnsPublishGetResponse, error) {
	var pResponse OnsPublishGetResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
