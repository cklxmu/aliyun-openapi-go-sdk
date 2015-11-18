package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeMetricListRequest struct {
	RpcRequest
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
}

func (r *DescribeMetricListRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *DescribeMetricListRequest) GetProject() string {
	return r.Project
}
func (r *DescribeMetricListRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *DescribeMetricListRequest) GetMetric() string {
	return r.Metric
}
func (r *DescribeMetricListRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *DescribeMetricListRequest) GetPeriod() string {
	return r.Period
}
func (r *DescribeMetricListRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeMetricListRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeMetricListRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeMetricListRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeMetricListRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *DescribeMetricListRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *DescribeMetricListRequest) SetPage(value string) {
	r.Page = value
	r.QueryParams.Set("Page", value)
}
func (r *DescribeMetricListRequest) GetPage() string {
	return r.Page
}
func (r *DescribeMetricListRequest) SetLength(value string) {
	r.Length = value
	r.QueryParams.Set("Length", value)
}
func (r *DescribeMetricListRequest) GetLength() string {
	return r.Length
}

func (r *DescribeMetricListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMetricList")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DescribeMetricListResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func DescribeMetricList(req *DescribeMetricListRequest, accessId, accessSecret string) (*DescribeMetricListResponse, error) {
	var pResponse DescribeMetricListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
