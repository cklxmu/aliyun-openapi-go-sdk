package bss

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeCouponDetailRequest struct {
	RpcRequest
	CouponNumber string
}

func (r *DescribeCouponDetailRequest) SetCouponNumber(value string) {
	r.CouponNumber = value
	r.QueryParams.Set("CouponNumber", value)
}
func (r *DescribeCouponDetailRequest) GetCouponNumber() string {
	return r.CouponNumber
}

func (r *DescribeCouponDetailRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCouponDetail")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DescribeCouponDetailResponse struct {
	CouponTemplateId int    `xml:"CouponTemplateId" json:"CouponTemplateId"`
	TotalAmount      string `xml:"TotalAmount" json:"TotalAmount"`
	BalanceAmount    string `xml:"BalanceAmount" json:"BalanceAmount"`
	FrozenAmount     string `xml:"FrozenAmount" json:"FrozenAmount"`
	ExpiredAmount    string `xml:"ExpiredAmount" json:"ExpiredAmount"`
	DeliveryTime     string `xml:"DeliveryTime" json:"DeliveryTime"`
	ExpiredTime      string `xml:"ExpiredTime" json:"ExpiredTime"`
	CouponNumber     string `xml:"CouponNumber" json:"CouponNumber"`
	Status           string `xml:"Status" json:"Status"`
	Description      string `xml:"Description" json:"Description"`
	CreationTime     string `xml:"CreationTime" json:"CreationTime"`
	ModificationTime string `xml:"ModificationTime" json:"ModificationTime"`
	PriceLimit       string `xml:"PriceLimit" json:"PriceLimit"`
	Application      string `xml:"Application" json:"Application"`
	ProductCodes     struct {
		ProductCode []string `xml:"ProductCode" json:"ProductCode"`
	} `xml:"ProductCodes" json:"ProductCodes"`
	TradeTypes struct {
		TradeType []string `xml:"TradeType" json:"TradeType"`
	} `xml:"TradeTypes" json:"TradeTypes"`
}

func DescribeCouponDetail(req *DescribeCouponDetailRequest, accessId, accessSecret string) (*DescribeCouponDetailResponse, error) {
	var pResponse DescribeCouponDetailResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
