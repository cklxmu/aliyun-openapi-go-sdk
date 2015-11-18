package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UnlockDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	OwnerAccount         string
}

func (r *UnlockDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *UnlockDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *UnlockDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *UnlockDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *UnlockDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *UnlockDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *UnlockDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *UnlockDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *UnlockDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *UnlockDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *UnlockDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UnlockDBInstance")
	r.SetProduct(Product)
}

type UnlockDBInstanceResponse struct {
}

func UnlockDBInstance(req *UnlockDBInstanceRequest, accessId, accessSecret string) (*UnlockDBInstanceResponse, error) {
	var pResponse UnlockDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
