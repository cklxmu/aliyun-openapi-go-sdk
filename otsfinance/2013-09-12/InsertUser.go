package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type InsertUserRequest struct {
	RpcRequest
	InstanceQuota int
	Description   string
}

func (r *InsertUserRequest) SetInstanceQuota(value int) {
	r.InstanceQuota = value
	r.QueryParams.Set("InstanceQuota", strconv.Itoa(value))
}
func (r *InsertUserRequest) GetInstanceQuota() int {
	return r.InstanceQuota
}
func (r *InsertUserRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *InsertUserRequest) GetDescription() string {
	return r.Description
}

func (r *InsertUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("InsertUser")
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type InsertUserResponse struct {
}

func InsertUser(req *InsertUserRequest, accessId, accessSecret string) (*InsertUserResponse, error) {
	var pResponse InsertUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
