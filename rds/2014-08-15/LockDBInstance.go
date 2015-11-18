package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type LockDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	LockReason           string
	OwnerAccount         string
}

func (r *LockDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *LockDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *LockDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *LockDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *LockDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *LockDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *LockDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *LockDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *LockDBInstanceRequest) SetLockReason(value string) {
	r.LockReason = value
	r.QueryParams.Set("LockReason", value)
}
func (r *LockDBInstanceRequest) GetLockReason() string {
	return r.LockReason
}
func (r *LockDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *LockDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *LockDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("LockDBInstance")
	r.SetProduct(Product)
}

type LockDBInstanceResponse struct {
}

func LockDBInstance(req *LockDBInstanceRequest, accessId, accessSecret string) (*LockDBInstanceResponse, error) {
	var pResponse LockDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
