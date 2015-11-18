package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetImgVpcInfoRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VpcId                string
	Label                string
}

func (r *GetImgVpcInfoRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *GetImgVpcInfoRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *GetImgVpcInfoRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *GetImgVpcInfoRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *GetImgVpcInfoRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *GetImgVpcInfoRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *GetImgVpcInfoRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GetImgVpcInfoRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *GetImgVpcInfoRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *GetImgVpcInfoRequest) GetRegion() string {
	return r.Region
}
func (r *GetImgVpcInfoRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *GetImgVpcInfoRequest) GetVpcId() string {
	return r.VpcId
}
func (r *GetImgVpcInfoRequest) SetLabel(value string) {
	r.Label = value
	r.QueryParams.Set("Label", value)
}
func (r *GetImgVpcInfoRequest) GetLabel() string {
	return r.Label
}

func (r *GetImgVpcInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetImgVpcInfo")
	r.SetProduct(Product)
}

type GetImgVpcInfoResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
	vip       string `xml:"vip" json:"vip"`
	vpcId     string `xml:"vpcId" json:"vpcId"`
}

func GetImgVpcInfo(req *GetImgVpcInfoRequest, accessId, accessSecret string) (*GetImgVpcInfoResponse, error) {
	var pResponse GetImgVpcInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
