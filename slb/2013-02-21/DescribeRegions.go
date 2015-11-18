package slb

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeRegionsRequest struct {
	RpcRequest
	HostId       string
	OwnerAccount string
}

func (r *DescribeRegionsRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *DescribeRegionsRequest) GetHostId() string {
	return r.HostId
}
func (r *DescribeRegionsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeRegionsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeRegionsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeRegions")
	r.SetProduct(Product)
}

type DescribeRegionsResponse struct {
	Regions struct {
		Region []struct {
			RegionId string `xml:"RegionId" json:"RegionId"`
		} `xml:"Region" json:"Region"`
	} `xml:"Regions" json:"Regions"`
}

func DescribeRegions(req *DescribeRegionsRequest, accessId, accessSecret string) (*DescribeRegionsResponse, error) {
	var pResponse DescribeRegionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
