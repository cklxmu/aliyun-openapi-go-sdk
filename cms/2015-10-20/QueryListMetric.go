package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryListMetricRequest struct {
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
	Filter     string
}

func (r *QueryListMetricRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *QueryListMetricRequest) GetProject() string {
	return r.Project
}
func (r *QueryListMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *QueryListMetricRequest) GetMetric() string {
	return r.Metric
}
func (r *QueryListMetricRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *QueryListMetricRequest) GetPeriod() string {
	return r.Period
}
func (r *QueryListMetricRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *QueryListMetricRequest) GetStartTime() string {
	return r.StartTime
}
func (r *QueryListMetricRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *QueryListMetricRequest) GetEndTime() string {
	return r.EndTime
}
func (r *QueryListMetricRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *QueryListMetricRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *QueryListMetricRequest) SetPage(value string) {
	r.Page = value
	r.QueryParams.Set("Page", value)
}
func (r *QueryListMetricRequest) GetPage() string {
	return r.Page
}
func (r *QueryListMetricRequest) SetLength(value string) {
	r.Length = value
	r.QueryParams.Set("Length", value)
}
func (r *QueryListMetricRequest) GetLength() string {
	return r.Length
}
func (r *QueryListMetricRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryListMetricRequest) GetExtend() string {
	return r.Extend
}
func (r *QueryListMetricRequest) SetFilter(value string) {
	r.Filter = value
	r.QueryParams.Set("Filter", value)
}
func (r *QueryListMetricRequest) GetFilter() string {
	return r.Filter
}

func (r *QueryListMetricRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryListMetric")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type QueryListMetricResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func QueryListMetric(req *QueryListMetricRequest, accessId, accessSecret string) (*QueryListMetricResponse, error) {
	var pResponse QueryListMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
