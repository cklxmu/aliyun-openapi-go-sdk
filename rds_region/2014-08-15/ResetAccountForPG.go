package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ResetAccountForPGRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	AccountPassword      string
	OwnerAccount         string
}

func (r *ResetAccountForPGRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ResetAccountForPGRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ResetAccountForPGRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ResetAccountForPGRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ResetAccountForPGRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ResetAccountForPGRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ResetAccountForPGRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ResetAccountForPGRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ResetAccountForPGRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *ResetAccountForPGRequest) GetAccountName() string {
	return r.AccountName
}
func (r *ResetAccountForPGRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *ResetAccountForPGRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *ResetAccountForPGRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ResetAccountForPGRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ResetAccountForPGRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ResetAccountForPG")
	r.SetProduct(Product)
}

type ResetAccountForPGResponse struct {
}

func ResetAccountForPG(req *ResetAccountForPGRequest, accessId, accessSecret string) (*ResetAccountForPGResponse, error) {
	var pResponse ResetAccountForPGResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
