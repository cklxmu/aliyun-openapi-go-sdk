package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DeleteDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DeleteDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DeleteDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DeleteDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DeleteDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDBInstance")
	r.SetProduct(Product)
}

type DeleteDBInstanceResponse struct {
}

func DeleteDBInstance(req *DeleteDBInstanceRequest, accessId, accessSecret string) (*DeleteDBInstanceResponse, error) {
	var pResponse DeleteDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
