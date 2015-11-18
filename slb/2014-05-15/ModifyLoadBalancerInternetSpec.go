package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyLoadBalancerInternetSpecRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	InternetChargeType   string
	Bandwidth            int
	OwnerAccount         string
	MasterZoneId         string
	SlaveZoneId          string
	MaxConnLimit         int
	SecurityStatus       string
}

func (r *ModifyLoadBalancerInternetSpecRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetMasterZoneId(value string) {
	r.MasterZoneId = value
	r.QueryParams.Set("MasterZoneId", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetMasterZoneId() string {
	return r.MasterZoneId
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetSlaveZoneId(value string) {
	r.SlaveZoneId = value
	r.QueryParams.Set("SlaveZoneId", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetSlaveZoneId() string {
	return r.SlaveZoneId
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *ModifyLoadBalancerInternetSpecRequest) SetSecurityStatus(value string) {
	r.SecurityStatus = value
	r.QueryParams.Set("SecurityStatus", value)
}
func (r *ModifyLoadBalancerInternetSpecRequest) GetSecurityStatus() string {
	return r.SecurityStatus
}

func (r *ModifyLoadBalancerInternetSpecRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyLoadBalancerInternetSpec")
	r.SetProduct(Product)
}

type ModifyLoadBalancerInternetSpecResponse struct {
}

func ModifyLoadBalancerInternetSpec(req *ModifyLoadBalancerInternetSpecRequest, accessId, accessSecret string) (*ModifyLoadBalancerInternetSpecResponse, error) {
	var pResponse ModifyLoadBalancerInternetSpecResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
