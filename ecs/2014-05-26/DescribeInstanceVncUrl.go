package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceVncUrlRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	OwnerAccount         string
}

func (r *DescribeInstanceVncUrlRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceVncUrlRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceVncUrlRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceVncUrlRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceVncUrlRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceVncUrlRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceVncUrlRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeInstanceVncUrlRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeInstanceVncUrlRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceVncUrlRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceVncUrlRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceVncUrl")
	r.SetProduct(Product)
}

type DescribeInstanceVncUrlResponse struct {
	VncUrl string `xml:"VncUrl" json:"VncUrl"`
}

func DescribeInstanceVncUrl(req *DescribeInstanceVncUrlRequest, accessId, accessSecret string) (*DescribeInstanceVncUrlResponse, error) {
	var pResponse DescribeInstanceVncUrlResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
