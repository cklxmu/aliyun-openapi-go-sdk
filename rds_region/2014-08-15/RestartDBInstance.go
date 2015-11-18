package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RestartDBInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *RestartDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RestartDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RestartDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RestartDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RestartDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RestartDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RestartDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *RestartDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *RestartDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *RestartDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *RestartDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RestartDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RestartDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RestartDBInstance")
	r.SetProduct(Product)
}

type RestartDBInstanceResponse struct {
}

func RestartDBInstance(req *RestartDBInstanceRequest, accessId, accessSecret string) (*RestartDBInstanceResponse, error) {
	var pResponse RestartDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
