package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetOssVpcInfoRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VpcId                string
	Label                string
}

func (r *GetOssVpcInfoRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *GetOssVpcInfoRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *GetOssVpcInfoRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *GetOssVpcInfoRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *GetOssVpcInfoRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *GetOssVpcInfoRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *GetOssVpcInfoRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GetOssVpcInfoRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *GetOssVpcInfoRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *GetOssVpcInfoRequest) GetRegion() string {
	return r.Region
}
func (r *GetOssVpcInfoRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *GetOssVpcInfoRequest) GetVpcId() string {
	return r.VpcId
}
func (r *GetOssVpcInfoRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *GetOssVpcInfoRequest) GetLabel() string {
	return r.Label
}

func (r *GetOssVpcInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetOssVpcInfo")
	r.SetProduct(Product)
}

type GetOssVpcInfoResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
	vip       string `xml:"vip" json:"vip"`
	vpcId     string `xml:"vpcId" json:"vpcId"`
}

func GetOssVpcInfo(req *GetOssVpcInfoRequest, accessId, accessSecret string) (*GetOssVpcInfoResponse, error) {
	var pResponse GetOssVpcInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
