package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateDimTableRequest struct {
	RpcRequest
	DimTableName string
	DimTable     string
}

func (r *UpdateDimTableRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *UpdateDimTableRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *UpdateDimTableRequest) SetDimTable(value string) {
	r.DimTable = value
	r.QueryParams.Set("DimTable", value)
}
func (r *UpdateDimTableRequest) GetDimTable() string {
	return r.DimTable
}

func (r *UpdateDimTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateDimTable")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateDimTableResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func UpdateDimTable(req *UpdateDimTableRequest, accessId, accessSecret string) (*UpdateDimTableResponse, error) {
	var pResponse UpdateDimTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
