package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RunExecutePlanRequest struct {
	RpcRequest
	ExecutePlanId int
}

func (r *RunExecutePlanRequest) SetExecutePlanId(value int) {
	r.ExecutePlanId = value
	r.QueryParams.Set("ExecutePlanId", strconv.Itoa(value))
}
func (r *RunExecutePlanRequest) GetExecutePlanId() int {
	return r.ExecutePlanId
}

func (r *RunExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RunExecutePlan")
	r.SetProduct(Product)
}

type RunExecutePlanResponse struct {
}

func RunExecutePlan(req *RunExecutePlanRequest, accessId, accessSecret string) (*RunExecutePlanResponse, error) {
	var pResponse RunExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
