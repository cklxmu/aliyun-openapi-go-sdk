package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsClusterNamesRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
}

func (r *OnsClusterNamesRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsClusterNamesRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsClusterNamesRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsClusterNamesRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsClusterNamesRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsClusterNamesRequest) GetPreventCache() int {
	return r.PreventCache
}

func (r *OnsClusterNamesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsClusterNames")
	r.SetProduct(Product)
}

type OnsClusterNamesResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		ClusterName []string `xml:"ClusterName" json:"ClusterName"`
	} `xml:"Data" json:"Data"`
}

func OnsClusterNames(req *OnsClusterNamesRequest, accessId, accessSecret string) (*OnsClusterNamesResponse, error) {
	var pResponse OnsClusterNamesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
