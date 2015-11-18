package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeMetricRequest struct {
	RpcRequest
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
}

func (r *DescribeMetricRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *DescribeMetricRequest) GetProject() string {
	return r.Project
}
func (r *DescribeMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *DescribeMetricRequest) GetMetric() string {
	return r.Metric
}
func (r *DescribeMetricRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *DescribeMetricRequest) GetPeriod() string {
	return r.Period
}
func (r *DescribeMetricRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeMetricRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeMetricRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeMetricRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeMetricRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *DescribeMetricRequest) GetDimensions() string {
	return r.Dimensions
}

func (r *DescribeMetricRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMetric")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DescribeMetricResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func DescribeMetric(req *DescribeMetricRequest, accessId, accessSecret string) (*DescribeMetricResponse, error) {
	var pResponse DescribeMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
