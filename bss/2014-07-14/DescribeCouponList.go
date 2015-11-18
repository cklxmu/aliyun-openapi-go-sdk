package bss

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCouponListRequest struct {
	RpcRequest
	Status            string
	StartDeliveryTime string
	EndDeliveryTime   string
	PageSize          int
	PageNum           int
}

func (r *DescribeCouponListRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeCouponListRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeCouponListRequest) SetStartDeliveryTime(value string) {
	r.StartDeliveryTime = value
	r.QueryParams.Set("StartDeliveryTime", value)
}
func (r *DescribeCouponListRequest) GetStartDeliveryTime() string {
	return r.StartDeliveryTime
}
func (r *DescribeCouponListRequest) SetEndDeliveryTime(value string) {
	r.EndDeliveryTime = value
	r.QueryParams.Set("EndDeliveryTime", value)
}
func (r *DescribeCouponListRequest) GetEndDeliveryTime() string {
	return r.EndDeliveryTime
}
func (r *DescribeCouponListRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeCouponListRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeCouponListRequest) SetPageNum(value int) {
	r.PageNum = value
	r.QueryParams.Set("PageNum", strconv.Itoa(value))
}
func (r *DescribeCouponListRequest) GetPageNum() int {
	return r.PageNum
}

func (r *DescribeCouponListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCouponList")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DescribeCouponListResponse struct {
	Coupons struct {
		Coupon []struct {
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
		} `xml:"Coupon" json:"Coupon"`
	} `xml:"Coupons" json:"Coupons"`
}

func DescribeCouponList(req *DescribeCouponListRequest, accessId, accessSecret string) (*DescribeCouponListResponse, error) {
	var pResponse DescribeCouponListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
