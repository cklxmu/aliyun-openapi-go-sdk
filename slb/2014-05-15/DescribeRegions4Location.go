package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRegions4LocationRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeRegions4LocationRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeRegions4LocationRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeRegions4LocationRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeRegions4LocationRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeRegions4LocationRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeRegions4LocationRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeRegions4LocationRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeRegions4LocationRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeRegions4LocationRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeRegions4Location")
	r.SetProduct(Product)
}

type DescribeRegions4LocationResponse struct {
	Regions struct {
		Region []struct {
			RegionId  string `xml:"RegionId" json:"RegionId"`
			LocalName string `xml:"LocalName" json:"LocalName"`
		} `xml:"Region" json:"Region"`
	} `xml:"Regions" json:"Regions"`
}

func DescribeRegions4Location(req *DescribeRegions4LocationRequest, accessId, accessSecret string) (*DescribeRegions4LocationResponse, error) {
	var pResponse DescribeRegions4LocationResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
