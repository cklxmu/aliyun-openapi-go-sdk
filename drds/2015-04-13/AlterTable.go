package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type AlterTableRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	DdlSql         string
}

func (r *AlterTableRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *AlterTableRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *AlterTableRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *AlterTableRequest) GetDbName() string {
	return r.DbName
}
func (r *AlterTableRequest) SetDdlSql(value string) {
	r.DdlSql = value
	r.QueryParams.Set("DdlSql", value)
}
func (r *AlterTableRequest) GetDdlSql() string {
	return r.DdlSql
}

func (r *AlterTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AlterTable")
	r.SetProduct(Product)
}

type AlterTableResponse struct {
	TaskId string `xml:"TaskId" json:"TaskId"`
}

func AlterTable(req *AlterTableRequest, accessId, accessSecret string) (*AlterTableResponse, error) {
	var pResponse AlterTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
