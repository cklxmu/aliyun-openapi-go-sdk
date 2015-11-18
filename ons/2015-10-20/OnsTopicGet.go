package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTopicGetRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
}

func (r *OnsTopicGetRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTopicGetRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTopicGetRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTopicGetRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTopicGetRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTopicGetRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTopicGetRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTopicGetRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsTopicGetRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTopicGet")
	r.SetProduct(Product)
}

type OnsTopicGetResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		PublishInfoDo []struct {
			Id           int    `xml:"Id" json:"Id"`
			ChannelId    int    `xml:"ChannelId" json:"ChannelId"`
			ChannelName  string `xml:"ChannelName" json:"ChannelName"`
			RegionId     string `xml:"RegionId" json:"RegionId"`
			RegionName   string `xml:"RegionName" json:"RegionName"`
			Topic        string `xml:"Topic" json:"Topic"`
			Owner        string `xml:"Owner" json:"Owner"`
			Relation     int    `xml:"Relation" json:"Relation"`
			RelationName string `xml:"RelationName" json:"RelationName"`
			Status       int    `xml:"Status" json:"Status"`
			StatusName   string `xml:"StatusName" json:"StatusName"`
			Appkey       int    `xml:"Appkey" json:"Appkey"`
			CreateTime   int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime   int    `xml:"UpdateTime" json:"UpdateTime"`
			Remark       string `xml:"Remark" json:"Remark"`
		} `xml:"PublishInfoDo" json:"PublishInfoDo"`
	} `xml:"Data" json:"Data"`
}

func OnsTopicGet(req *OnsTopicGetRequest, accessId, accessSecret string) (*OnsTopicGetResponse, error) {
	var pResponse OnsTopicGetResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
