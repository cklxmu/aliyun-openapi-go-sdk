package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type VerifyPasswordRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Password             string
}

func (r *VerifyPasswordRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *VerifyPasswordRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *VerifyPasswordRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *VerifyPasswordRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *VerifyPasswordRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *VerifyPasswordRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *VerifyPasswordRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *VerifyPasswordRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *VerifyPasswordRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *VerifyPasswordRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *VerifyPasswordRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *VerifyPasswordRequest) GetPassword() string {
	return r.Password
}

func (r *VerifyPasswordRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("VerifyPassword")
	r.SetProduct(Product)
}

type VerifyPasswordResponse struct {
	Result bool `xml:"Result" json:"Result"`
}

func VerifyPassword(req *VerifyPasswordRequest, accessId, accessSecret string) (*VerifyPasswordResponse, error) {
	var pResponse VerifyPasswordResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
