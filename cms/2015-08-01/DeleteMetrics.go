package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
}

func (r *DeleteMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *DeleteMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *DeleteMetricsRequest) GetMetricName() string {
	return r.MetricName
}

func (r *DeleteMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteMetrics")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func DeleteMetrics(req *DeleteMetricsRequest, accessId, accessSecret string) (*DeleteMetricsResponse, error) {
	var pResponse DeleteMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
