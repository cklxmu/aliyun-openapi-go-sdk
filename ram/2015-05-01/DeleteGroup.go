package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteGroupRequest struct {
	RpcRequest
	GroupName string
}

func (r *DeleteGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *DeleteGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *DeleteGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteGroupResponse struct {
}

func DeleteGroup(req *DeleteGroupRequest, accessId, accessSecret string) (*DeleteGroupResponse, error) {
	var pResponse DeleteGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
