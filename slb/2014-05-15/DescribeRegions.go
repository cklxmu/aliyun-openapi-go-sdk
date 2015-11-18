package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRegionsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeRegionsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeRegionsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeRegionsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeRegionsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeRegionsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeRegionsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
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
			RegionId  string `xml:"RegionId" json:"RegionId"`
			LocalName string `xml:"LocalName" json:"LocalName"`
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
