package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceConfigRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *DescribeInstanceConfigRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceConfigRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceConfigRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceConfigRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceConfigRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceConfigRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceConfigRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceConfigRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeInstanceConfigRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeInstanceConfigRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DescribeInstanceConfigRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceConfig")
	r.SetProduct(Product)
}

type DescribeInstanceConfigResponse struct {
	Config string `xml:"Config" json:"Config"`
}

func DescribeInstanceConfig(req *DescribeInstanceConfigRequest, accessId, accessSecret string) (*DescribeInstanceConfigResponse, error) {
	var pResponse DescribeInstanceConfigResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
