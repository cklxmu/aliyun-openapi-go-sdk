package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListMetricsPlanRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	MetricName       string
	Page             int
	PageSize         int
}

func (r *ListMetricsPlanRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *ListMetricsPlanRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListMetricsPlanRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *ListMetricsPlanRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *ListMetricsPlanRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *ListMetricsPlanRequest) GetMetricName() string {
	return r.MetricName
}
func (r *ListMetricsPlanRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListMetricsPlanRequest) GetPage() int {
	return r.Page
}
func (r *ListMetricsPlanRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListMetricsPlanRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListMetricsPlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListMetricsPlan")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListMetricsPlanResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListMetricsPlan(req *ListMetricsPlanRequest, accessId, accessSecret string) (*ListMetricsPlanResponse, error) {
	var pResponse ListMetricsPlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
