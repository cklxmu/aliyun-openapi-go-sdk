package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type BatchPutDimTableDataRequest struct {
	RpcRequest
	DimTableName string
	Body         string
}

func (r *BatchPutDimTableDataRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *BatchPutDimTableDataRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *BatchPutDimTableDataRequest) SetBody(value string) {
	r.Body = value
	r.QueryParams.Set("Body", value)
}
func (r *BatchPutDimTableDataRequest) GetBody() string {
	return r.Body
}

func (r *BatchPutDimTableDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchPutDimTableData")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type BatchPutDimTableDataResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func BatchPutDimTableData(req *BatchPutDimTableDataRequest, accessId, accessSecret string) (*BatchPutDimTableDataResponse, error) {
	var pResponse BatchPutDimTableDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
