package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeMetricDatumRequest struct {
	RpcRequest
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Page       string
	Length     string
	Extend     string
}

func (r *DescribeMetricDatumRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *DescribeMetricDatumRequest) GetProject() string {
	return r.Project
}
func (r *DescribeMetricDatumRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *DescribeMetricDatumRequest) GetMetric() string {
	return r.Metric
}
func (r *DescribeMetricDatumRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *DescribeMetricDatumRequest) GetPeriod() string {
	return r.Period
}
func (r *DescribeMetricDatumRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeMetricDatumRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeMetricDatumRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeMetricDatumRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeMetricDatumRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *DescribeMetricDatumRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *DescribeMetricDatumRequest) SetPage(value string) {
	r.Page = value
	r.QueryParams.Set("Page", value)
}
func (r *DescribeMetricDatumRequest) GetPage() string {
	return r.Page
}
func (r *DescribeMetricDatumRequest) SetLength(value string) {
	r.Length = value
	r.QueryParams.Set("Length", value)
}
func (r *DescribeMetricDatumRequest) GetLength() string {
	return r.Length
}
func (r *DescribeMetricDatumRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *DescribeMetricDatumRequest) GetExtend() string {
	return r.Extend
}

func (r *DescribeMetricDatumRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMetricDatum")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DescribeMetricDatumResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func DescribeMetricDatum(req *DescribeMetricDatumRequest, accessId, accessSecret string) (*DescribeMetricDatumResponse, error) {
	var pResponse DescribeMetricDatumResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
