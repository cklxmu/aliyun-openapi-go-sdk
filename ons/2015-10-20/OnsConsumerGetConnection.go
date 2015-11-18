package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsConsumerGetConnectionRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
}

func (r *OnsConsumerGetConnectionRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsConsumerGetConnectionRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsConsumerGetConnectionRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsConsumerGetConnectionRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsConsumerGetConnectionRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsConsumerGetConnectionRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsConsumerGetConnectionRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsConsumerGetConnectionRequest) GetConsumerId() string {
	return r.ConsumerId
}

func (r *OnsConsumerGetConnectionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsConsumerGetConnection")
	r.SetProduct(Product)
}

type OnsConsumerGetConnectionResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		ConnectionList struct {
			ConnectionDo []struct {
				ClientId   string `xml:"ClientId" json:"ClientId"`
				ClientAddr string `xml:"ClientAddr" json:"ClientAddr"`
				Language   string `xml:"Language" json:"Language"`
				Version    string `xml:"Version" json:"Version"`
			} `xml:"ConnectionDo" json:"ConnectionDo"`
		} `xml:"ConnectionList" json:"ConnectionList"`
	} `xml:"Data" json:"Data"`
}

func OnsConsumerGetConnection(req *OnsConsumerGetConnectionRequest, accessId, accessSecret string) (*OnsConsumerGetConnectionResponse, error) {
	var pResponse OnsConsumerGetConnectionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
