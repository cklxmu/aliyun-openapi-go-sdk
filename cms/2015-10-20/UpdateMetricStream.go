package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	MetricStream     string
}

func (r *UpdateMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *UpdateMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *UpdateMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *UpdateMetricStreamRequest) SetMetricStream(value string) {
	r.MetricStream = value
	r.QueryParams.Set("MetricStream", value)
}
func (r *UpdateMetricStreamRequest) GetMetricStream() string {
	return r.MetricStream
}

func (r *UpdateMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateMetricStream")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func UpdateMetricStream(req *UpdateMetricStreamRequest, accessId, accessSecret string) (*UpdateMetricStreamResponse, error) {
	var pResponse UpdateMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
