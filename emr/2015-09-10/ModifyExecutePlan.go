package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyExecutePlanRequest struct {
	RpcRequest
	ClusterId     int
	ExecutePlanId int
	Name          string
	Strategy      int
	TimeInterval  int
	StartTime     string
	TimeUnit      string
	JobId         string
}

func (r *ModifyExecutePlanRequest) SetClusterId(value int) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", strconv.Itoa(value))
}
func (r *ModifyExecutePlanRequest) GetClusterId() int {
	return r.ClusterId
}
func (r *ModifyExecutePlanRequest) SetExecutePlanId(value int) {
	r.ExecutePlanId = value
	r.QueryParams.Set("ExecutePlanId", strconv.Itoa(value))
}
func (r *ModifyExecutePlanRequest) GetExecutePlanId() int {
	return r.ExecutePlanId
}
func (r *ModifyExecutePlanRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *ModifyExecutePlanRequest) GetName() string {
	return r.Name
}
func (r *ModifyExecutePlanRequest) SetStrategy(value int) {
	r.Strategy = value
	r.QueryParams.Set("Strategy", strconv.Itoa(value))
}
func (r *ModifyExecutePlanRequest) GetStrategy() int {
	return r.Strategy
}
func (r *ModifyExecutePlanRequest) SetTimeInterval(value int) {
	r.TimeInterval = value
	r.QueryParams.Set("TimeInterval", strconv.Itoa(value))
}
func (r *ModifyExecutePlanRequest) GetTimeInterval() int {
	return r.TimeInterval
}
func (r *ModifyExecutePlanRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *ModifyExecutePlanRequest) GetStartTime() string {
	return r.StartTime
}
func (r *ModifyExecutePlanRequest) SetTimeUnit(value string) {
	r.TimeUnit = value
	r.QueryParams.Set("TimeUnit", value)
}
func (r *ModifyExecutePlanRequest) GetTimeUnit() string {
	return r.TimeUnit
}
func (r *ModifyExecutePlanRequest) SetJobId(value string) {
	r.JobId = value
	r.QueryParams.Set("JobId", value)
}
func (r *ModifyExecutePlanRequest) GetJobId() string {
	return r.JobId
}

func (r *ModifyExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyExecutePlan")
	r.SetProduct(Product)
}

type ModifyExecutePlanResponse struct {
	ExecutePlanId int `xml:"ExecutePlanId" json:"ExecutePlanId"`
}

func ModifyExecutePlan(req *ModifyExecutePlanRequest, accessId, accessSecret string) (*ModifyExecutePlanResponse, error) {
	var pResponse ModifyExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
