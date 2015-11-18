package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ImportDatabaseBetweenInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	SourceDBInstanceId   string
	DBInfo               string
	OwnerAccount         string
}

func (r *ImportDatabaseBetweenInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ImportDatabaseBetweenInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ImportDatabaseBetweenInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ImportDatabaseBetweenInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ImportDatabaseBetweenInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ImportDatabaseBetweenInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ImportDatabaseBetweenInstancesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ImportDatabaseBetweenInstancesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ImportDatabaseBetweenInstancesRequest) SetSourceDBInstanceId(value string) {
	r.SourceDBInstanceId = value
	r.QueryParams.Set("SourceDBInstanceId", value)
}
func (r *ImportDatabaseBetweenInstancesRequest) GetSourceDBInstanceId() string {
	return r.SourceDBInstanceId
}
func (r *ImportDatabaseBetweenInstancesRequest) SetDBInfo(value string) {
	r.DBInfo = value
	r.QueryParams.Set("DBInfo", value)
}
func (r *ImportDatabaseBetweenInstancesRequest) GetDBInfo() string {
	return r.DBInfo
}
func (r *ImportDatabaseBetweenInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ImportDatabaseBetweenInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ImportDatabaseBetweenInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ImportDatabaseBetweenInstances")
	r.SetProduct(Product)
}

type ImportDatabaseBetweenInstancesResponse struct {
	ImportId string `xml:"ImportId" json:"ImportId"`
}

func ImportDatabaseBetweenInstances(req *ImportDatabaseBetweenInstancesRequest, accessId, accessSecret string) (*ImportDatabaseBetweenInstancesResponse, error) {
	var pResponse ImportDatabaseBetweenInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
