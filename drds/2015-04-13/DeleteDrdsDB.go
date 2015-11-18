package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDrdsDBRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
}

func (r *DeleteDrdsDBRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DeleteDrdsDBRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DeleteDrdsDBRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DeleteDrdsDBRequest) GetDbName() string {
	return r.DbName
}

func (r *DeleteDrdsDBRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDrdsDB")
	r.SetProduct(Product)
}

type DeleteDrdsDBResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func DeleteDrdsDB(req *DeleteDrdsDBRequest, accessId, accessSecret string) (*DeleteDrdsDBResponse, error) {
	var pResponse DeleteDrdsDBResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
