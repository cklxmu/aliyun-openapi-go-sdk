package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetLogHubMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
}

func (r *GetLogHubMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetLogHubMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetLogHubMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *GetLogHubMetricRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetLogHubMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/logHubMetrics/MetricName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetLogHubMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetLogHubMetric(req *GetLogHubMetricRequest, accessId, accessSecret string) (*GetLogHubMetricResponse, error) {
	var pResponse GetLogHubMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
