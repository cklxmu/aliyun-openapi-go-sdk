package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RestoreDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	BackupId             string
	OwnerAccount         string
}

func (r *RestoreDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RestoreDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RestoreDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RestoreDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RestoreDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RestoreDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RestoreDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *RestoreDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *RestoreDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *RestoreDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *RestoreDBInstanceRequest) SetBackupId(value string) {
	r.BackupId = value
	r.QueryParams.Set("BackupId", value)
}
func (r *RestoreDBInstanceRequest) GetBackupId() string {
	return r.BackupId
}
func (r *RestoreDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RestoreDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RestoreDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RestoreDBInstance")
	r.SetProduct(Product)
}

type RestoreDBInstanceResponse struct {
}

func RestoreDBInstance(req *RestoreDBInstanceRequest, accessId, accessSecret string) (*RestoreDBInstanceResponse, error) {
	var pResponse RestoreDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
