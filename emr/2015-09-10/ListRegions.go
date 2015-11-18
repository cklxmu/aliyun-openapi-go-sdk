package emr

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListRegionsRequest struct {
	RpcRequest
}

func (r *ListRegionsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListRegions")
	r.SetProduct(Product)
}

type ListRegionsResponse struct {
	RegionTypes struct {
		RegionType []struct {
			Name  string `xml:"Name" json:"Name"`
			Id    string `xml:"Id" json:"Id"`
			AsUrl string `xml:"AsUrl" json:"AsUrl"`
		} `xml:"RegionType" json:"RegionType"`
	} `xml:"RegionTypes" json:"RegionTypes"`
}

func ListRegions(req *ListRegionsRequest, accessId, accessSecret string) (*ListRegionsResponse, error) {
	var pResponse ListRegionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
