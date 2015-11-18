package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReleaseClusterRequest struct {
	RpcRequest
	ClusterId int
}

func (r *ReleaseClusterRequest) SetClusterId(value int) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", strconv.Itoa(value))
}
func (r *ReleaseClusterRequest) GetClusterId() int {
	return r.ClusterId
}

func (r *ReleaseClusterRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReleaseCluster")
	r.SetProduct(Product)
}

type ReleaseClusterResponse struct {
}

func ReleaseCluster(req *ReleaseClusterRequest, accessId, accessSecret string) (*ReleaseClusterResponse, error) {
	var pResponse ReleaseClusterResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
