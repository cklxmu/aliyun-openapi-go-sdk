package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopExecutePlanRequest struct {
	RpcRequest
	Id int
}

func (r *StopExecutePlanRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *StopExecutePlanRequest) GetId() int {
	return r.Id
}

func (r *StopExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopExecutePlan")
	r.SetProduct(Product)
}

type StopExecutePlanResponse struct {
}

func StopExecutePlan(req *StopExecutePlanRequest, accessId, accessSecret string) (*StopExecutePlanResponse, error) {
	var pResponse StopExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
