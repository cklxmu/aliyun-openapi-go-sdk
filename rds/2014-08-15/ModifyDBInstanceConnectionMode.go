package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceConnectionModeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	ConnectionMode       string
	OwnerAccount         string
}

func (r *ModifyDBInstanceConnectionModeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceConnectionModeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceConnectionModeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceConnectionModeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceConnectionModeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceConnectionModeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceConnectionModeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceConnectionModeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceConnectionModeRequest) SetConnectionMode(value string) {
	r.ConnectionMode = value
	r.QueryParams.Set("ConnectionMode", value)
}
func (r *ModifyDBInstanceConnectionModeRequest) GetConnectionMode() string {
	return r.ConnectionMode
}
func (r *ModifyDBInstanceConnectionModeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceConnectionModeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceConnectionModeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceConnectionMode")
	r.SetProduct(Product)
}

type ModifyDBInstanceConnectionModeResponse struct {
}

func ModifyDBInstanceConnectionMode(req *ModifyDBInstanceConnectionModeRequest, accessId, accessSecret string) (*ModifyDBInstanceConnectionModeResponse, error) {
	var pResponse ModifyDBInstanceConnectionModeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
