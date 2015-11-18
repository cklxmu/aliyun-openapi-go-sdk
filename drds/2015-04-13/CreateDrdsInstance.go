package drds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDrdsInstanceRequest struct {
	RpcRequest
	Description   string
	ZoneId        string
	Type          string
	Quantity      int
	Specification string
	PayType       string
	VpcId         string
	VswitchId     string
}

func (r *CreateDrdsInstanceRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateDrdsInstanceRequest) GetDescription() string {
	return r.Description
}
func (r *CreateDrdsInstanceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateDrdsInstanceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateDrdsInstanceRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *CreateDrdsInstanceRequest) GetType() string {
	return r.Type
}
func (r *CreateDrdsInstanceRequest) SetQuantity(value int) {
	r.Quantity = value
	r.QueryParams.Set("Quantity", strconv.Itoa(value))
}
func (r *CreateDrdsInstanceRequest) GetQuantity() int {
	return r.Quantity
}
func (r *CreateDrdsInstanceRequest) SetSpecification(value string) {
	r.Specification = value
	r.QueryParams.Set("Specification", value)
}
func (r *CreateDrdsInstanceRequest) GetSpecification() string {
	return r.Specification
}
func (r *CreateDrdsInstanceRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *CreateDrdsInstanceRequest) GetPayType() string {
	return r.PayType
}
func (r *CreateDrdsInstanceRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateDrdsInstanceRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateDrdsInstanceRequest) SetVswitchId(value string) {
	r.VswitchId = value
	r.QueryParams.Set("VswitchId", value)
}
func (r *CreateDrdsInstanceRequest) GetVswitchId() string {
	return r.VswitchId
}

func (r *CreateDrdsInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDrdsInstance")
	r.SetProduct(Product)
}

type CreateDrdsInstanceResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		OrderId            int `xml:"OrderId" json:"OrderId"`
		DrdsInstanceIdList struct {
			DrdsInstanceId []string `xml:"DrdsInstanceId" json:"DrdsInstanceId"`
		} `xml:"DrdsInstanceIdList" json:"DrdsInstanceIdList"`
	} `xml:"Data" json:"Data"`
}

func CreateDrdsInstance(req *CreateDrdsInstanceRequest, accessId, accessSecret string) (*CreateDrdsInstanceResponse, error) {
	var pResponse CreateDrdsInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
