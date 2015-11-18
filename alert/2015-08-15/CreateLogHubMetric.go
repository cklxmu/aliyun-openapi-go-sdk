package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateLogHubMetricRequest struct {
	RoaRequest
	ProjectName string
	Metric      string
}

func (r *CreateLogHubMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateLogHubMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateLogHubMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *CreateLogHubMetricRequest) GetMetric() string {
	return r.Metric
}

func (r *CreateLogHubMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/logHubMetrics"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateLogHubMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateLogHubMetric(req *CreateLogHubMetricRequest, accessId, accessSecret string) (*CreateLogHubMetricResponse, error) {
	var pResponse CreateLogHubMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
