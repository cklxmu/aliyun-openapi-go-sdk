package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessageGetByTopicRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
}

func (r *OnsMessageGetByTopicRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessageGetByTopicRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessageGetByTopicRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessageGetByTopicRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessageGetByTopicRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessageGetByTopicRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessageGetByTopicRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsMessageGetByTopicRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsMessageGetByTopicRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessageGetByTopic")
	r.SetProduct(Product)
}

type OnsMessageGetByTopicResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		OnsRestMessageDo []struct {
			Topic                     string `xml:"Topic" json:"Topic"`
			Flag                      int    `xml:"Flag" json:"Flag"`
			Body                      string `xml:"Body" json:"Body"`
			QueueId                   int    `xml:"QueueId" json:"QueueId"`
			StoreSize                 int    `xml:"StoreSize" json:"StoreSize"`
			QueueOffset               int    `xml:"QueueOffset" json:"QueueOffset"`
			SysFlag                   int    `xml:"SysFlag" json:"SysFlag"`
			BornTimestamp             int    `xml:"BornTimestamp" json:"BornTimestamp"`
			BornHost                  string `xml:"BornHost" json:"BornHost"`
			StoreTimestamp            int    `xml:"StoreTimestamp" json:"StoreTimestamp"`
			StoreHost                 string `xml:"StoreHost" json:"StoreHost"`
			MsgId                     string `xml:"MsgId" json:"MsgId"`
			CommitLogOffset           int    `xml:"CommitLogOffset" json:"CommitLogOffset"`
			BodyCRC                   int    `xml:"BodyCRC" json:"BodyCRC"`
			ReconsumeTimes            int    `xml:"ReconsumeTimes" json:"ReconsumeTimes"`
			PreparedTransactionOffset int    `xml:"PreparedTransactionOffset" json:"PreparedTransactionOffset"`
			PropertyList              struct {
				MessageProperty []struct {
					Name  string `xml:"Name" json:"Name"`
					Value string `xml:"Value" json:"Value"`
				} `xml:"MessageProperty" json:"MessageProperty"`
			} `xml:"PropertyList" json:"PropertyList"`
		} `xml:"OnsRestMessageDo" json:"OnsRestMessageDo"`
	} `xml:"Data" json:"Data"`
}

func OnsMessageGetByTopic(req *OnsMessageGetByTopicRequest, accessId, accessSecret string) (*OnsMessageGetByTopicResponse, error) {
	var pResponse OnsMessageGetByTopicResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
