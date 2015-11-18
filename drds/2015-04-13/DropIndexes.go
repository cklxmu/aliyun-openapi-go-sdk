package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DropIndexesRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	Table          string
	Indexes        string
}

func (r *DropIndexesRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DropIndexesRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DropIndexesRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DropIndexesRequest) GetDbName() string {
	return r.DbName
}
func (r *DropIndexesRequest) SetTable(value string) {
	r.Table = value
	r.QueryParams.Set("Table", value)
}
func (r *DropIndexesRequest) GetTable() string {
	return r.Table
}
func (r *DropIndexesRequest) SetIndexes(value string) {
	r.Indexes = value
	r.QueryParams.Set("Indexes", value)
}
func (r *DropIndexesRequest) GetIndexes() string {
	return r.Indexes
}

func (r *DropIndexesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DropIndexes")
	r.SetProduct(Product)
}

type DropIndexesResponse struct {
	TaskId string `xml:"TaskId" json:"TaskId"`
}

func DropIndexes(req *DropIndexesRequest, accessId, accessSecret string) (*DropIndexesResponse, error) {
	var pResponse DropIndexesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
