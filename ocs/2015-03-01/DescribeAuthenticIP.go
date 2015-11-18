package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeAuthenticIPRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *DescribeAuthenticIPRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeAuthenticIPRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeAuthenticIPRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeAuthenticIPRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeAuthenticIPRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeAuthenticIPRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeAuthenticIPRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeAuthenticIPRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeAuthenticIPRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeAuthenticIPRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DescribeAuthenticIPRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeAuthenticIP")
	r.SetProduct(Product)
}

type DescribeAuthenticIPResponse struct {
	AuthenticIPs string `xml:"AuthenticIPs" json:"AuthenticIPs"`
}

func DescribeAuthenticIP(req *DescribeAuthenticIPRequest, accessId, accessSecret string) (*DescribeAuthenticIPResponse, error) {
	var pResponse DescribeAuthenticIPResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
