package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UpdateUserRequest struct {
	RpcRequest
	InstanceQuota int
	Description   string
}

func (r *UpdateUserRequest) SetInstanceQuota(value int) {
	r.InstanceQuota = value
	r.QueryParams.Set("InstanceQuota", strconv.Itoa(value))
}
func (r *UpdateUserRequest) GetInstanceQuota() int {
	return r.InstanceQuota
}
func (r *UpdateUserRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *UpdateUserRequest) GetDescription() string {
	return r.Description
}

func (r *UpdateUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateUser")
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type UpdateUserResponse struct {
}

func UpdateUser(req *UpdateUserRequest, accessId, accessSecret string) (*UpdateUserResponse, error) {
	var pResponse UpdateUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
