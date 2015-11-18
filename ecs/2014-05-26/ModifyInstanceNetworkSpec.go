package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceNetworkSpecRequest struct {
	RpcRequest
	OwnerId                 int
	ResourceOwnerAccount    string
	ResourceOwnerId         int
	InstanceId              string
	InternetMaxBandwidthOut int
	InternetMaxBandwidthIn  int
	OwnerAccount            string
}

func (r *ModifyInstanceNetworkSpecRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceNetworkSpecRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceNetworkSpecRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceNetworkSpecRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceNetworkSpecRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceNetworkSpecRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceNetworkSpecRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceNetworkSpecRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(value int) {
	r.InternetMaxBandwidthOut = value
	r.QueryParams.Set("InternetMaxBandwidthOut", strconv.Itoa(value))
}
func (r *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() int {
	return r.InternetMaxBandwidthOut
}
func (r *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthIn(value int) {
	r.InternetMaxBandwidthIn = value
	r.QueryParams.Set("InternetMaxBandwidthIn", strconv.Itoa(value))
}
func (r *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthIn() int {
	return r.InternetMaxBandwidthIn
}
func (r *ModifyInstanceNetworkSpecRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceNetworkSpecRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyInstanceNetworkSpecRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceNetworkSpec")
	r.SetProduct(Product)
}

type ModifyInstanceNetworkSpecResponse struct {
}

func ModifyInstanceNetworkSpec(req *ModifyInstanceNetworkSpecRequest, accessId, accessSecret string) (*ModifyInstanceNetworkSpecResponse, error) {
	var pResponse ModifyInstanceNetworkSpecResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
