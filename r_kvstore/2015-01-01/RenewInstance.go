package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RenewInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Capacity             string
	Period               int
	AutoPay              bool
}

func (r *RenewInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RenewInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RenewInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RenewInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RenewInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RenewInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *RenewInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *RenewInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *RenewInstanceRequest) SetCapacity(value string) {
	r.Capacity = value
	r.QueryParams.Set("Capacity", value)
}
func (r *RenewInstanceRequest) GetCapacity() string {
	return r.Capacity
}
func (r *RenewInstanceRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetPeriod() int {
	return r.Period
}
func (r *RenewInstanceRequest) SetAutoPay(value bool) {
	r.AutoPay = value
	r.QueryParams.Set("AutoPay", strconv.FormatBool(value))
}
func (r *RenewInstanceRequest) GetAutoPay() bool {
	return r.AutoPay
}

func (r *RenewInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RenewInstance")
	r.SetProduct(Product)
}

type RenewInstanceResponse struct {
	OrderId string `xml:"OrderId" json:"OrderId"`
	EndTime string `xml:"EndTime" json:"EndTime"`
}

func RenewInstance(req *RenewInstanceRequest, accessId, accessSecret string) (*RenewInstanceResponse, error) {
	var pResponse RenewInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
