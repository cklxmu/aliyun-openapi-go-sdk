package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Instances            string
	Token                string
	AutoPay              bool
}

func (r *CreateInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateInstancesRequest) SetInstances(value string) {
	r.Instances = value
	r.QueryParams.Set("Instances", value)
}
func (r *CreateInstancesRequest) GetInstances() string {
	return r.Instances
}
func (r *CreateInstancesRequest) SetToken(value string) {
	r.Token = value
	r.QueryParams.Set("Token", value)
}
func (r *CreateInstancesRequest) GetToken() string {
	return r.Token
}
func (r *CreateInstancesRequest) SetAutoPay(value bool) {
	r.AutoPay = value
	r.QueryParams.Set("AutoPay", strconv.FormatBool(value))
}
func (r *CreateInstancesRequest) GetAutoPay() bool {
	return r.AutoPay
}

func (r *CreateInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateInstances")
	r.SetProduct(Product)
}

type CreateInstancesResponse struct {
	OrderId string `xml:"OrderId" json:"OrderId"`
}

func CreateInstances(req *CreateInstancesRequest, accessId, accessSecret string) (*CreateInstancesResponse, error) {
	var pResponse CreateInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
