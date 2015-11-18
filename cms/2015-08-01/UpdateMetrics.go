package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateMetricsRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	MetricName       string
	Metrics          string
}

func (r *UpdateMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *UpdateMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateMetricsRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *UpdateMetricsRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *UpdateMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *UpdateMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *UpdateMetricsRequest) SetMetrics(value string) {
	r.Metrics = value
	r.QueryParams.Set("Metrics", value)
}
func (r *UpdateMetricsRequest) GetMetrics() string {
	return r.Metrics
}

func (r *UpdateMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func UpdateMetrics(req *UpdateMetricsRequest, accessId, accessSecret string) (*UpdateMetricsResponse, error) {
	var pResponse UpdateMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
