package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBInstanceMaintainTimeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	MaintainTime         string
	OwnerAccount         string
}

func (r *ModifyDBInstanceMaintainTimeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetMaintainTime(value string) {
	r.MaintainTime = value
	r.QueryParams.Set("MaintainTime", value)
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetMaintainTime() string {
	return r.MaintainTime
}
func (r *ModifyDBInstanceMaintainTimeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBInstanceMaintainTimeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBInstanceMaintainTimeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBInstanceMaintainTime")
	r.SetProduct(Product)
}

type ModifyDBInstanceMaintainTimeResponse struct {
}

func ModifyDBInstanceMaintainTime(req *ModifyDBInstanceMaintainTimeRequest, accessId, accessSecret string) (*ModifyDBInstanceMaintainTimeResponse, error) {
	var pResponse ModifyDBInstanceMaintainTimeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
