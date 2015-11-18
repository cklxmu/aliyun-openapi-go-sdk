package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopOssInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *StopOssInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StopOssInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StopOssInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StopOssInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StopOssInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StopOssInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StopOssInstanceRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *StopOssInstanceRequest) Getregion() string {
	return r.region
}
func (r *StopOssInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StopOssInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StopOssInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopOssInstance")
	r.SetProduct(Product)
}

type StopOssInstanceResponse struct {
}

func StopOssInstance(req *StopOssInstanceRequest, accessId, accessSecret string) (*StopOssInstanceResponse, error) {
	var pResponse StopOssInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
