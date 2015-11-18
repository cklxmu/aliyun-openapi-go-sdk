package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AllocateEipAddressRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	Bandwidth            string
	InternetChargeType   string
	OwnerAccount         string
}

func (r *AllocateEipAddressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AllocateEipAddressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AllocateEipAddressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AllocateEipAddressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AllocateEipAddressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AllocateEipAddressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AllocateEipAddressRequest) SetBandwidth(value string) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", value)
}
func (r *AllocateEipAddressRequest) GetBandwidth() string {
	return r.Bandwidth
}
func (r *AllocateEipAddressRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *AllocateEipAddressRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *AllocateEipAddressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AllocateEipAddressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AllocateEipAddressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AllocateEipAddress")
	r.SetProduct(Product)
}

type AllocateEipAddressResponse struct {
	AllocationId string `xml:"AllocationId" json:"AllocationId"`
	EipAddress   string `xml:"EipAddress" json:"EipAddress"`
}

func AllocateEipAddress(req *AllocateEipAddressRequest, accessId, accessSecret string) (*AllocateEipAddressResponse, error) {
	var pResponse AllocateEipAddressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
