package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyBackupPolicyRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	DBInstanceId          string
	PreferredBackupTime   string
	PreferredBackupPeriod string
	OwnerAccount          string
}

func (r *ModifyBackupPolicyRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyBackupPolicyRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyBackupPolicyRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyBackupPolicyRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyBackupPolicyRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyBackupPolicyRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyBackupPolicyRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyBackupPolicyRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyBackupPolicyRequest) SetPreferredBackupTime(value string) {
	r.PreferredBackupTime = value
	r.QueryParams.Set("PreferredBackupTime", value)
}
func (r *ModifyBackupPolicyRequest) GetPreferredBackupTime() string {
	return r.PreferredBackupTime
}
func (r *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(value string) {
	r.PreferredBackupPeriod = value
	r.QueryParams.Set("PreferredBackupPeriod", value)
}
func (r *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() string {
	return r.PreferredBackupPeriod
}
func (r *ModifyBackupPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyBackupPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyBackupPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyBackupPolicy")
	r.SetProduct(Product)
}

type ModifyBackupPolicyResponse struct {
}

func ModifyBackupPolicy(req *ModifyBackupPolicyRequest, accessId, accessSecret string) (*ModifyBackupPolicyResponse, error) {
	var pResponse ModifyBackupPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
