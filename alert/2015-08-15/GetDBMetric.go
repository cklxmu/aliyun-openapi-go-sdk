package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetDBMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
}

func (r *GetDBMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetDBMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetDBMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.PathParams.Set("MetricName", value)
}
func (r *GetDBMetricRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetDBMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/dbMetrics/MetricName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetDBMetricResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetDBMetric(req *GetDBMetricRequest, accessId, accessSecret string) (*GetDBMetricResponse, error) {
	var pResponse GetDBMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
