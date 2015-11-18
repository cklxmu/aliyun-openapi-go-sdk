package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type KillExecutePlanRecordNodeRequest struct {
	RpcRequest
	ExecutePlanRecordNodeId int
}

func (r *KillExecutePlanRecordNodeRequest) SetExecutePlanRecordNodeId(value int) {
	r.ExecutePlanRecordNodeId = value
	r.QueryParams.Set("ExecutePlanRecordNodeId", strconv.Itoa(value))
}
func (r *KillExecutePlanRecordNodeRequest) GetExecutePlanRecordNodeId() int {
	return r.ExecutePlanRecordNodeId
}

func (r *KillExecutePlanRecordNodeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("KillExecutePlanRecordNode")
	r.SetProduct(Product)
}

type KillExecutePlanRecordNodeResponse struct {
}

func KillExecutePlanRecordNode(req *KillExecutePlanRecordNodeRequest, accessId, accessSecret string) (*KillExecutePlanRecordNodeResponse, error) {
	var pResponse KillExecutePlanRecordNodeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
