package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceDescriptionRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ClientToken           string
	DBInstanceId          string
	DBInstanceDescription string
	OwnerAccount          string
}

func (r *ModifyDBInstanceDescriptionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceDescriptionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceDescriptionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceDescriptionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceDescriptionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceDescriptionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceDescriptionRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyDBInstanceDescriptionRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceDescriptionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *ModifyDBInstanceDescriptionRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *ModifyDBInstanceDescriptionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceDescriptionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceDescriptionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceDescription")
	r.SetProduct(Product)
}

type ModifyDBInstanceDescriptionResponse struct {
}

func ModifyDBInstanceDescription(req *ModifyDBInstanceDescriptionRequest, accessId, accessSecret string) (*ModifyDBInstanceDescriptionResponse, error) {
	var pResponse ModifyDBInstanceDescriptionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
