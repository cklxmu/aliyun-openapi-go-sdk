package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeEipAddressesRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	Status                 string
	EipAddress             string
	AllocationId           string
	PageNumber             int
	PageSize               int
	OwnerAccount           string
	Filter_1_Key           string
	Filter_2_Key           string
	Filter_1_Value         string
	Filter_2_Value         string
	LockReason             string
	AssociatedInstanceType string
	AssociatedInstanceId   string
}

func (r *DescribeEipAddressesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeEipAddressesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeEipAddressesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeEipAddressesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeEipAddressesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeEipAddressesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeEipAddressesRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeEipAddressesRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeEipAddressesRequest) SetEipAddress(value string) {
	r.EipAddress = value
	r.QueryParams.Set("EipAddress", value)
}
func (r *DescribeEipAddressesRequest) GetEipAddress() string {
	return r.EipAddress
}
func (r *DescribeEipAddressesRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *DescribeEipAddressesRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *DescribeEipAddressesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeEipAddressesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeEipAddressesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeEipAddressesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeEipAddressesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeEipAddressesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeEipAddressesRequest) SetFilter_1_Key(value string) {
	r.Filter_1_Key = value
	r.QueryParams.Set("Filter.1.Key", value)
}
func (r *DescribeEipAddressesRequest) GetFilter_1_Key() string {
	return r.Filter_1_Key
}
func (r *DescribeEipAddressesRequest) SetFilter_2_Key(value string) {
	r.Filter_2_Key = value
	r.QueryParams.Set("Filter.2.Key", value)
}
func (r *DescribeEipAddressesRequest) GetFilter_2_Key() string {
	return r.Filter_2_Key
}
func (r *DescribeEipAddressesRequest) SetFilter_1_Value(value string) {
	r.Filter_1_Value = value
	r.QueryParams.Set("Filter.1.Value", value)
}
func (r *DescribeEipAddressesRequest) GetFilter_1_Value() string {
	return r.Filter_1_Value
}
func (r *DescribeEipAddressesRequest) SetFilter_2_Value(value string) {
	r.Filter_2_Value = value
	r.QueryParams.Set("Filter.2.Value", value)
}
func (r *DescribeEipAddressesRequest) GetFilter_2_Value() string {
	return r.Filter_2_Value
}
func (r *DescribeEipAddressesRequest) SetLockReason(value string) {
	r.LockReason = value
	r.QueryParams.Set("LockReason", value)
}
func (r *DescribeEipAddressesRequest) GetLockReason() string {
	return r.LockReason
}
func (r *DescribeEipAddressesRequest) SetAssociatedInstanceType(value string) {
	r.AssociatedInstanceType = value
	r.QueryParams.Set("AssociatedInstanceType", value)
}
func (r *DescribeEipAddressesRequest) GetAssociatedInstanceType() string {
	return r.AssociatedInstanceType
}
func (r *DescribeEipAddressesRequest) SetAssociatedInstanceId(value string) {
	r.AssociatedInstanceId = value
	r.QueryParams.Set("AssociatedInstanceId", value)
}
func (r *DescribeEipAddressesRequest) GetAssociatedInstanceId() string {
	return r.AssociatedInstanceId
}

func (r *DescribeEipAddressesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeEipAddresses")
	r.SetProduct(Product)
}

type DescribeEipAddressesResponse struct {
	TotalCount   int `xml:"TotalCount" json:"TotalCount"`
	PageNumber   int `xml:"PageNumber" json:"PageNumber"`
	PageSize     int `xml:"PageSize" json:"PageSize"`
	EipAddresses struct {
		EipAddress []struct {
			RegionId           string `xml:"RegionId" json:"RegionId"`
			IpAddress          string `xml:"IpAddress" json:"IpAddress"`
			AllocationId       string `xml:"AllocationId" json:"AllocationId"`
			Status             string `xml:"Status" json:"Status"`
			InstanceId         string `xml:"InstanceId" json:"InstanceId"`
			Bandwidth          string `xml:"Bandwidth" json:"Bandwidth"`
			InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
			AllocationTime     string `xml:"AllocationTime" json:"AllocationTime"`
			InstanceType       string `xml:"InstanceType" json:"InstanceType"`
			OperationLocks     struct {
				LockReason []struct {
					LockReason string `xml:"LockReason" json:"LockReason"`
				} `xml:"LockReason" json:"LockReason"`
			} `xml:"OperationLocks" json:"OperationLocks"`
		} `xml:"EipAddress" json:"EipAddress"`
	} `xml:"EipAddresses" json:"EipAddresses"`
}

func DescribeEipAddresses(req *DescribeEipAddressesRequest, accessId, accessSecret string) (*DescribeEipAddressesResponse, error) {
	var pResponse DescribeEipAddressesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
