package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateDBMetricRequest struct {
	RoaRequest
	ProjectName string
	Metric      string
}

func (r *CreateDBMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateDBMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateDBMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *CreateDBMetricRequest) GetMetric() string {
	return r.Metric
}

func (r *CreateDBMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/dbMetrics"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateDBMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateDBMetric(req *CreateDBMetricRequest, accessId, accessSecret string) (*CreateDBMetricResponse, error) {
	var pResponse CreateDBMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
