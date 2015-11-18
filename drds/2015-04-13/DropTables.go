package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DropTablesRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	Tables         string
}

func (r *DropTablesRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DropTablesRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DropTablesRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DropTablesRequest) GetDbName() string {
	return r.DbName
}
func (r *DropTablesRequest) SetTables(value string) {
	r.Tables = value
	r.QueryParams.Set("Tables", value)
}
func (r *DropTablesRequest) GetTables() string {
	return r.Tables
}

func (r *DropTablesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DropTables")
	r.SetProduct(Product)
}

type DropTablesResponse struct {
	TaskId string `xml:"TaskId" json:"TaskId"`
}

func DropTables(req *DropTablesRequest, accessId, accessSecret string) (*DropTablesResponse, error) {
	var pResponse DropTablesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
