package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RestartOssInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *RestartOssInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RestartOssInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RestartOssInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RestartOssInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RestartOssInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RestartOssInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RestartOssInstanceRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *RestartOssInstanceRequest) Getregion() string {
	return r.region
}
func (r *RestartOssInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RestartOssInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RestartOssInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RestartOssInstance")
	r.SetProduct(Product)
}

type RestartOssInstanceResponse struct {
}

func RestartOssInstance(req *RestartOssInstanceRequest, accessId, accessSecret string) (*RestartOssInstanceResponse, error) {
	var pResponse RestartOssInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
