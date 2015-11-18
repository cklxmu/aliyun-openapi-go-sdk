package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeVpcZoneNosRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	OwnerAccount         string
	Region               string
	ZoneId               string
}

func (r *DescribeVpcZoneNosRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeVpcZoneNosRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeVpcZoneNosRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeVpcZoneNosRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeVpcZoneNosRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeVpcZoneNosRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeVpcZoneNosRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeVpcZoneNosRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeVpcZoneNosRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeVpcZoneNosRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeVpcZoneNosRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *DescribeVpcZoneNosRequest) GetRegion() string {
	return r.Region
}
func (r *DescribeVpcZoneNosRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeVpcZoneNosRequest) GetZoneId() string {
	return r.ZoneId
}

func (r *DescribeVpcZoneNosRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeVpcZoneNos")
	r.SetProduct(Product)
}

type DescribeVpcZoneNosResponse struct {
	Items struct {
		VpcZoneId []struct {
			ZoneId    string `xml:"ZoneId" json:"ZoneId"`
			Region    string `xml:"Region" json:"Region"`
			SubDomain string `xml:"SubDomain" json:"SubDomain"`
		} `xml:"VpcZoneId" json:"VpcZoneId"`
	} `xml:"Items" json:"Items"`
}

func DescribeVpcZoneNos(req *DescribeVpcZoneNosRequest, accessId, accessSecret string) (*DescribeVpcZoneNosResponse, error) {
	var pResponse DescribeVpcZoneNosResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
