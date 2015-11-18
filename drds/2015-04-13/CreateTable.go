package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateTableRequest struct {
	RpcRequest
	DrdsInstanceId     string
	DbName             string
	DdlSql             string
	ShardType          string
	ShardKey           string
	AllowFullTableScan string
}

func (r *CreateTableRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *CreateTableRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *CreateTableRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *CreateTableRequest) GetDbName() string {
	return r.DbName
}
func (r *CreateTableRequest) SetDdlSql(value string) {
	r.DdlSql = value
	r.QueryParams.Set("DdlSql", value)
}
func (r *CreateTableRequest) GetDdlSql() string {
	return r.DdlSql
}
func (r *CreateTableRequest) SetShardType(value string) {
	r.ShardType = value
	r.QueryParams.Set("ShardType", value)
}
func (r *CreateTableRequest) GetShardType() string {
	return r.ShardType
}
func (r *CreateTableRequest) SetShardKey(value string) {
	r.ShardKey = value
	r.QueryParams.Set("ShardKey", value)
}
func (r *CreateTableRequest) GetShardKey() string {
	return r.ShardKey
}
func (r *CreateTableRequest) SetAllowFullTableScan(value string) {
	r.AllowFullTableScan = value
	r.QueryParams.Set("AllowFullTableScan", value)
}
func (r *CreateTableRequest) GetAllowFullTableScan() string {
	return r.AllowFullTableScan
}

func (r *CreateTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateTable")
	r.SetProduct(Product)
}

type CreateTableResponse struct {
	TaskId string `xml:"TaskId" json:"TaskId"`
}

func CreateTable(req *CreateTableRequest, accessId, accessSecret string) (*CreateTableResponse, error) {
	var pResponse CreateTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
