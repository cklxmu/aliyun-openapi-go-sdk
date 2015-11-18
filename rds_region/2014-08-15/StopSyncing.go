package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopSyncingRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	ImportId             int
	OwnerAccount         string
}

func (r *StopSyncingRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StopSyncingRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StopSyncingRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StopSyncingRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StopSyncingRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StopSyncingRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StopSyncingRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *StopSyncingRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *StopSyncingRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *StopSyncingRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *StopSyncingRequest) SetImportId(value int) {
	r.ImportId = value
	r.QueryParams.Set("ImportId", strconv.Itoa(value))
}
func (r *StopSyncingRequest) GetImportId() int {
	return r.ImportId
}
func (r *StopSyncingRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StopSyncingRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StopSyncingRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopSyncing")
	r.SetProduct(Product)
}

type StopSyncingResponse struct {
}

func StopSyncing(req *StopSyncingRequest, accessId, accessSecret string) (*StopSyncingResponse, error) {
	var pResponse StopSyncingResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
