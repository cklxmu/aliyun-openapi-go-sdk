package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateJobRequest struct {
	RpcRequest
	JobName    string
	JobType    string
	Parameter  string
	JobFailAct int
}

func (r *CreateJobRequest) SetJobName(value string) {
	r.JobName = value
	r.QueryParams.Set("JobName", value)
}
func (r *CreateJobRequest) GetJobName() string {
	return r.JobName
}
func (r *CreateJobRequest) SetJobType(value string) {
	r.JobType = value
	r.QueryParams.Set("JobType", value)
}
func (r *CreateJobRequest) GetJobType() string {
	return r.JobType
}
func (r *CreateJobRequest) SetParameter(value string) {
	r.Parameter = value
	r.QueryParams.Set("Parameter", value)
}
func (r *CreateJobRequest) GetParameter() string {
	return r.Parameter
}
func (r *CreateJobRequest) SetJobFailAct(value int) {
	r.JobFailAct = value
	r.QueryParams.Set("JobFailAct", strconv.Itoa(value))
}
func (r *CreateJobRequest) GetJobFailAct() int {
	return r.JobFailAct
}

func (r *CreateJobRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateJob")
	r.SetProduct(Product)
}

type CreateJobResponse struct {
	Id int `xml:"Id" json:"Id"`
}

func CreateJob(req *CreateJobRequest, accessId, accessSecret string) (*CreateJobResponse, error) {
	var pResponse CreateJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
