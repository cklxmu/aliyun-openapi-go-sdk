package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeUserInfoRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeUserInfoRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeUserInfoRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeUserInfoRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeUserInfoRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeUserInfoRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeUserInfoRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeUserInfoRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeUserInfoRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeUserInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeUserInfo")
	r.SetProduct(Product)
}

type DescribeUserInfoResponse struct {
	IsAuthentication      bool `xml:"IsAuthentication" json:"IsAuthentication"`
	IsFinance             bool `xml:"IsFinance" json:"IsFinance"`
	BalanceAmount         int  `xml:"BalanceAmount" json:"BalanceAmount"`
	AlreadyHasResourceNum struct {
		EcsBought int `xml:"EcsBought" json:"EcsBought"`
		KVSBought int `xml:"KVSBought" json:"KVSBought"`
	} `xml:"AlreadyHasResourceNum" json:"AlreadyHasResourceNum"`
}

func DescribeUserInfo(req *DescribeUserInfoRequest, accessId, accessSecret string) (*DescribeUserInfoResponse, error) {
	var pResponse DescribeUserInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
