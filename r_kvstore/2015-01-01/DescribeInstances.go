package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceIds          string
	InstanceStatus       string
	ChargeType           string
	PageNumber           int
	PageSize             int
}

func (r *DescribeInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeInstancesRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *DescribeInstancesRequest) GetInstanceIds() string {
	return r.InstanceIds
}
func (r *DescribeInstancesRequest) SetInstanceStatus(value string) {
	r.InstanceStatus = value
	r.QueryParams.Set("InstanceStatus", value)
}
func (r *DescribeInstancesRequest) GetInstanceStatus() string {
	return r.InstanceStatus
}
func (r *DescribeInstancesRequest) SetChargeType(value string) {
	r.ChargeType = value
	r.QueryParams.Set("ChargeType", value)
}
func (r *DescribeInstancesRequest) GetChargeType() string {
	return r.ChargeType
}
func (r *DescribeInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageSize() int {
	return r.PageSize
}

func (r *DescribeInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstances")
	r.SetProduct(Product)
}

type DescribeInstancesResponse struct {
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	Instances  struct {
		KVStoreInstance []struct {
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
			CreateTime       string `xml:"CreateTime" json:"CreateTime"`
			EndTime          string `xml:"EndTime" json:"EndTime"`
		} `xml:"KVStoreInstance" json:"KVStoreInstance"`
	} `xml:"Instances" json:"Instances"`
}

func DescribeInstances(req *DescribeInstancesRequest, accessId, accessSecret string) (*DescribeInstancesResponse, error) {
	var pResponse DescribeInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
