package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteImgVpcRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VpcId                string
	Label                string
}

func (r *DeleteImgVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteImgVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteImgVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteImgVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteImgVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteImgVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteImgVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteImgVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DeleteImgVpcRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *DeleteImgVpcRequest) GetRegion() string {
	return r.Region
}
func (r *DeleteImgVpcRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DeleteImgVpcRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DeleteImgVpcRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *DeleteImgVpcRequest) GetLabel() string {
	return r.Label
}

func (r *DeleteImgVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteImgVpc")
	r.SetProduct(Product)
}

type DeleteImgVpcResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
}

func DeleteImgVpc(req *DeleteImgVpcRequest, accessId, accessSecret string) (*DeleteImgVpcResponse, error) {
	var pResponse DeleteImgVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
