package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListMetricsRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	MetricName       string
	Page             int
	PageSize         int
}

func (r *ListMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *ListMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListMetricsRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *ListMetricsRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *ListMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListMetricsRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListMetricsRequest) GetPage() int {
	return r.Page
}
func (r *ListMetricsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListMetricsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListMetricsResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListMetrics(req *ListMetricsRequest, accessId, accessSecret string) (*ListMetricsResponse, error) {
	var pResponse ListMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
