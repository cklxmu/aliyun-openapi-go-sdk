package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessageGetByMsgIdRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	MsgId        string
}

func (r *OnsMessageGetByMsgIdRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessageGetByMsgIdRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessageGetByMsgIdRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessageGetByMsgIdRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessageGetByMsgIdRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessageGetByMsgIdRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessageGetByMsgIdRequest) SetMsgId(value string) {
	r.MsgId = value
	r.QueryParams.Set("MsgId", value)
}
func (r *OnsMessageGetByMsgIdRequest) GetMsgId() string {
	return r.MsgId
}

func (r *OnsMessageGetByMsgIdRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessageGetByMsgId")
	r.SetProduct(Product)
}

type OnsMessageGetByMsgIdResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
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
	} `xml:"Data" json:"Data"`
}

func OnsMessageGetByMsgId(req *OnsMessageGetByMsgIdRequest, accessId, accessSecret string) (*OnsMessageGetByMsgIdResponse, error) {
	var pResponse OnsMessageGetByMsgIdResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
