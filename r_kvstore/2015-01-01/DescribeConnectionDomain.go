package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeConnectionDomainRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ConnectionDomain     string
}

func (r *DescribeConnectionDomainRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeConnectionDomainRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeConnectionDomainRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeConnectionDomainRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeConnectionDomainRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeConnectionDomainRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeConnectionDomainRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeConnectionDomainRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeConnectionDomainRequest) SetConnectionDomain(value string) {
	r.ConnectionDomain = value
	r.QueryParams.Set("ConnectionDomain", value)
}
func (r *DescribeConnectionDomainRequest) GetConnectionDomain() string {
	return r.ConnectionDomain
}

func (r *DescribeConnectionDomainRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeConnectionDomain")
	r.SetProduct(Product)
}

type DescribeConnectionDomainResponse struct {
	InstanceId string `xml:"InstanceId" json:"InstanceId"`
}

func DescribeConnectionDomain(req *DescribeConnectionDomainRequest, accessId, accessSecret string) (*DescribeConnectionDomainResponse, error) {
	var pResponse DescribeConnectionDomainResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
