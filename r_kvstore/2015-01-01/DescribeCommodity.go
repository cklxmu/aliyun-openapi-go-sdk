package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCommodityRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	CommodityCode        string
	InstanceId           string
	OrderType            string
}

func (r *DescribeCommodityRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCommodityRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCommodityRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCommodityRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCommodityRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCommodityRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeCommodityRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeCommodityRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeCommodityRequest) SetCommodityCode(value string) {
	r.CommodityCode = value
	r.QueryParams.Set("CommodityCode", value)
}
func (r *DescribeCommodityRequest) GetCommodityCode() string {
	return r.CommodityCode
}
func (r *DescribeCommodityRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeCommodityRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeCommodityRequest) SetOrderType(value string) {
	r.OrderType = value
	r.QueryParams.Set("OrderType", value)
}
func (r *DescribeCommodityRequest) GetOrderType() string {
	return r.OrderType
}

func (r *DescribeCommodityRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCommodity")
	r.SetProduct(Product)
}

type DescribeCommodityResponse struct {
	Commodity string `xml:"Commodity" json:"Commodity"`
}

func DescribeCommodity(req *DescribeCommodityRequest, accessId, accessSecret string) (*DescribeCommodityResponse, error) {
	var pResponse DescribeCommodityResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
