package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetDimTableRequest struct {
	RpcRequest
	DimTableName string
}

func (r *GetDimTableRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *GetDimTableRequest) GetDimTableName() string {
	return r.DimTableName
}

func (r *GetDimTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetDimTable")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetDimTableResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetDimTable(req *GetDimTableRequest, accessId, accessSecret string) (*GetDimTableResponse, error) {
	var pResponse GetDimTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
