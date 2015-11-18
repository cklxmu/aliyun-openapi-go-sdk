package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteLogHubMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
}

func (r *DeleteLogHubMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteLogHubMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteLogHubMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *DeleteLogHubMetricRequest) GetMetricName() string {
	return r.MetricName
}

func (r *DeleteLogHubMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/logHubMetrics/MetricName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteLogHubMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteLogHubMetric(req *DeleteLogHubMetricRequest, accessId, accessSecret string) (*DeleteLogHubMetricResponse, error) {
	var pResponse DeleteLogHubMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
