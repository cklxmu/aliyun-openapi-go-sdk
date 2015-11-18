package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyAccountDescriptionRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	AccountDescription   string
	OwnerAccount         string
}

func (r *ModifyAccountDescriptionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyAccountDescriptionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyAccountDescriptionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyAccountDescriptionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyAccountDescriptionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyAccountDescriptionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyAccountDescriptionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyAccountDescriptionRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *ModifyAccountDescriptionRequest) GetAccountName() string {
	return r.AccountName
}
func (r *ModifyAccountDescriptionRequest) SetAccountDescription(value string) {
	r.AccountDescription = value
	r.QueryParams.Set("AccountDescription", value)
}
func (r *ModifyAccountDescriptionRequest) GetAccountDescription() string {
	return r.AccountDescription
}
func (r *ModifyAccountDescriptionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyAccountDescriptionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyAccountDescriptionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyAccountDescription")
	r.SetProduct(Product)
}

type ModifyAccountDescriptionResponse struct {
}

func ModifyAccountDescription(req *ModifyAccountDescriptionRequest, accessId, accessSecret string) (*ModifyAccountDescriptionResponse, error) {
	var pResponse ModifyAccountDescriptionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
