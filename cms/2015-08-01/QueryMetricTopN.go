package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryMetricTopNRequest struct {
	RpcRequest
	Project    string
	Metric     string
	Period     string
	StartTime  string
	EndTime    string
	Dimensions string
	ValueKey   string
	Top        string
	Extend     string
}

func (r *QueryMetricTopNRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *QueryMetricTopNRequest) GetProject() string {
	return r.Project
}
func (r *QueryMetricTopNRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *QueryMetricTopNRequest) GetMetric() string {
	return r.Metric
}
func (r *QueryMetricTopNRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *QueryMetricTopNRequest) GetPeriod() string {
	return r.Period
}
func (r *QueryMetricTopNRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *QueryMetricTopNRequest) GetStartTime() string {
	return r.StartTime
}
func (r *QueryMetricTopNRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *QueryMetricTopNRequest) GetEndTime() string {
	return r.EndTime
}
func (r *QueryMetricTopNRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *QueryMetricTopNRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *QueryMetricTopNRequest) SetValueKey(value string) {
	r.ValueKey = value
	r.QueryParams.Set("ValueKey", value)
}
func (r *QueryMetricTopNRequest) GetValueKey() string {
	return r.ValueKey
}
func (r *QueryMetricTopNRequest) SetTop(value string) {
	r.Top = value
	r.QueryParams.Set("Top", value)
}
func (r *QueryMetricTopNRequest) GetTop() string {
	return r.Top
}
func (r *QueryMetricTopNRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryMetricTopNRequest) GetExtend() string {
	return r.Extend
}

func (r *QueryMetricTopNRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryMetricTopN")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type QueryMetricTopNResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func QueryMetricTopN(req *QueryMetricTopNRequest, accessId, accessSecret string) (*QueryMetricTopNResponse, error) {
	var pResponse QueryMetricTopNResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
