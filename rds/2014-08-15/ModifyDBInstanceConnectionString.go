package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceConnectionStringRequest struct {
	RpcRequest
	OwnerId                 int
	ResourceOwnerAccount    string
	ResourceOwnerId         int
	DBInstanceId            string
	CurrentConnectionString string
	ConnectionStringPrefix  string
	Port                    string
	OwnerAccount            string
}

func (r *ModifyDBInstanceConnectionStringRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceConnectionStringRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(value string) {
	r.CurrentConnectionString = value
	r.QueryParams.Set("CurrentConnectionString", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetCurrentConnectionString() string {
	return r.CurrentConnectionString
}
func (r *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(value string) {
	r.ConnectionStringPrefix = value
	r.QueryParams.Set("ConnectionStringPrefix", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetConnectionStringPrefix() string {
	return r.ConnectionStringPrefix
}
func (r *ModifyDBInstanceConnectionStringRequest) SetPort(value string) {
	r.Port = value
	r.QueryParams.Set("Port", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetPort() string {
	return r.Port
}
func (r *ModifyDBInstanceConnectionStringRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceConnectionStringRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceConnectionStringRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceConnectionString")
	r.SetProduct(Product)
}

type ModifyDBInstanceConnectionStringResponse struct {
}

func ModifyDBInstanceConnectionString(req *ModifyDBInstanceConnectionStringRequest, accessId, accessSecret string) (*ModifyDBInstanceConnectionStringResponse, error) {
	var pResponse ModifyDBInstanceConnectionStringResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
