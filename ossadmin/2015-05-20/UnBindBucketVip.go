package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UnBindBucketVipRequest struct {
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

func (r *UnBindBucketVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *UnBindBucketVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *UnBindBucketVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *UnBindBucketVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *UnBindBucketVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *UnBindBucketVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *UnBindBucketVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *UnBindBucketVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *UnBindBucketVipRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *UnBindBucketVipRequest) GetRegion() string {
	return r.Region
}
func (r *UnBindBucketVipRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *UnBindBucketVipRequest) GetVpcId() string {
	return r.VpcId
}
func (r *UnBindBucketVipRequest) SetVip(value string) {
	r.Vip = value
	r.QueryParams.Set("Vip", value)
}
func (r *UnBindBucketVipRequest) GetVip() string {
	return r.Vip
}
func (r *UnBindBucketVipRequest) SetBucketName(value string) {
	r.BucketName = value
	r.QueryParams.Set("BucketName", value)
}
func (r *UnBindBucketVipRequest) GetBucketName() string {
	return r.BucketName
}

func (r *UnBindBucketVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UnBindBucketVip")
	r.SetProduct(Product)
}

type UnBindBucketVipResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
}

func UnBindBucketVip(req *UnBindBucketVipRequest, accessId, accessSecret string) (*UnBindBucketVipResponse, error) {
	var pResponse UnBindBucketVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
