package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRegionsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
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
func (r *DescribeRegionsRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeRegionsRequest) GetClientToken() string {
	return r.ClientToken
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
		RDSRegion []struct {
			RegionId string `xml:"RegionId" json:"RegionId"`
			ZoneId   string `xml:"ZoneId" json:"ZoneId"`
		} `xml:"RDSRegion" json:"RDSRegion"`
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
