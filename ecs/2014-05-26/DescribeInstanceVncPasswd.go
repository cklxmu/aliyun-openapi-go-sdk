package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceVncPasswdRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	OwnerAccount         string
}

func (r *DescribeInstanceVncPasswdRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceVncPasswdRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceVncPasswdRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceVncPasswdRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceVncPasswdRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceVncPasswdRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceVncPasswdRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeInstanceVncPasswdRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeInstanceVncPasswdRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceVncPasswdRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceVncPasswdRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceVncPasswd")
	r.SetProduct(Product)
}

type DescribeInstanceVncPasswdResponse struct {
	VncPasswd string `xml:"VncPasswd" json:"VncPasswd"`
}

func DescribeInstanceVncPasswd(req *DescribeInstanceVncPasswdRequest, accessId, accessSecret string) (*DescribeInstanceVncPasswdResponse, error) {
	var pResponse DescribeInstanceVncPasswdResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
