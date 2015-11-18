package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsCloudGetAppkeyListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	IsvId        int
}

func (r *OnsCloudGetAppkeyListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsCloudGetAppkeyListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsCloudGetAppkeyListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsCloudGetAppkeyListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsCloudGetAppkeyListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsCloudGetAppkeyListRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsCloudGetAppkeyListRequest) SetIsvId(value int) {
	r.IsvId = value
	r.QueryParams.Set("IsvId", strconv.Itoa(value))
}
func (r *OnsCloudGetAppkeyListRequest) GetIsvId() int {
	return r.IsvId
}

func (r *OnsCloudGetAppkeyListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsCloudGetAppkeyList")
	r.SetProduct(Product)
}

type OnsCloudGetAppkeyListResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		AppStatus    int    `xml:"AppStatus" json:"AppStatus"`
		Appkey       string `xml:"Appkey" json:"Appkey"`
		Description  string `xml:"Description" json:"Description"`
		IsvId        int    `xml:"IsvId" json:"IsvId"`
		SupportEmail string `xml:"SupportEmail" json:"SupportEmail"`
		Title        string `xml:"Title" json:"Title"`
		Wangwang     string `xml:"Wangwang" json:"Wangwang"`
	} `xml:"Data" json:"Data"`
}

func OnsCloudGetAppkeyList(req *OnsCloudGetAppkeyListRequest, accessId, accessSecret string) (*OnsCloudGetAppkeyListResponse, error) {
	var pResponse OnsCloudGetAppkeyListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
