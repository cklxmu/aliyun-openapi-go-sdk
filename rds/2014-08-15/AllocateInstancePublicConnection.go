package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AllocateInstancePublicConnectionRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	DBInstanceId           string
	ConnectionStringPrefix string
	Port                   string
	OwnerAccount           string
}

func (r *AllocateInstancePublicConnectionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AllocateInstancePublicConnectionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AllocateInstancePublicConnectionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AllocateInstancePublicConnectionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AllocateInstancePublicConnectionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *AllocateInstancePublicConnectionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(value string) {
	r.ConnectionStringPrefix = value
	r.QueryParams.Set("ConnectionStringPrefix", value)
}
func (r *AllocateInstancePublicConnectionRequest) GetConnectionStringPrefix() string {
	return r.ConnectionStringPrefix
}
func (r *AllocateInstancePublicConnectionRequest) SetPort(value string) {
	r.Port = value
	r.QueryParams.Set("Port", value)
}
func (r *AllocateInstancePublicConnectionRequest) GetPort() string {
	return r.Port
}
func (r *AllocateInstancePublicConnectionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AllocateInstancePublicConnectionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AllocateInstancePublicConnectionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AllocateInstancePublicConnection")
	r.SetProduct(Product)
}

type AllocateInstancePublicConnectionResponse struct {
}

func AllocateInstancePublicConnection(req *AllocateInstancePublicConnectionRequest, accessId, accessSecret string) (*AllocateInstancePublicConnectionResponse, error) {
	var pResponse AllocateInstancePublicConnectionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
