package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListDimTableDataRequest struct {
	RpcRequest
	DimTableName string
	Key          string
}

func (r *ListDimTableDataRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *ListDimTableDataRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *ListDimTableDataRequest) SetKey(value string) {
	r.Key = value
	r.QueryParams.Set("Key", value)
}
func (r *ListDimTableDataRequest) GetKey() string {
	return r.Key
}

func (r *ListDimTableDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListDimTableData")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListDimTableDataResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func ListDimTableData(req *ListDimTableDataRequest, accessId, accessSecret string) (*ListDimTableDataResponse, error) {
	var pResponse ListDimTableDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
