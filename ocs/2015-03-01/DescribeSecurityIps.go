package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSecurityIpsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *DescribeSecurityIpsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityIpsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSecurityIpsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSecurityIpsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSecurityIpsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityIpsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSecurityIpsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSecurityIpsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeSecurityIpsRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeSecurityIpsRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DescribeSecurityIpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSecurityIps")
	r.SetProduct(Product)
}

type DescribeSecurityIpsResponse struct {
	SecurityIps string `xml:"SecurityIps" json:"SecurityIps"`
}

func DescribeSecurityIps(req *DescribeSecurityIpsRequest, accessId, accessSecret string) (*DescribeSecurityIpsResponse, error) {
	var pResponse DescribeSecurityIpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
