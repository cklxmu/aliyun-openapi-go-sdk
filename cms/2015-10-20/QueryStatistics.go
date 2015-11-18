package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryStatisticsRequest struct {
	RpcRequest
	Project      string
	Metric       string
	Period       string
	StartTime    string
	EndTime      string
	Dimensions   string
	TargetPeriod string
	Function     string
	Extend       string
}

func (r *QueryStatisticsRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *QueryStatisticsRequest) GetProject() string {
	return r.Project
}
func (r *QueryStatisticsRequest) SetMetric(value string) {
	r.Metric = value
	r.QueryParams.Set("Metric", value)
}
func (r *QueryStatisticsRequest) GetMetric() string {
	return r.Metric
}
func (r *QueryStatisticsRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *QueryStatisticsRequest) GetPeriod() string {
	return r.Period
}
func (r *QueryStatisticsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *QueryStatisticsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *QueryStatisticsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *QueryStatisticsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *QueryStatisticsRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *QueryStatisticsRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *QueryStatisticsRequest) SetTargetPeriod(value string) {
	r.TargetPeriod = value
	r.QueryParams.Set("TargetPeriod", value)
}
func (r *QueryStatisticsRequest) GetTargetPeriod() string {
	return r.TargetPeriod
}
func (r *QueryStatisticsRequest) SetFunction(value string) {
	r.Function = value
	r.QueryParams.Set("Function", value)
}
func (r *QueryStatisticsRequest) GetFunction() string {
	return r.Function
}
func (r *QueryStatisticsRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryStatisticsRequest) GetExtend() string {
	return r.Extend
}

func (r *QueryStatisticsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryStatistics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type QueryStatisticsResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func QueryStatistics(req *QueryStatisticsRequest, accessId, accessSecret string) (*QueryStatisticsResponse, error) {
	var pResponse QueryStatisticsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
