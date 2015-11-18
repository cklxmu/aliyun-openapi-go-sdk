package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceSpecRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	DBInstanceClass      string
	DBInstanceStorage    int
	PayType              string
	OwnerAccount         string
}

func (r *ModifyDBInstanceSpecRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceSpecRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceSpecRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceSpecRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceSpecRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceSpecRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceSpecRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyDBInstanceSpecRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyDBInstanceSpecRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceSpecRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceSpecRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *ModifyDBInstanceSpecRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *ModifyDBInstanceSpecRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *ModifyDBInstanceSpecRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *ModifyDBInstanceSpecRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *ModifyDBInstanceSpecRequest) GetPayType() string {
	return r.PayType
}
func (r *ModifyDBInstanceSpecRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceSpecRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceSpecRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceSpec")
	r.SetProduct(Product)
}

type ModifyDBInstanceSpecResponse struct {
}

func ModifyDBInstanceSpec(req *ModifyDBInstanceSpecRequest, accessId, accessSecret string) (*ModifyDBInstanceSpecResponse, error) {
	var pResponse ModifyDBInstanceSpecResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
