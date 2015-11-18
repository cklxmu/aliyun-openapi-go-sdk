package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateDBMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
	Metric      string
}

func (r *UpdateDBMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateDBMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateDBMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *UpdateDBMetricRequest) GetMetricName() string {
	return r.MetricName
}
func (r *UpdateDBMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *UpdateDBMetricRequest) GetMetric() string {
	return r.Metric
}

func (r *UpdateDBMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/dbMetrics/MetricName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateDBMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateDBMetric(req *UpdateDBMetricRequest, accessId, accessSecret string) (*UpdateDBMetricResponse, error) {
	var pResponse UpdateDBMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
