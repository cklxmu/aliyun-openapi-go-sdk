package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListExecutePlanNodeInstancesRequest struct {
	RpcRequest
	ExecutePlanWorkNodeId int
}

func (r *ListExecutePlanNodeInstancesRequest) SetExecutePlanWorkNodeId(value int) {
	r.ExecutePlanWorkNodeId = value
	r.QueryParams.Set("ExecutePlanWorkNodeId", strconv.Itoa(value))
}
func (r *ListExecutePlanNodeInstancesRequest) GetExecutePlanWorkNodeId() int {
	return r.ExecutePlanWorkNodeId
}

func (r *ListExecutePlanNodeInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListExecutePlanNodeInstances")
	r.SetProduct(Product)
}

type ListExecutePlanNodeInstancesResponse struct {
	ExecutePlanNodeInstance struct {
		ExecutePlanNodeInstanceInfo []struct {
			ApplicationId string `xml:"ApplicationId" json:"ApplicationId"`
			InstanceInfo  string `xml:"InstanceInfo" json:"InstanceInfo"`
			ContainerInfo string `xml:"ContainerInfo" json:"ContainerInfo"`
		} `xml:"ExecutePlanNodeInstanceInfo" json:"ExecutePlanNodeInstanceInfo"`
	} `xml:"ExecutePlanNodeInstance" json:"ExecutePlanNodeInstance"`
}

func ListExecutePlanNodeInstances(req *ListExecutePlanNodeInstancesRequest, accessId, accessSecret string) (*ListExecutePlanNodeInstancesResponse, error) {
	var pResponse ListExecutePlanNodeInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
