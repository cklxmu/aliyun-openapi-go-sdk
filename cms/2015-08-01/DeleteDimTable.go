package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDimTableRequest struct {
	RpcRequest
	DimTableName string
}

func (r *DeleteDimTableRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *DeleteDimTableRequest) GetDimTableName() string {
	return r.DimTableName
}

func (r *DeleteDimTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDimTable")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteDimTableResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func DeleteDimTable(req *DeleteDimTableRequest, accessId, accessSecret string) (*DeleteDimTableResponse, error) {
	var pResponse DeleteDimTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
