package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteOssVpcRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VpcId                string
	Label                string
}

func (r *DeleteOssVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteOssVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteOssVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteOssVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteOssVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteOssVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteOssVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteOssVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DeleteOssVpcRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *DeleteOssVpcRequest) GetRegion() string {
	return r.Region
}
func (r *DeleteOssVpcRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DeleteOssVpcRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DeleteOssVpcRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *DeleteOssVpcRequest) GetLabel() string {
	return r.Label
}

func (r *DeleteOssVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteOssVpc")
	r.SetProduct(Product)
}

type DeleteOssVpcResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
}

func DeleteOssVpc(req *DeleteOssVpcRequest, accessId, accessSecret string) (*DeleteOssVpcResponse, error) {
	var pResponse DeleteOssVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
