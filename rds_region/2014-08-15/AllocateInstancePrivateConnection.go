package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AllocateInstancePrivateConnectionRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	DBInstanceId           string
	ConnectionStringPrefix string
	Port                   string
	OwnerAccount           string
}

func (r *AllocateInstancePrivateConnectionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AllocateInstancePrivateConnectionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AllocateInstancePrivateConnectionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AllocateInstancePrivateConnectionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AllocateInstancePrivateConnectionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AllocateInstancePrivateConnectionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AllocateInstancePrivateConnectionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *AllocateInstancePrivateConnectionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *AllocateInstancePrivateConnectionRequest) SetConnectionStringPrefix(value string) {
	r.ConnectionStringPrefix = value
	r.QueryParams.Set("ConnectionStringPrefix", value)
}
func (r *AllocateInstancePrivateConnectionRequest) GetConnectionStringPrefix() string {
	return r.ConnectionStringPrefix
}
func (r *AllocateInstancePrivateConnectionRequest) SetPort(value string) {
	r.Port = value
	r.QueryParams.Set("Port", value)
}
func (r *AllocateInstancePrivateConnectionRequest) GetPort() string {
	return r.Port
}
func (r *AllocateInstancePrivateConnectionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AllocateInstancePrivateConnectionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AllocateInstancePrivateConnectionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AllocateInstancePrivateConnection")
	r.SetProduct(Product)
}

type AllocateInstancePrivateConnectionResponse struct {
}

func AllocateInstancePrivateConnection(req *AllocateInstancePrivateConnectionRequest, accessId, accessSecret string) (*AllocateInstancePrivateConnectionResponse, error) {
	var pResponse AllocateInstancePrivateConnectionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
