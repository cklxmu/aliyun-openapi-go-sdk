package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsSubscriptionGetRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsSubscriptionGetRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsSubscriptionGetRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsSubscriptionGetRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsSubscriptionGetRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsSubscriptionGetRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsSubscriptionGetRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsSubscriptionGetRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsSubscriptionGetRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsSubscriptionGetRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsSubscriptionGetRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsSubscriptionGetRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsSubscriptionGet")
	r.SetProduct(Product)
}

type OnsSubscriptionGetResponse struct {
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

func OnsSubscriptionGet(req *OnsSubscriptionGetRequest, accessId, accessSecret string) (*OnsSubscriptionGetResponse, error) {
	var pResponse OnsSubscriptionGetResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
