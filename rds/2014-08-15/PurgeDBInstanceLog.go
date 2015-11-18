package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PurgeDBInstanceLogRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *PurgeDBInstanceLogRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *PurgeDBInstanceLogRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *PurgeDBInstanceLogRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *PurgeDBInstanceLogRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *PurgeDBInstanceLogRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *PurgeDBInstanceLogRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *PurgeDBInstanceLogRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *PurgeDBInstanceLogRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *PurgeDBInstanceLogRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *PurgeDBInstanceLogRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *PurgeDBInstanceLogRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *PurgeDBInstanceLogRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *PurgeDBInstanceLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PurgeDBInstanceLog")
	r.SetProduct(Product)
}

type PurgeDBInstanceLogResponse struct {
}

func PurgeDBInstanceLog(req *PurgeDBInstanceLogRequest, accessId, accessSecret string) (*PurgeDBInstanceLogResponse, error) {
	var pResponse PurgeDBInstanceLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
