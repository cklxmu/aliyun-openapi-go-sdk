package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteScalingConfigurationRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	ScalingConfigurationId string
	OwnerAccount           string
}

func (r *DeleteScalingConfigurationRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingConfigurationRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteScalingConfigurationRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteScalingConfigurationRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteScalingConfigurationRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingConfigurationRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteScalingConfigurationRequest) SetScalingConfigurationId(value string) {
	r.ScalingConfigurationId = value
	r.QueryParams.Set("ScalingConfigurationId", value)
}
func (r *DeleteScalingConfigurationRequest) GetScalingConfigurationId() string {
	return r.ScalingConfigurationId
}
func (r *DeleteScalingConfigurationRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteScalingConfigurationRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteScalingConfigurationRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteScalingConfiguration")
	r.SetProduct(Product)
}

type DeleteScalingConfigurationResponse struct {
}

func DeleteScalingConfiguration(req *DeleteScalingConfigurationRequest, accessId, accessSecret string) (*DeleteScalingConfigurationResponse, error) {
	var pResponse DeleteScalingConfigurationResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
