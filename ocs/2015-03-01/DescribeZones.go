package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeZonesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ZoneId               string
}

func (r *DescribeZonesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeZonesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeZonesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeZonesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeZonesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeZonesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeZonesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeZonesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeZonesRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeZonesRequest) GetZoneId() string {
	return r.ZoneId
}

func (r *DescribeZonesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeZones")
	r.SetProduct(Product)
}

type DescribeZonesResponse struct {
	Zones struct {
		OcsZone []struct {
			ZoneId      string `xml:"ZoneId" json:"ZoneId"`
			Name        string `xml:"Name" json:"Name"`
			Description string `xml:"Description" json:"Description"`
		} `xml:"OcsZone" json:"OcsZone"`
	} `xml:"Zones" json:"Zones"`
}

func DescribeZones(req *DescribeZonesRequest, accessId, accessSecret string) (*DescribeZonesResponse, error) {
	var pResponse DescribeZonesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
