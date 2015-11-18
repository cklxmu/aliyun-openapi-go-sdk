package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AddressType          string
	InternetChargeType   string
	Bandwidth            int
	ClientToken          string
	LoadBalancerName     string
	VpcId                string
	VSwitchId            string
	OwnerAccount         string
	MasterZoneId         string
	SlaveZoneId          string
	MaxConnLimit         int
	SecurityStatus       string
}

func (r *CreateLoadBalancerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerRequest) SetAddressType(value string) {
	r.AddressType = value
	r.QueryParams.Set("AddressType", value)
}
func (r *CreateLoadBalancerRequest) GetAddressType() string {
	return r.AddressType
}
func (r *CreateLoadBalancerRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *CreateLoadBalancerRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *CreateLoadBalancerRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *CreateLoadBalancerRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateLoadBalancerRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateLoadBalancerRequest) SetLoadBalancerName(value string) {
	r.LoadBalancerName = value
	r.QueryParams.Set("LoadBalancerName", value)
}
func (r *CreateLoadBalancerRequest) GetLoadBalancerName() string {
	return r.LoadBalancerName
}
func (r *CreateLoadBalancerRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateLoadBalancerRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateLoadBalancerRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *CreateLoadBalancerRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *CreateLoadBalancerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateLoadBalancerRequest) SetMasterZoneId(value string) {
	r.MasterZoneId = value
	r.QueryParams.Set("MasterZoneId", value)
}
func (r *CreateLoadBalancerRequest) GetMasterZoneId() string {
	return r.MasterZoneId
}
func (r *CreateLoadBalancerRequest) SetSlaveZoneId(value string) {
	r.SlaveZoneId = value
	r.QueryParams.Set("SlaveZoneId", value)
}
func (r *CreateLoadBalancerRequest) GetSlaveZoneId() string {
	return r.SlaveZoneId
}
func (r *CreateLoadBalancerRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *CreateLoadBalancerRequest) SetSecurityStatus(value string) {
	r.SecurityStatus = value
	r.QueryParams.Set("SecurityStatus", value)
}
func (r *CreateLoadBalancerRequest) GetSecurityStatus() string {
	return r.SecurityStatus
}

func (r *CreateLoadBalancerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancer")
	r.SetProduct(Product)
}

type CreateLoadBalancerResponse struct {
	LoadBalancerId   string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	Address          string `xml:"Address" json:"Address"`
	LoadBalancerName string `xml:"LoadBalancerName" json:"LoadBalancerName"`
	VpcId            string `xml:"VpcId" json:"VpcId"`
	VSwitchId        string `xml:"VSwitchId" json:"VSwitchId"`
	NetworkType      string `xml:"NetworkType" json:"NetworkType"`
}

func CreateLoadBalancer(req *CreateLoadBalancerRequest, accessId, accessSecret string) (*CreateLoadBalancerResponse, error) {
	var pResponse CreateLoadBalancerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
