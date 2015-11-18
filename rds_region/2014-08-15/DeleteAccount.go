package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteAccountRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	OwnerAccount         string
}

func (r *DeleteAccountRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteAccountRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteAccountRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteAccountRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteAccountRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteAccountRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteAccountRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DeleteAccountRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DeleteAccountRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *DeleteAccountRequest) GetAccountName() string {
	return r.AccountName
}
func (r *DeleteAccountRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteAccountRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteAccount")
	r.SetProduct(Product)
}

type DeleteAccountResponse struct {
}

func DeleteAccount(req *DeleteAccountRequest, accessId, accessSecret string) (*DeleteAccountResponse, error) {
	var pResponse DeleteAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
