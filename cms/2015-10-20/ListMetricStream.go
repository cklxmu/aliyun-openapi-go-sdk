package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	Page             int
	PageSize         int
}

func (r *ListMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *ListMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *ListMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *ListMetricStreamRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListMetricStreamRequest) GetPage() int {
	return r.Page
}
func (r *ListMetricStreamRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListMetricStreamRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListMetricStream")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListMetricStreamResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListMetricStream(req *ListMetricStreamRequest, accessId, accessSecret string) (*ListMetricStreamResponse, error) {
	var pResponse ListMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
