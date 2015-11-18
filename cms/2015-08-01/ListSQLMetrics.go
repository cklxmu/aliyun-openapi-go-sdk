package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListSQLMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
	Page        int
	PageSize    int
}

func (r *ListSQLMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *ListSQLMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListSQLMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListSQLMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListSQLMetricsRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListSQLMetricsRequest) GetPage() int {
	return r.Page
}
func (r *ListSQLMetricsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListSQLMetricsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListSQLMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListSQLMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListSQLMetricsResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListSQLMetrics(req *ListSQLMetricsRequest, accessId, accessSecret string) (*ListSQLMetricsResponse, error) {
	var pResponse ListSQLMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
