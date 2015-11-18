package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateTempDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	BackupId             int
	RestoreTime          string
	OwnerAccount         string
}

func (r *CreateTempDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateTempDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateTempDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateTempDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateTempDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateTempDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateTempDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateTempDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateTempDBInstanceRequest) SetBackupId(value int) {
	r.BackupId = value
	r.QueryParams.Set("BackupId", strconv.Itoa(value))
}
func (r *CreateTempDBInstanceRequest) GetBackupId() int {
	return r.BackupId
}
func (r *CreateTempDBInstanceRequest) SetRestoreTime(value string) {
	r.RestoreTime = value
	r.QueryParams.Set("RestoreTime", value)
}
func (r *CreateTempDBInstanceRequest) GetRestoreTime() string {
	return r.RestoreTime
}
func (r *CreateTempDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateTempDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateTempDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateTempDBInstance")
	r.SetProduct(Product)
}

type CreateTempDBInstanceResponse struct {
	TempDBInstanceId string `xml:"TempDBInstanceId" json:"TempDBInstanceId"`
}

func CreateTempDBInstance(req *CreateTempDBInstanceRequest, accessId, accessSecret string) (*CreateTempDBInstanceResponse, error) {
	var pResponse CreateTempDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
