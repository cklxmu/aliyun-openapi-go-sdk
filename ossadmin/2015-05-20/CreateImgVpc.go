package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateImgVpcRequest struct {
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

func (r *CreateImgVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateImgVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateImgVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateImgVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateImgVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateImgVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateImgVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateImgVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateImgVpcRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *CreateImgVpcRequest) GetRegion() string {
	return r.Region
}
func (r *CreateImgVpcRequest) SetVirtualSwitchId(value string) {
	r.VirtualSwitchId = value
	r.QueryParams.Set("VirtualSwitchId", value)
}
func (r *CreateImgVpcRequest) GetVirtualSwitchId() string {
	return r.VirtualSwitchId
}
func (r *CreateImgVpcRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateImgVpcRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateImgVpcRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *CreateImgVpcRequest) GetLabel() string {
	return r.Label
}

func (r *CreateImgVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateImgVpc")
	r.SetProduct(Product)
}

type CreateImgVpcResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
	vip       string `xml:"vip" json:"vip"`
	vpcId     string `xml:"vpcId" json:"vpcId"`
}

func CreateImgVpc(req *CreateImgVpcRequest, accessId, accessSecret string) (*CreateImgVpcResponse, error) {
	var pResponse CreateImgVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
