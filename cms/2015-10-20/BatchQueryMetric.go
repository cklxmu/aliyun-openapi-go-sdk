package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type BatchQueryMetricRequest struct {
	RpcRequest
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	Extend     string
	Filter     string
}

func (r *BatchQueryMetricRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *BatchQueryMetricRequest) GetProject() string {
	return r.Project
}
func (r *BatchQueryMetricRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *BatchQueryMetricRequest) GetMetric() string {
	return r.Metric
}
func (r *BatchQueryMetricRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *BatchQueryMetricRequest) GetPeriod() string {
	return r.Period
}
func (r *BatchQueryMetricRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *BatchQueryMetricRequest) GetStartTime() string {
	return r.StartTime
}
func (r *BatchQueryMetricRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *BatchQueryMetricRequest) GetEndTime() string {
	return r.EndTime
}
func (r *BatchQueryMetricRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *BatchQueryMetricRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *BatchQueryMetricRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *BatchQueryMetricRequest) GetExtend() string {
	return r.Extend
}
func (r *BatchQueryMetricRequest) SetFilter(value string) {
	r.Filter = value
	r.QueryParams.Set("Filter", value)
}
func (r *BatchQueryMetricRequest) GetFilter() string {
	return r.Filter
}

func (r *BatchQueryMetricRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchQueryMetric")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type BatchQueryMetricResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func BatchQueryMetric(req *BatchQueryMetricRequest, accessId, accessSecret string) (*BatchQueryMetricResponse, error) {
	var pResponse BatchQueryMetricResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
