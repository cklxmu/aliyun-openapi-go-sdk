package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	OwnerAccount         string
}

func (r *DescribeInstanceAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceAttributeRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeInstanceAttributeRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeInstanceAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceAttribute")
	r.SetProduct(Product)
}

type DescribeInstanceAttributeResponse struct {
	InstanceId              string `xml:"InstanceId" json:"InstanceId"`
	InstanceName            string `xml:"InstanceName" json:"InstanceName"`
	ImageId                 string `xml:"ImageId" json:"ImageId"`
	RegionId                string `xml:"RegionId" json:"RegionId"`
	ZoneId                  string `xml:"ZoneId" json:"ZoneId"`
	ClusterId               string `xml:"ClusterId" json:"ClusterId"`
	InstanceType            string `xml:"InstanceType" json:"InstanceType"`
	HostName                string `xml:"HostName" json:"HostName"`
	Status                  string `xml:"Status" json:"Status"`
	InternetChargeType      string `xml:"InternetChargeType" json:"InternetChargeType"`
	InternetMaxBandwidthIn  int    `xml:"InternetMaxBandwidthIn" json:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int    `xml:"InternetMaxBandwidthOut" json:"InternetMaxBandwidthOut"`
	VlanId                  string `xml:"VlanId" json:"VlanId"`
	SerialNumber            string `xml:"SerialNumber" json:"SerialNumber"`
	CreationTime            string `xml:"CreationTime" json:"CreationTime"`
	Description             string `xml:"Description" json:"Description"`
	InstanceNetworkType     string `xml:"InstanceNetworkType" json:"InstanceNetworkType"`
	IoOptimized             string `xml:"IoOptimized" json:"IoOptimized"`
	InstanceChargeType      string `xml:"InstanceChargeType" json:"InstanceChargeType"`
	ExpiredTime             string `xml:"ExpiredTime" json:"ExpiredTime"`
	SecurityGroupIds        struct {
		SecurityGroupId []string `xml:"SecurityGroupId" json:"SecurityGroupId"`
	} `xml:"SecurityGroupIds" json:"SecurityGroupIds"`
	PublicIpAddress struct {
		IpAddress []string `xml:"IpAddress" json:"IpAddress"`
	} `xml:"PublicIpAddress" json:"PublicIpAddress"`
	InnerIpAddress struct {
		IpAddress []string `xml:"IpAddress" json:"IpAddress"`
	} `xml:"InnerIpAddress" json:"InnerIpAddress"`
	OperationLocks struct {
		LockReason []struct {
			LockReason string `xml:"LockReason" json:"LockReason"`
		} `xml:"LockReason" json:"LockReason"`
	} `xml:"OperationLocks" json:"OperationLocks"`
	VpcAttributes struct {
		VpcId            string `xml:"VpcId" json:"VpcId"`
		VSwitchId        string `xml:"VSwitchId" json:"VSwitchId"`
		NatIpAddress     string `xml:"NatIpAddress" json:"NatIpAddress"`
		PrivateIpAddress struct {
			IpAddress []string `xml:"IpAddress" json:"IpAddress"`
		} `xml:"PrivateIpAddress" json:"PrivateIpAddress"`
	} `xml:"VpcAttributes" json:"VpcAttributes"`
	EipAddress struct {
		AllocationId       string `xml:"AllocationId" json:"AllocationId"`
		IpAddress          string `xml:"IpAddress" json:"IpAddress"`
		Bandwidth          int    `xml:"Bandwidth" json:"Bandwidth"`
		InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
	} `xml:"EipAddress" json:"EipAddress"`
}

func DescribeInstanceAttribute(req *DescribeInstanceAttributeRequest, accessId, accessSecret string) (*DescribeInstanceAttributeResponse, error) {
	var pResponse DescribeInstanceAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
