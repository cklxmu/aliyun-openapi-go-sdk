package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type StopMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *StopMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *StopMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *StopMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *StopMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *StopMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopMetricStream")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type StopMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func StopMetricStream(req *StopMetricStreamRequest, accessId, accessSecret string) (*StopMetricStreamResponse, error) {
	var pResponse StopMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
