package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLimitationRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Limitation           string
}

func (r *DescribeLimitationRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLimitationRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLimitationRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLimitationRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLimitationRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLimitationRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLimitationRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLimitationRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeLimitationRequest) SetLimitation(value string) {
	r.Limitation = value
	r.QueryParams.Set("Limitation", value)
}
func (r *DescribeLimitationRequest) GetLimitation() string {
	return r.Limitation
}

func (r *DescribeLimitationRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLimitation")
	r.SetProduct(Product)
}

type DescribeLimitationResponse struct {
	Limitation string `xml:"Limitation" json:"Limitation"`
	Value      string `xml:"Value" json:"Value"`
}

func DescribeLimitation(req *DescribeLimitationRequest, accessId, accessSecret string) (*DescribeLimitationResponse, error) {
	var pResponse DescribeLimitationResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
