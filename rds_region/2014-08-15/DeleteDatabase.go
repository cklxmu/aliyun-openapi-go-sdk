package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteDatabaseRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	OwnerAccount         string
}

func (r *DeleteDatabaseRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteDatabaseRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteDatabaseRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteDatabaseRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteDatabaseRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteDatabaseRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteDatabaseRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DeleteDatabaseRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DeleteDatabaseRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *DeleteDatabaseRequest) GetDBName() string {
	return r.DBName
}
func (r *DeleteDatabaseRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteDatabaseRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteDatabaseRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDatabase")
	r.SetProduct(Product)
}

type DeleteDatabaseResponse struct {
}

func DeleteDatabase(req *DeleteDatabaseRequest, accessId, accessSecret string) (*DeleteDatabaseResponse, error) {
	var pResponse DeleteDatabaseResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
