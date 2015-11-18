package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateOssVpcRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VirtualSwitchId      string
	VpcId                string
	Label                string
}

func (r *CreateOssVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateOssVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateOssVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateOssVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateOssVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateOssVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateOssVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateOssVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateOssVpcRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *CreateOssVpcRequest) GetRegion() string {
	return r.Region
}
func (r *CreateOssVpcRequest) SetVirtualSwitchId(value string) {
	r.VirtualSwitchId = value
	r.QueryParams.Set("VirtualSwitchId", value)
}
func (r *CreateOssVpcRequest) GetVirtualSwitchId() string {
	return r.VirtualSwitchId
}
func (r *CreateOssVpcRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateOssVpcRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateOssVpcRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *CreateOssVpcRequest) GetLabel() string {
	return r.Label
}

func (r *CreateOssVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateOssVpc")
	r.SetProduct(Product)
}

type CreateOssVpcResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
	vip       string `xml:"vip" json:"vip"`
	vpcId     string `xml:"vpcId" json:"vpcId"`
}

func CreateOssVpc(req *CreateOssVpcRequest, accessId, accessSecret string) (*CreateOssVpcResponse, error) {
	var pResponse CreateOssVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
