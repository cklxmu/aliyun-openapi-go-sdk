package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsConsumerTimeSpanRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsConsumerTimeSpanRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsConsumerTimeSpanRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsConsumerTimeSpanRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsConsumerTimeSpanRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsConsumerTimeSpanRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsConsumerTimeSpanRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsConsumerTimeSpanRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsConsumerTimeSpanRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsConsumerTimeSpanRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsConsumerTimeSpanRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsConsumerTimeSpanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsConsumerTimeSpan")
	r.SetProduct(Product)
}

type OnsConsumerTimeSpanResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		Topic            string `xml:"Topic" json:"Topic"`
		MinTimeStamp     int    `xml:"MinTimeStamp" json:"MinTimeStamp"`
		MaxTimeStamp     int    `xml:"MaxTimeStamp" json:"MaxTimeStamp"`
		ConsumeTimeStamp int    `xml:"ConsumeTimeStamp" json:"ConsumeTimeStamp"`
		QueueTimeSpans   struct {
			QueueTimeSpan []struct {
				MinTimeStamp     int `xml:"MinTimeStamp" json:"MinTimeStamp"`
				MaxTimeStamp     int `xml:"MaxTimeStamp" json:"MaxTimeStamp"`
				ConsumeTimeStamp int `xml:"ConsumeTimeStamp" json:"ConsumeTimeStamp"`
				MessageQueue     struct {
					Topic      string `xml:"Topic" json:"Topic"`
					BrokerName string `xml:"BrokerName" json:"BrokerName"`
					QueueId    int    `xml:"QueueId" json:"QueueId"`
				} `xml:"MessageQueue" json:"MessageQueue"`
			} `xml:"QueueTimeSpan" json:"QueueTimeSpan"`
		} `xml:"QueueTimeSpans" json:"QueueTimeSpans"`
	} `xml:"Data" json:"Data"`
}

func OnsConsumerTimeSpan(req *OnsConsumerTimeSpanRequest, accessId, accessSecret string) (*OnsConsumerTimeSpanResponse, error) {
	var pResponse OnsConsumerTimeSpanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
