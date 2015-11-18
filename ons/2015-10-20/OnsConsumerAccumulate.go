package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsConsumerAccumulateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Detail       bool
}

func (r *OnsConsumerAccumulateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsConsumerAccumulateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsConsumerAccumulateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsConsumerAccumulateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsConsumerAccumulateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsConsumerAccumulateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsConsumerAccumulateRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsConsumerAccumulateRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsConsumerAccumulateRequest) SetDetail(value bool) {
	r.Detail = value
	r.QueryParams.Set("Detail", strconv.FormatBool(value))
}
func (r *OnsConsumerAccumulateRequest) GetDetail() bool {
	return r.Detail
}

func (r *OnsConsumerAccumulateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsConsumerAccumulate")
	r.SetProduct(Product)
}

type OnsConsumerAccumulateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		Online            bool    `xml:"Online" json:"Online"`
		TotalDiff         int     `xml:"TotalDiff" json:"TotalDiff"`
		ConsumeTps        float32 `xml:"ConsumeTps" json:"ConsumeTps"`
		LastTimestamp     int     `xml:"LastTimestamp" json:"LastTimestamp"`
		DelayTime         int     `xml:"DelayTime" json:"DelayTime"`
		DetailInTopicList struct {
			DetailInTopicDo []struct {
				Topic         string  `xml:"Topic" json:"Topic"`
				ConsumeTps    float32 `xml:"ConsumeTps" json:"ConsumeTps"`
				LastTimestamp int     `xml:"LastTimestamp" json:"LastTimestamp"`
				DelayTime     int     `xml:"DelayTime" json:"DelayTime"`
				OffsetList    struct {
					ConsumeQueueOffset []struct {
						Topic          string `xml:"Topic" json:"Topic"`
						BrokerName     string `xml:"BrokerName" json:"BrokerName"`
						QueueId        int    `xml:"QueueId" json:"QueueId"`
						BrokerOffset   int    `xml:"BrokerOffset" json:"BrokerOffset"`
						ConsumerOffset int    `xml:"ConsumerOffset" json:"ConsumerOffset"`
						LastTimestamp  int    `xml:"LastTimestamp" json:"LastTimestamp"`
					} `xml:"ConsumeQueueOffset" json:"ConsumeQueueOffset"`
				} `xml:"OffsetList" json:"OffsetList"`
			} `xml:"DetailInTopicDo" json:"DetailInTopicDo"`
		} `xml:"DetailInTopicList" json:"DetailInTopicList"`
	} `xml:"Data" json:"Data"`
}

func OnsConsumerAccumulate(req *OnsConsumerAccumulateRequest, accessId, accessSecret string) (*OnsConsumerAccumulateResponse, error) {
	var pResponse OnsConsumerAccumulateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
