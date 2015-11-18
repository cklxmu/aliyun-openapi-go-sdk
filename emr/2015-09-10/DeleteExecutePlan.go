package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteExecutePlanRequest struct {
	RpcRequest
	Id int
}

func (r *DeleteExecutePlanRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *DeleteExecutePlanRequest) GetId() int {
	return r.Id
}

func (r *DeleteExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteExecutePlan")
	r.SetProduct(Product)
}

type DeleteExecutePlanResponse struct {
}

func DeleteExecutePlan(req *DeleteExecutePlanRequest, accessId, accessSecret string) (*DeleteExecutePlanResponse, error) {
	var pResponse DeleteExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
