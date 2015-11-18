package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type MigrateToOtherZoneRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	ZoneId               string
	OwnerAccount         string
}

func (r *MigrateToOtherZoneRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *MigrateToOtherZoneRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *MigrateToOtherZoneRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *MigrateToOtherZoneRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *MigrateToOtherZoneRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *MigrateToOtherZoneRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *MigrateToOtherZoneRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *MigrateToOtherZoneRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *MigrateToOtherZoneRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *MigrateToOtherZoneRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *MigrateToOtherZoneRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *MigrateToOtherZoneRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *MigrateToOtherZoneRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("MigrateToOtherZone")
	r.SetProduct(Product)
}

type MigrateToOtherZoneResponse struct {
}

func MigrateToOtherZone(req *MigrateToOtherZoneRequest, accessId, accessSecret string) (*MigrateToOtherZoneResponse, error) {
	var pResponse MigrateToOtherZoneResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
