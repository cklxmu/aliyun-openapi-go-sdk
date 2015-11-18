package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type StartMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *StartMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *StartMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *StartMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *StartMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *StartMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartMetricStream")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type StartMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func StartMetricStream(req *StartMetricStreamRequest, accessId, accessSecret string) (*StartMetricStreamResponse, error) {
	var pResponse StartMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
