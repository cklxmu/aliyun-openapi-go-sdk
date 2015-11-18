package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeAutoSnapshotPolicyRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeAutoSnapshotPolicyRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeAutoSnapshotPolicyRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeAutoSnapshotPolicyRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeAutoSnapshotPolicyRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeAutoSnapshotPolicyRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeAutoSnapshotPolicyRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeAutoSnapshotPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeAutoSnapshotPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeAutoSnapshotPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeAutoSnapshotPolicy")
	r.SetProduct(Product)
}

type DescribeAutoSnapshotPolicyResponse struct {
	AutoSnapshotOccupation int `xml:"AutoSnapshotOccupation" json:"AutoSnapshotOccupation"`
	AutoSnapshotPolicy     struct {
		SystemDiskPolicyEnabled           string `xml:"SystemDiskPolicyEnabled" json:"SystemDiskPolicyEnabled"`
		SystemDiskPolicyTimePeriod        string `xml:"SystemDiskPolicyTimePeriod" json:"SystemDiskPolicyTimePeriod"`
		SystemDiskPolicyRetentionDays     string `xml:"SystemDiskPolicyRetentionDays" json:"SystemDiskPolicyRetentionDays"`
		SystemDiskPolicyRetentionLastWeek string `xml:"SystemDiskPolicyRetentionLastWeek" json:"SystemDiskPolicyRetentionLastWeek"`
		DataDiskPolicyEnabled             string `xml:"DataDiskPolicyEnabled" json:"DataDiskPolicyEnabled"`
		DataDiskPolicyTimePeriod          string `xml:"DataDiskPolicyTimePeriod" json:"DataDiskPolicyTimePeriod"`
		DataDiskPolicyRetentionDays       string `xml:"DataDiskPolicyRetentionDays" json:"DataDiskPolicyRetentionDays"`
		DataDiskPolicyRetentionLastWeek   string `xml:"DataDiskPolicyRetentionLastWeek" json:"DataDiskPolicyRetentionLastWeek"`
	} `xml:"AutoSnapshotPolicy" json:"AutoSnapshotPolicy"`
	AutoSnapshotExcutionStatus struct {
		SystemDiskExcutionStatus string `xml:"SystemDiskExcutionStatus" json:"SystemDiskExcutionStatus"`
		DataDiskExcutionStatus   string `xml:"DataDiskExcutionStatus" json:"DataDiskExcutionStatus"`
	} `xml:"AutoSnapshotExcutionStatus" json:"AutoSnapshotExcutionStatus"`
}

func DescribeAutoSnapshotPolicy(req *DescribeAutoSnapshotPolicyRequest, accessId, accessSecret string) (*DescribeAutoSnapshotPolicyResponse, error) {
	var pResponse DescribeAutoSnapshotPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
