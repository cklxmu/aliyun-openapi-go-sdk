package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsWarnListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Topic        string
}

func (r *OnsWarnListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsWarnListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsWarnListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsWarnListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsWarnListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsWarnListRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsWarnListRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsWarnListRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsWarnListRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsWarnListRequest) GetTopic() string {
	return r.Topic
}

func (r *OnsWarnListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsWarnList")
	r.SetProduct(Product)
}

type OnsWarnListResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		WarnViewDO []struct {
			AliyunPK   string `xml:"AliyunPK" json:"AliyunPK"`
			ConsumerID string `xml:"ConsumerID" json:"ConsumerID"`
			Topic      string `xml:"Topic" json:"Topic"`
			Threshold  string `xml:"Threshold" json:"Threshold"`
			Status     bool   `xml:"Status" json:"Status"`
			contacts   struct {
				YunContact []struct {
					Name   string `xml:"Name" json:"Name"`
					Value  string `xml:"Value" json:"Value"`
					Type   string `xml:"Type" json:"Type"`
					Params string `xml:"Params" json:"Params"`
					Ari    string `xml:"Ari" json:"Ari"`
				} `xml:"YunContact" json:"YunContact"`
			} `xml:"contacts" json:"contacts"`
		} `xml:"WarnViewDO" json:"WarnViewDO"`
	} `xml:"Data" json:"Data"`
}

func OnsWarnList(req *OnsWarnListRequest, accessId, accessSecret string) (*OnsWarnListResponse, error) {
	var pResponse OnsWarnListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
