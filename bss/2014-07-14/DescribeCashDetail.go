package bss

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeCashDetailRequest struct {
	RpcRequest
}

func (r *DescribeCashDetailRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCashDetail")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DescribeCashDetailResponse struct {
	BalanceAmount        string `xml:"BalanceAmount" json:"BalanceAmount"`
	AmountOwed           string `xml:"AmountOwed" json:"AmountOwed"`
	EnableThresholdAlert string `xml:"EnableThresholdAlert" json:"EnableThresholdAlert"`
	MiniAlertThreshold   int    `xml:"MiniAlertThreshold" json:"MiniAlertThreshold"`
	FrozenAmount         string `xml:"FrozenAmount" json:"FrozenAmount"`
	CreditCardAmount     string `xml:"CreditCardAmount" json:"CreditCardAmount"`
	RemmitanceAmount     string `xml:"RemmitanceAmount" json:"RemmitanceAmount"`
	CreditLimit          string `xml:"CreditLimit" json:"CreditLimit"`
	AvailableCredit      string `xml:"AvailableCredit" json:"AvailableCredit"`
}

func DescribeCashDetail(req *DescribeCashDetailRequest, accessId, accessSecret string) (*DescribeCashDetailResponse, error) {
	var pResponse DescribeCashDetailResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
