package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteJobRequest struct {
	RpcRequest
	Id int
}

func (r *DeleteJobRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *DeleteJobRequest) GetId() int {
	return r.Id
}

func (r *DeleteJobRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteJob")
	r.SetProduct(Product)
}

type DeleteJobResponse struct {
}

func DeleteJob(req *DeleteJobRequest, accessId, accessSecret string) (*DeleteJobResponse, error) {
	var pResponse DeleteJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
