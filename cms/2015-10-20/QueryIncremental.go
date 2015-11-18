package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryIncrementalRequest struct {
	RpcRequest
	Project      string
	Metric       string
	Period       string
	StartTime    string
	EndTime      string
	Dimensions   string
	TargetPeriod string
	Columns      string
	Extend       string
}

func (r *QueryIncrementalRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *QueryIncrementalRequest) GetProject() string {
	return r.Project
}
func (r *QueryIncrementalRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *QueryIncrementalRequest) GetMetric() string {
	return r.Metric
}
func (r *QueryIncrementalRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *QueryIncrementalRequest) GetPeriod() string {
	return r.Period
}
func (r *QueryIncrementalRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *QueryIncrementalRequest) GetStartTime() string {
	return r.StartTime
}
func (r *QueryIncrementalRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *QueryIncrementalRequest) GetEndTime() string {
	return r.EndTime
}
func (r *QueryIncrementalRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *QueryIncrementalRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *QueryIncrementalRequest) SetTargetPeriod(value string) {
	r.TargetPeriod = value
	r.QueryParams.Set("TargetPeriod", value)
}
func (r *QueryIncrementalRequest) GetTargetPeriod() string {
	return r.TargetPeriod
}
func (r *QueryIncrementalRequest) SetColumns(value string) {
	r.Columns = value
	r.QueryParams.Set("Columns", value)
}
func (r *QueryIncrementalRequest) GetColumns() string {
	return r.Columns
}
func (r *QueryIncrementalRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryIncrementalRequest) GetExtend() string {
	return r.Extend
}

func (r *QueryIncrementalRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryIncremental")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type QueryIncrementalResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func QueryIncremental(req *QueryIncrementalRequest, accessId, accessSecret string) (*QueryIncrementalResponse, error) {
	var pResponse QueryIncrementalResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
