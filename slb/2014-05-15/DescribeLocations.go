package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLocationsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Namespace            string
	NamespaceUid         string
}

func (r *DescribeLocationsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLocationsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLocationsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLocationsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLocationsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLocationsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLocationsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLocationsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeLocationsRequest) SetNamespace(value string) {
	r.Namespace = value
	r.QueryParams.Set("Namespace", value)
}
func (r *DescribeLocationsRequest) GetNamespace() string {
	return r.Namespace
}
func (r *DescribeLocationsRequest) SetNamespaceUid(value string) {
	r.NamespaceUid = value
	r.QueryParams.Set("NamespaceUid", value)
}
func (r *DescribeLocationsRequest) GetNamespaceUid() string {
	return r.NamespaceUid
}

func (r *DescribeLocationsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLocations")
	r.SetProduct(Product)
}

type DescribeLocationsResponse struct {
}

func DescribeLocations(req *DescribeLocationsRequest, accessId, accessSecret string) (*DescribeLocationsResponse, error) {
	var pResponse DescribeLocationsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
