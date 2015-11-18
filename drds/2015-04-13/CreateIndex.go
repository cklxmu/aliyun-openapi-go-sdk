package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateIndexRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	DdlSql         string
}

func (r *CreateIndexRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *CreateIndexRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *CreateIndexRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *CreateIndexRequest) GetDbName() string {
	return r.DbName
}
func (r *CreateIndexRequest) SetDdlSql(value string) {
	r.DdlSql = value
	r.QueryParams.Set("DdlSql", value)
}
func (r *CreateIndexRequest) GetDdlSql() string {
	return r.DdlSql
}

func (r *CreateIndexRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateIndex")
	r.SetProduct(Product)
}

type CreateIndexResponse struct {
	TaskId string `xml:"TaskId" json:"TaskId"`
}

func CreateIndex(req *CreateIndexRequest, accessId, accessSecret string) (*CreateIndexResponse, error) {
	var pResponse CreateIndexResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
