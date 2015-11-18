package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDBMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
}

func (r *DeleteDBMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteDBMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteDBMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *DeleteDBMetricRequest) GetMetricName() string {
	return r.MetricName
}

func (r *DeleteDBMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/dbMetrics/MetricName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteDBMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteDBMetric(req *DeleteDBMetricRequest, accessId, accessSecret string) (*DeleteDBMetricResponse, error) {
	var pResponse DeleteDBMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
