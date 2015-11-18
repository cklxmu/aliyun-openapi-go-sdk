package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateMetricStreamRequest struct {
	RpcRequest
	ProjectName  string
	MetricStream string
}

func (r *CreateMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *CreateMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateMetricStreamRequest) SetMetricStream(value string) {
	r.MetricStream = value
	r.QueryParams.Set("MetricStream", value)
}
func (r *CreateMetricStreamRequest) GetMetricStream() string {
	return r.MetricStream
}

func (r *CreateMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateMetricStream")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func CreateMetricStream(req *CreateMetricStreamRequest, accessId, accessSecret string) (*CreateMetricStreamResponse, error) {
	var pResponse CreateMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
