package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifySecurityIpsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	SecurityIps          string
	OwnerAccount         string
}

func (r *ModifySecurityIpsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityIpsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifySecurityIpsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifySecurityIpsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifySecurityIpsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityIpsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifySecurityIpsRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifySecurityIpsRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifySecurityIpsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifySecurityIpsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifySecurityIpsRequest) SetSecurityIps(value string) {
	r.SecurityIps = value
	r.QueryParams.Set("SecurityIps", value)
}
func (r *ModifySecurityIpsRequest) GetSecurityIps() string {
	return r.SecurityIps
}
func (r *ModifySecurityIpsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifySecurityIpsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifySecurityIpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifySecurityIps")
	r.SetProduct(Product)
}

type ModifySecurityIpsResponse struct {
}

func ModifySecurityIps(req *ModifySecurityIpsRequest, accessId, accessSecret string) (*ModifySecurityIpsResponse, error) {
	var pResponse ModifySecurityIpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
