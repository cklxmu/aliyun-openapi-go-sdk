package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type PutDimTableDataRequest struct {
	RpcRequest
	DimTableName string
	Body         string
}

func (r *PutDimTableDataRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *PutDimTableDataRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *PutDimTableDataRequest) SetBody(value string) {
	r.Body = value
	r.QueryParams.Set("Body", value)
}
func (r *PutDimTableDataRequest) GetBody() string {
	return r.Body
}

func (r *PutDimTableDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutDimTableData")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type PutDimTableDataResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func PutDimTableData(req *PutDimTableDataRequest, accessId, accessSecret string) (*PutDimTableDataResponse, error) {
	var pResponse PutDimTableDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
