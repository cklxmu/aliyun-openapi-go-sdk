package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTopicStatusRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	Detail       bool
}

func (r *OnsTopicStatusRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTopicStatusRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTopicStatusRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTopicStatusRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTopicStatusRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTopicStatusRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTopicStatusRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTopicStatusRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsTopicStatusRequest) SetDetail(value bool) {
	r.Detail = value
	r.QueryParams.Set("Detail", strconv.FormatBool(value))
}
func (r *OnsTopicStatusRequest) GetDetail() bool {
	return r.Detail
}

func (r *OnsTopicStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTopicStatus")
	r.SetProduct(Product)
}

type OnsTopicStatusResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		TotalCount    int `xml:"TotalCount" json:"TotalCount"`
		LastTimeStamp int `xml:"LastTimeStamp" json:"LastTimeStamp"`
		OffsetTable   struct {
			TopicQueueOffset []struct {
				Topic               string `xml:"Topic" json:"Topic"`
				BrokerName          string `xml:"BrokerName" json:"BrokerName"`
				QueueId             int    `xml:"QueueId" json:"QueueId"`
				MinOffset           int    `xml:"MinOffset" json:"MinOffset"`
				MaxOffset           int    `xml:"MaxOffset" json:"MaxOffset"`
				LastUpdateTimestamp int    `xml:"LastUpdateTimestamp" json:"LastUpdateTimestamp"`
			} `xml:"TopicQueueOffset" json:"TopicQueueOffset"`
		} `xml:"OffsetTable" json:"OffsetTable"`
	} `xml:"Data" json:"Data"`
}

func OnsTopicStatus(req *OnsTopicStatusRequest, accessId, accessSecret string) (*OnsTopicStatusResponse, error) {
	var pResponse OnsTopicStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
