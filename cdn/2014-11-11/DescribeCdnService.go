package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCdnServiceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
}

func (r *DescribeCdnServiceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnServiceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCdnServiceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCdnServiceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCdnServiceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnServiceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}

func (r *DescribeCdnServiceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCdnService")
	r.SetProduct(Product)
}

type DescribeCdnServiceResponse struct {
	InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
	OpeningTime        string `xml:"OpeningTime" json:"OpeningTime"`
	ChangingChargeType string `xml:"ChangingChargeType" json:"ChangingChargeType"`
	ChangingAffectTime string `xml:"ChangingAffectTime" json:"ChangingAffectTime"`
	OperationLocks     struct {
		LockReason []struct {
			LockReason string `xml:"LockReason" json:"LockReason"`
		} `xml:"LockReason" json:"LockReason"`
	} `xml:"OperationLocks" json:"OperationLocks"`
}

func DescribeCdnService(req *DescribeCdnServiceRequest, accessId, accessSecret string) (*DescribeCdnServiceResponse, error) {
	var pResponse DescribeCdnServiceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
