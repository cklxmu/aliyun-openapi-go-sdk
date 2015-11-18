package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsMessageTraceRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	MsgId        string
}

func (r *OnsMessageTraceRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsMessageTraceRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsMessageTraceRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsMessageTraceRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsMessageTraceRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsMessageTraceRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsMessageTraceRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsMessageTraceRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsMessageTraceRequest) SetMsgId(value string) {
	r.MsgId = value
	r.QueryParams.Set("MsgId", value)
}
func (r *OnsMessageTraceRequest) GetMsgId() string {
	return r.MsgId
}

func (r *OnsMessageTraceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsMessageTrace")
	r.SetProduct(Product)
}

type OnsMessageTraceResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		MessageTrack []struct {
			ConsumerGroup string `xml:"ConsumerGroup" json:"ConsumerGroup"`
			TrackType     string `xml:"TrackType" json:"TrackType"`
			ExceptionDesc string `xml:"ExceptionDesc" json:"ExceptionDesc"`
		} `xml:"MessageTrack" json:"MessageTrack"`
	} `xml:"Data" json:"Data"`
}

func OnsMessageTrace(req *OnsMessageTraceRequest, accessId, accessSecret string) (*OnsMessageTraceResponse, error) {
	var pResponse OnsMessageTraceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
