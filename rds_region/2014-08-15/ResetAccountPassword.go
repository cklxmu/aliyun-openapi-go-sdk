package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ResetAccountPasswordRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	AccountPassword      string
	OwnerAccount         string
}

func (r *ResetAccountPasswordRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ResetAccountPasswordRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ResetAccountPasswordRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ResetAccountPasswordRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ResetAccountPasswordRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ResetAccountPasswordRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ResetAccountPasswordRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ResetAccountPasswordRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ResetAccountPasswordRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *ResetAccountPasswordRequest) GetAccountName() string {
	return r.AccountName
}
func (r *ResetAccountPasswordRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *ResetAccountPasswordRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *ResetAccountPasswordRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ResetAccountPasswordRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ResetAccountPasswordRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ResetAccountPassword")
	r.SetProduct(Product)
}

type ResetAccountPasswordResponse struct {
}

func ResetAccountPassword(req *ResetAccountPasswordRequest, accessId, accessSecret string) (*ResetAccountPasswordResponse, error) {
	var pResponse ResetAccountPasswordResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
