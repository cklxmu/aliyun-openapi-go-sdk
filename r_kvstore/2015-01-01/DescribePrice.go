package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribePriceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	Capacity             int
	OrderType            string
	ZoneId               string
	ChargeType           string
	Period               int
	Quantity             int
	InstanceId           string
}

func (r *DescribePriceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribePriceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribePriceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribePriceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribePriceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribePriceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribePriceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribePriceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribePriceRequest) SetCapacity(value int) {
	r.Capacity = value
	r.QueryParams.Set("Capacity", strconv.Itoa(value))
}
func (r *DescribePriceRequest) GetCapacity() int {
	return r.Capacity
}
func (r *DescribePriceRequest) SetOrderType(value string) {
	r.OrderType = value
	r.QueryParams.Set("OrderType", value)
}
func (r *DescribePriceRequest) GetOrderType() string {
	return r.OrderType
}
func (r *DescribePriceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribePriceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *DescribePriceRequest) SetChargeType(value string) {
	r.ChargeType = value
	r.QueryParams.Set("ChargeType", value)
}
func (r *DescribePriceRequest) GetChargeType() string {
	return r.ChargeType
}
func (r *DescribePriceRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *DescribePriceRequest) GetPeriod() int {
	return r.Period
}
func (r *DescribePriceRequest) SetQuantity(value int) {
	r.Quantity = value
	r.QueryParams.Set("Quantity", strconv.Itoa(value))
}
func (r *DescribePriceRequest) GetQuantity() int {
	return r.Quantity
}
func (r *DescribePriceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribePriceRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DescribePriceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribePrice")
	r.SetProduct(Product)
}

type DescribePriceResponse struct {
	Rules struct {
		Rule []struct {
			RuleDescId int    `xml:"RuleDescId" json:"RuleDescId"`
			Name       string `xml:"Name" json:"Name"`
			Title      string `xml:"Title" json:"Title"`
		} `xml:"Rule" json:"Rule"`
	} `xml:"Rules" json:"Rules"`
	Order struct {
		OriginalAmount float32 `xml:"OriginalAmount" json:"OriginalAmount"`
		TradeAmount    float32 `xml:"TradeAmount" json:"TradeAmount"`
		DiscountAmount float32 `xml:"DiscountAmount" json:"DiscountAmount"`
		RuleIds        struct {
			RuleId []string `xml:"RuleId" json:"RuleId"`
		} `xml:"RuleIds" json:"RuleIds"`
	} `xml:"Order" json:"Order"`
}

func DescribePrice(req *DescribePriceRequest, accessId, accessSecret string) (*DescribePriceResponse, error) {
	var pResponse DescribePriceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
