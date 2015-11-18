package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReleaseOssInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *ReleaseOssInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReleaseOssInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReleaseOssInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReleaseOssInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReleaseOssInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReleaseOssInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReleaseOssInstanceRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *ReleaseOssInstanceRequest) Getregion() string {
	return r.region
}
func (r *ReleaseOssInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReleaseOssInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ReleaseOssInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReleaseOssInstance")
	r.SetProduct(Product)
}

type ReleaseOssInstanceResponse struct {
}

func ReleaseOssInstance(req *ReleaseOssInstanceRequest, accessId, accessSecret string) (*ReleaseOssInstanceResponse, error) {
	var pResponse ReleaseOssInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
