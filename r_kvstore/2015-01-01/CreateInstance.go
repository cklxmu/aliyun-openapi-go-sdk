package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceName         string
	Password             string
	Capacity             int
	ZoneId               string
	Config               string
	ChargeType           string
	Period               int
	Token                string
}

func (r *CreateInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *CreateInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *CreateInstanceRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *CreateInstanceRequest) GetPassword() string {
	return r.Password
}
func (r *CreateInstanceRequest) SetCapacity(value int) {
	r.Capacity = value
	r.QueryParams.Set("Capacity", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetCapacity() int {
	return r.Capacity
}
func (r *CreateInstanceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateInstanceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateInstanceRequest) SetConfig(value string) {
	r.Config = value
	r.QueryParams.Set("Config", value)
}
func (r *CreateInstanceRequest) GetConfig() string {
	return r.Config
}
func (r *CreateInstanceRequest) SetChargeType(value string) {
	r.ChargeType = value
	r.QueryParams.Set("ChargeType", value)
}
func (r *CreateInstanceRequest) GetChargeType() string {
	return r.ChargeType
}
func (r *CreateInstanceRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetPeriod() int {
	return r.Period
}
func (r *CreateInstanceRequest) SetToken(value string) {
	r.Token = value
	r.QueryParams.Set("Token", value)
}
func (r *CreateInstanceRequest) GetToken() string {
	return r.Token
}

func (r *CreateInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateInstance")
	r.SetProduct(Product)
}

type CreateInstanceResponse struct {
	InstanceId       string `xml:"InstanceId" json:"InstanceId"`
	InstanceName     string `xml:"InstanceName" json:"InstanceName"`
	ConnectionDomain string `xml:"ConnectionDomain" json:"ConnectionDomain"`
	Port             int    `xml:"Port" json:"Port"`
	UserName         string `xml:"UserName" json:"UserName"`
	InstanceStatus   string `xml:"InstanceStatus" json:"InstanceStatus"`
	RegionId         string `xml:"RegionId" json:"RegionId"`
	Capacity         int    `xml:"Capacity" json:"Capacity"`
	QPS              int    `xml:"QPS" json:"QPS"`
	Bandwidth        int    `xml:"Bandwidth" json:"Bandwidth"`
	Connections      int    `xml:"Connections" json:"Connections"`
	ZoneId           string `xml:"ZoneId" json:"ZoneId"`
	Config           string `xml:"Config" json:"Config"`
	ChargeType       string `xml:"ChargeType" json:"ChargeType"`
	EndTime          string `xml:"EndTime" json:"EndTime"`
}

func CreateInstance(req *CreateInstanceRequest, accessId, accessSecret string) (*CreateInstanceResponse, error) {
	var pResponse CreateInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
