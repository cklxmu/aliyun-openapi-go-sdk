package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessageGetByKeyRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	Key          string
}

func (r *OnsMessageGetByKeyRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessageGetByKeyRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessageGetByKeyRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessageGetByKeyRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessageGetByKeyRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessageGetByKeyRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessageGetByKeyRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsMessageGetByKeyRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsMessageGetByKeyRequest) SetKey(value string) {
	r.Key = value
	r.QueryParams.Set("Key", value)
}
func (r *OnsMessageGetByKeyRequest) GetKey() string {
	return r.Key
}

func (r *OnsMessageGetByKeyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessageGetByKey")
	r.SetProduct(Product)
}

type OnsMessageGetByKeyResponse struct {
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

func OnsMessageGetByKey(req *OnsMessageGetByKeyRequest, accessId, accessSecret string) (*OnsMessageGetByKeyResponse, error) {
	var pResponse OnsMessageGetByKeyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
