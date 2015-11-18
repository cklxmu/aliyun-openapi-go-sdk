package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListLogHubMetricRequest struct {
	RoaRequest
	ProjectName string
	MetricName  string
	Page        int
	PageSize    int
}

func (r *ListLogHubMetricRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListLogHubMetricRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListLogHubMetricRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListLogHubMetricRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListLogHubMetricRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListLogHubMetricRequest) GetPage() int {
	return r.Page
}
func (r *ListLogHubMetricRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListLogHubMetricRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListLogHubMetricRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/logHubMetrics"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListLogHubMetricResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListLogHubMetric(req *ListLogHubMetricRequest, accessId, accessSecret string) (*ListLogHubMetricResponse, error) {
	var pResponse ListLogHubMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
