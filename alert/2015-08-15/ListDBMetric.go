package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListDBMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
	Page        int
	PageSize    int
}

func (r *ListDBMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListDBMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListDBMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListDBMetricRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListDBMetricRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListDBMetricRequest) GetPage() int {
	return r.Page
}
func (r *ListDBMetricRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListDBMetricRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListDBMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/dbMetrics"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListDBMetricResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListDBMetric(req *ListDBMetricRequest, accessId, accessSecret string) (*ListDBMetricResponse, error) {
	var pResponse ListDBMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
