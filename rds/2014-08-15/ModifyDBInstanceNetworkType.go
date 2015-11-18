package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceNetworkTypeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	InstanceNetworkType  string
	VPCId                string
	VSwitchId            string
	PrivateIpAddress     string
	OwnerAccount         string
}

func (r *ModifyDBInstanceNetworkTypeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(value string) {
	r.InstanceNetworkType = value
	r.QueryParams.Set("InstanceNetworkType", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetInstanceNetworkType() string {
	return r.InstanceNetworkType
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetVPCId(value string) {
	r.VPCId = value
	r.QueryParams.Set("VPCId", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetVPCId() string {
	return r.VPCId
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(value string) {
	r.PrivateIpAddress = value
	r.QueryParams.Set("PrivateIpAddress", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetPrivateIpAddress() string {
	return r.PrivateIpAddress
}
func (r *ModifyDBInstanceNetworkTypeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceNetworkTypeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceNetworkTypeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceNetworkType")
	r.SetProduct(Product)
}

type ModifyDBInstanceNetworkTypeResponse struct {
}

func ModifyDBInstanceNetworkType(req *ModifyDBInstanceNetworkTypeRequest, accessId, accessSecret string) (*ModifyDBInstanceNetworkTypeResponse, error) {
	var pResponse ModifyDBInstanceNetworkTypeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
