package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BindBucketVipRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Region               string
	VpcId                string
	Vip                  string
	BucketName           string
}

func (r *BindBucketVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *BindBucketVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *BindBucketVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *BindBucketVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *BindBucketVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *BindBucketVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *BindBucketVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *BindBucketVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *BindBucketVipRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *BindBucketVipRequest) GetRegion() string {
	return r.Region
}
func (r *BindBucketVipRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *BindBucketVipRequest) GetVpcId() string {
	return r.VpcId
}
func (r *BindBucketVipRequest) SetVip(value string) {
	r.Vip = value
	r.QueryParams.Set("Vip", value)
}
func (r *BindBucketVipRequest) GetVip() string {
	return r.Vip
}
func (r *BindBucketVipRequest) SetBucketName(value string) {
	r.BucketName = value
	r.QueryParams.Set("BucketName", value)
}
func (r *BindBucketVipRequest) GetBucketName() string {
	return r.BucketName
}

func (r *BindBucketVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BindBucketVip")
	r.SetProduct(Product)
}

type BindBucketVipResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
}

func BindBucketVip(req *BindBucketVipRequest, accessId, accessSecret string) (*BindBucketVipResponse, error) {
	var pResponse BindBucketVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
