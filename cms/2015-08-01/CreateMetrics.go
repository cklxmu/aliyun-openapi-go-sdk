package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateMetricsRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	Metrics          string
}

func (r *CreateMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *CreateMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateMetricsRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *CreateMetricsRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *CreateMetricsRequest) SetMetrics(value string) {
	r.Metrics = value
	r.QueryParams.Set("Metrics", value)
}
func (r *CreateMetricsRequest) GetMetrics() string {
	return r.Metrics
}

func (r *CreateMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func CreateMetrics(req *CreateMetricsRequest, accessId, accessSecret string) (*CreateMetricsResponse, error) {
	var pResponse CreateMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
