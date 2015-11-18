package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeVpcsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeVpcsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeVpcsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeVpcsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeVpcsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeVpcsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeVpcsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeVpcsRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeVpcsRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeVpcsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeVpcsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeVpcsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeVpcsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeVpcsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeVpcsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeVpcsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeVpcs")
	r.SetProduct(Product)
}

type DescribeVpcsResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Vpcs       struct {
		Vpc []struct {
			VpcId        string `xml:"VpcId" json:"VpcId"`
			RegionId     string `xml:"RegionId" json:"RegionId"`
			Status       string `xml:"Status" json:"Status"`
			VpcName      string `xml:"VpcName" json:"VpcName"`
			CreationTime string `xml:"CreationTime" json:"CreationTime"`
			CidrBlock    string `xml:"CidrBlock" json:"CidrBlock"`
			VRouterId    string `xml:"VRouterId" json:"VRouterId"`
			Description  string `xml:"Description" json:"Description"`
			VSwitchIds   struct {
				VSwitchId []string `xml:"VSwitchId" json:"VSwitchId"`
			} `xml:"VSwitchIds" json:"VSwitchIds"`
			UserCidrs struct {
				UserCidr []string `xml:"UserCidr" json:"UserCidr"`
			} `xml:"UserCidrs" json:"UserCidrs"`
		} `xml:"Vpc" json:"Vpc"`
	} `xml:"Vpcs" json:"Vpcs"`
}

func DescribeVpcs(req *DescribeVpcsRequest, accessId, accessSecret string) (*DescribeVpcsResponse, error) {
	var pResponse DescribeVpcsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
