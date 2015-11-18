package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyJobRequest struct {
	RpcRequest
	Id           int
	Name         string
	Type         string
	EnvParameter string
	FailAct      int
}

func (r *ModifyJobRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *ModifyJobRequest) GetId() int {
	return r.Id
}
func (r *ModifyJobRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *ModifyJobRequest) GetName() string {
	return r.Name
}
func (r *ModifyJobRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *ModifyJobRequest) GetType() string {
	return r.Type
}
func (r *ModifyJobRequest) SetEnvParameter(value string) {
	r.EnvParameter = value
	r.QueryParams.Set("EnvParameter", value)
}
func (r *ModifyJobRequest) GetEnvParameter() string {
	return r.EnvParameter
}
func (r *ModifyJobRequest) SetFailAct(value int) {
	r.FailAct = value
	r.QueryParams.Set("FailAct", strconv.Itoa(value))
}
func (r *ModifyJobRequest) GetFailAct() int {
	return r.FailAct
}

func (r *ModifyJobRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyJob")
	r.SetProduct(Product)
}

type ModifyJobResponse struct {
}

func ModifyJob(req *ModifyJobRequest, accessId, accessSecret string) (*ModifyJobResponse, error) {
	var pResponse ModifyJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
