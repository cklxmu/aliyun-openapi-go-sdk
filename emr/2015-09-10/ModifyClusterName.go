package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyClusterNameRequest struct {
	RpcRequest
	ClusterId   int
	ClusterName string
}

func (r *ModifyClusterNameRequest) SetClusterId(value int) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", strconv.Itoa(value))
}
func (r *ModifyClusterNameRequest) GetClusterId() int {
	return r.ClusterId
}
func (r *ModifyClusterNameRequest) SetClusterName(value string) {
	r.ClusterName = value
	r.QueryParams.Set("ClusterName", value)
}
func (r *ModifyClusterNameRequest) GetClusterName() string {
	return r.ClusterName
}

func (r *ModifyClusterNameRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyClusterName")
	r.SetProduct(Product)
}

type ModifyClusterNameResponse struct {
}

func ModifyClusterName(req *ModifyClusterNameRequest, accessId, accessSecret string) (*ModifyClusterNameResponse, error) {
	var pResponse ModifyClusterNameResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
