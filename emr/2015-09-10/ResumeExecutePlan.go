package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ResumeExecutePlanRequest struct {
	RpcRequest
	Id int
}

func (r *ResumeExecutePlanRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *ResumeExecutePlanRequest) GetId() int {
	return r.Id
}

func (r *ResumeExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ResumeExecutePlan")
	r.SetProduct(Product)
}

type ResumeExecutePlanResponse struct {
}

func ResumeExecutePlan(req *ResumeExecutePlanRequest, accessId, accessSecret string) (*ResumeExecutePlanResponse, error) {
	var pResponse ResumeExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
