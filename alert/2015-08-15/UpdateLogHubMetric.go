package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateLogHubMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
	Metric      string
}

func (r *UpdateLogHubMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateLogHubMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateLogHubMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *UpdateLogHubMetricRequest) GetMetricName() string {
	return r.MetricName
}
func (r *UpdateLogHubMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *UpdateLogHubMetricRequest) GetMetric() string {
	return r.Metric
}

func (r *UpdateLogHubMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/logHubMetrics/MetricName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateLogHubMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateLogHubMetric(req *UpdateLogHubMetricRequest, accessId, accessSecret string) (*UpdateLogHubMetricResponse, error) {
	var pResponse UpdateLogHubMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
