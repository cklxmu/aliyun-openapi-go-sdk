package dts

import (
	. "aliyun-openapi-go-sdk/core"
)

type drcGuidRouteApiRequest struct {
	RpcRequest
	guid string
}

func (r *drcGuidRouteApiRequest) Setguid(value string) {
	r.guid = value
	r.QueryParams.Set("guid", value)
}
func (r *drcGuidRouteApiRequest) Getguid() string {
	return r.guid
}

func (r *drcGuidRouteApiRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("drcGuidRouteApi")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type drcGuidRouteApiResponse struct {
	success bool   `xml:"success" json:"success"`
	data    string `xml:"data" json:"data"`
}

func drcGuidRouteApi(req *drcGuidRouteApiRequest, accessId, accessSecret string) (*drcGuidRouteApiResponse, error) {
	var pResponse drcGuidRouteApiResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
