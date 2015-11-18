package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceCountRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeInstanceCountRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceCountRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceCountRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceCountRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceCountRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceCountRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceCountRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceCountRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceCountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceCount")
	r.SetProduct(Product)
}

type DescribeInstanceCountResponse struct {
	InstancesCount int `xml:"InstancesCount" json:"InstancesCount"`
}

func DescribeInstanceCount(req *DescribeInstanceCountRequest, accessId, accessSecret string) (*DescribeInstanceCountResponse, error) {
	var pResponse DescribeInstanceCountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
