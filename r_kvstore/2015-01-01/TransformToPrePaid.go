package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type TransformToPrePaidRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Period               int
	AutoPay              bool
}

func (r *TransformToPrePaidRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *TransformToPrePaidRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *TransformToPrePaidRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *TransformToPrePaidRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *TransformToPrePaidRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *TransformToPrePaidRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *TransformToPrePaidRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *TransformToPrePaidRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *TransformToPrePaidRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *TransformToPrePaidRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *TransformToPrePaidRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *TransformToPrePaidRequest) GetPeriod() int {
	return r.Period
}
func (r *TransformToPrePaidRequest) SetAutoPay(value bool) {
	r.AutoPay = value
	r.QueryParams.Set("AutoPay", strconv.FormatBool(value))
}
func (r *TransformToPrePaidRequest) GetAutoPay() bool {
	return r.AutoPay
}

func (r *TransformToPrePaidRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TransformToPrePaid")
	r.SetProduct(Product)
}

type TransformToPrePaidResponse struct {
	OrderId string `xml:"OrderId" json:"OrderId"`
	EndTime string `xml:"EndTime" json:"EndTime"`
}

func TransformToPrePaid(req *TransformToPrePaidRequest, accessId, accessSecret string) (*TransformToPrePaidResponse, error) {
	var pResponse TransformToPrePaidResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
