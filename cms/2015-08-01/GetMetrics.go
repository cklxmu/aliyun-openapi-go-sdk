package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetMetricsRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	MetricName       string
}

func (r *GetMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetMetricsRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *GetMetricsRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *GetMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *GetMetricsRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetMetrics(req *GetMetricsRequest, accessId, accessSecret string) (*GetMetricsResponse, error) {
	var pResponse GetMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
