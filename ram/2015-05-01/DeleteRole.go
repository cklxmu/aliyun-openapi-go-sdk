package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteRoleRequest struct {
	RpcRequest
	RoleName string
}

func (r *DeleteRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *DeleteRoleRequest) GetRoleName() string {
	return r.RoleName
}

func (r *DeleteRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteRoleResponse struct {
}

func DeleteRole(req *DeleteRoleRequest, accessId, accessSecret string) (*DeleteRoleResponse, error) {
	var pResponse DeleteRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
