package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateDimTableRequest struct {
	RpcRequest
	DimTable string
}

func (r *CreateDimTableRequest) SetDimTable(value string) {
	r.DimTable = value
	r.QueryParams.Set("DimTable", value)
}
func (r *CreateDimTableRequest) GetDimTable() string {
	return r.DimTable
}

func (r *CreateDimTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDimTable")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateDimTableResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  int    `xml:"Result" json:"Result"`
}

func CreateDimTable(req *CreateDimTableRequest, accessId, accessSecret string) (*CreateDimTableResponse, error) {
	var pResponse CreateDimTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
