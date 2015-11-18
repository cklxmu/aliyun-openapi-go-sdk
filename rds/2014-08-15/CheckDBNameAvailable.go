package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CheckDBNameAvailableRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	DBName               string
	OwnerAccount         string
}

func (r *CheckDBNameAvailableRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CheckDBNameAvailableRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CheckDBNameAvailableRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CheckDBNameAvailableRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CheckDBNameAvailableRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CheckDBNameAvailableRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CheckDBNameAvailableRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CheckDBNameAvailableRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CheckDBNameAvailableRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CheckDBNameAvailableRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CheckDBNameAvailableRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *CheckDBNameAvailableRequest) GetDBName() string {
	return r.DBName
}
func (r *CheckDBNameAvailableRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CheckDBNameAvailableRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CheckDBNameAvailableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CheckDBNameAvailable")
	r.SetProduct(Product)
}

type CheckDBNameAvailableResponse struct {
}

func CheckDBNameAvailable(req *CheckDBNameAvailableRequest, accessId, accessSecret string) (*CheckDBNameAvailableResponse, error) {
	var pResponse CheckDBNameAvailableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
