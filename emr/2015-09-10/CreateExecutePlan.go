package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateExecutePlanRequest struct {
	RpcRequest
	ClusterId    int
	Name         string
	Strategy     int
	TimeInterval int
	StartTime    string
	TimeUnit     string
	JobId        string
}

func (r *CreateExecutePlanRequest) SetClusterId(value int) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", strconv.Itoa(value))
}
func (r *CreateExecutePlanRequest) GetClusterId() int {
	return r.ClusterId
}
func (r *CreateExecutePlanRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *CreateExecutePlanRequest) GetName() string {
	return r.Name
}
func (r *CreateExecutePlanRequest) SetStrategy(value int) {
	r.Strategy = value
	r.QueryParams.Set("Strategy", strconv.Itoa(value))
}
func (r *CreateExecutePlanRequest) GetStrategy() int {
	return r.Strategy
}
func (r *CreateExecutePlanRequest) SetTimeInterval(value int) {
	r.TimeInterval = value
	r.QueryParams.Set("TimeInterval", strconv.Itoa(value))
}
func (r *CreateExecutePlanRequest) GetTimeInterval() int {
	return r.TimeInterval
}
func (r *CreateExecutePlanRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *CreateExecutePlanRequest) GetStartTime() string {
	return r.StartTime
}
func (r *CreateExecutePlanRequest) SetTimeUnit(value string) {
	r.TimeUnit = value
	r.QueryParams.Set("TimeUnit", value)
}
func (r *CreateExecutePlanRequest) GetTimeUnit() string {
	return r.TimeUnit
}
func (r *CreateExecutePlanRequest) SetJobId(value string) {
	r.JobId = value
	r.QueryParams.Set("JobId", value)
}
func (r *CreateExecutePlanRequest) GetJobId() string {
	return r.JobId
}

func (r *CreateExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateExecutePlan")
	r.SetProduct(Product)
}

type CreateExecutePlanResponse struct {
	Id int `xml:"Id" json:"Id"`
}

func CreateExecutePlan(req *CreateExecutePlanRequest, accessId, accessSecret string) (*CreateExecutePlanResponse, error) {
	var pResponse CreateExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
