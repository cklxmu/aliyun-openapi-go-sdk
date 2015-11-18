package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSecurityGroupAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SecurityGroupId      string
	NicType              string
	OwnerAccount         string
	Direction            string
}

func (r *DescribeSecurityGroupAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSecurityGroupAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSecurityGroupAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSecurityGroupAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *DescribeSecurityGroupAttributeRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *DescribeSecurityGroupAttributeRequest) SetNicType(value string) {
	r.NicType = value
	r.QueryParams.Set("NicType", value)
}
func (r *DescribeSecurityGroupAttributeRequest) GetNicType() string {
	return r.NicType
}
func (r *DescribeSecurityGroupAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSecurityGroupAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeSecurityGroupAttributeRequest) SetDirection(value string) {
	r.Direction = value
	r.QueryParams.Set("Direction", value)
}
func (r *DescribeSecurityGroupAttributeRequest) GetDirection() string {
	return r.Direction
}

func (r *DescribeSecurityGroupAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSecurityGroupAttribute")
	r.SetProduct(Product)
}

type DescribeSecurityGroupAttributeResponse struct {
	RegionId          string `xml:"RegionId" json:"RegionId"`
	SecurityGroupId   string `xml:"SecurityGroupId" json:"SecurityGroupId"`
	Description       string `xml:"Description" json:"Description"`
	SecurityGroupName string `xml:"SecurityGroupName" json:"SecurityGroupName"`
	VpcId             string `xml:"VpcId" json:"VpcId"`
	Permissions       struct {
		Permission []struct {
			IpProtocol              string `xml:"IpProtocol" json:"IpProtocol"`
			PortRange               string `xml:"PortRange" json:"PortRange"`
			SourceGroupId           string `xml:"SourceGroupId" json:"SourceGroupId"`
			SourceCidrIp            string `xml:"SourceCidrIp" json:"SourceCidrIp"`
			Policy                  string `xml:"Policy" json:"Policy"`
			NicType                 string `xml:"NicType" json:"NicType"`
			SourceGroupOwnerAccount string `xml:"SourceGroupOwnerAccount" json:"SourceGroupOwnerAccount"`
			DestGroupId             string `xml:"DestGroupId" json:"DestGroupId"`
			DestCidrIp              string `xml:"DestCidrIp" json:"DestCidrIp"`
			DestGroupOwnerAccount   string `xml:"DestGroupOwnerAccount" json:"DestGroupOwnerAccount"`
			Priority                string `xml:"Priority" json:"Priority"`
			Direction               string `xml:"Direction" json:"Direction"`
		} `xml:"Permission" json:"Permission"`
	} `xml:"Permissions" json:"Permissions"`
}

func DescribeSecurityGroupAttribute(req *DescribeSecurityGroupAttributeRequest, accessId, accessSecret string) (*DescribeSecurityGroupAttributeResponse, error) {
	var pResponse DescribeSecurityGroupAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
