package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListSqlMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
	Page        int
	PageSize    int
}

func (r *ListSqlMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *ListSqlMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListSqlMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListSqlMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListSqlMetricsRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListSqlMetricsRequest) GetPage() int {
	return r.Page
}
func (r *ListSqlMetricsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListSqlMetricsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListSqlMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListSqlMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListSqlMetricsResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListSqlMetrics(req *ListSqlMetricsRequest, accessId, accessSecret string) (*ListSqlMetricsResponse, error) {
	var pResponse ListSqlMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
