package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDimTableDataRequest struct {
	RpcRequest
	DimTableName string
	Key          string
}

func (r *DeleteDimTableDataRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *DeleteDimTableDataRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *DeleteDimTableDataRequest) SetKey(value string) {
	r.Key = value
	r.QueryParams.Set("Key", value)
}
func (r *DeleteDimTableDataRequest) GetKey() string {
	return r.Key
}

func (r *DeleteDimTableDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDimTableData")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteDimTableDataResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func DeleteDimTableData(req *DeleteDimTableDataRequest, accessId, accessSecret string) (*DeleteDimTableDataResponse, error) {
	var pResponse DeleteDimTableDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
