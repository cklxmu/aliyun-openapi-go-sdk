package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryMetricRequest struct {
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

func (r *QueryMetricRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *QueryMetricRequest) GetProject() string {
	return r.Project
}
func (r *QueryMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *QueryMetricRequest) GetMetric() string {
	return r.Metric
}
func (r *QueryMetricRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *QueryMetricRequest) GetPeriod() string {
	return r.Period
}
func (r *QueryMetricRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *QueryMetricRequest) GetStartTime() string {
	return r.StartTime
}
func (r *QueryMetricRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *QueryMetricRequest) GetEndTime() string {
	return r.EndTime
}
func (r *QueryMetricRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *QueryMetricRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *QueryMetricRequest) SetPage(value string) {
	r.Page = value
	r.QueryParams.Set("Page", value)
}
func (r *QueryMetricRequest) GetPage() string {
	return r.Page
}
func (r *QueryMetricRequest) SetLength(value string) {
	r.Length = value
	r.QueryParams.Set("Length", value)
}
func (r *QueryMetricRequest) GetLength() string {
	return r.Length
}
func (r *QueryMetricRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryMetricRequest) GetExtend() string {
	return r.Extend
}

func (r *QueryMetricRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryMetric")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type QueryMetricResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func QueryMetric(req *QueryMetricRequest, accessId, accessSecret string) (*QueryMetricResponse, error) {
	var pResponse QueryMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
