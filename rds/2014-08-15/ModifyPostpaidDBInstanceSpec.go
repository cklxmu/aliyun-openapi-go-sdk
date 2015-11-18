package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyPostpaidDBInstanceSpecRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	DBInstanceClass      string
	DBInstanceStorage    int
	OwnerAccount         string
}

func (r *ModifyPostpaidDBInstanceSpecRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *ModifyPostpaidDBInstanceSpecRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyPostpaidDBInstanceSpecRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyPostpaidDBInstanceSpecRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyPostpaidDBInstanceSpec")
	r.SetProduct(Product)
}

type ModifyPostpaidDBInstanceSpecResponse struct {
}

func ModifyPostpaidDBInstanceSpec(req *ModifyPostpaidDBInstanceSpecRequest, accessId, accessSecret string) (*ModifyPostpaidDBInstanceSpecResponse, error) {
	var pResponse ModifyPostpaidDBInstanceSpecResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
