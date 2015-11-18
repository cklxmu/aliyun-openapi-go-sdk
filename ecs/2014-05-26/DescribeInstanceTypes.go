package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceTypesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeInstanceTypesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceTypesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceTypesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceTypesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceTypesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceTypesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceTypesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceTypesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceTypesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceTypes")
	r.SetProduct(Product)
}

type DescribeInstanceTypesResponse struct {
	InstanceTypes struct {
		InstanceType []struct {
			InstanceTypeId string  `xml:"InstanceTypeId" json:"InstanceTypeId"`
			CpuCoreCount   int     `xml:"CpuCoreCount" json:"CpuCoreCount"`
			MemorySize     float32 `xml:"MemorySize" json:"MemorySize"`
		} `xml:"InstanceType" json:"InstanceType"`
	} `xml:"InstanceTypes" json:"InstanceTypes"`
}

func DescribeInstanceTypes(req *DescribeInstanceTypesRequest, accessId, accessSecret string) (*DescribeInstanceTypesResponse, error) {
	var pResponse DescribeInstanceTypesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
