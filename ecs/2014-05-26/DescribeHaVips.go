package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeHaVipsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	PageNumber           int
	PageSize             int
	Filter               string
}

func (r *DescribeHaVipsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeHaVipsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeHaVipsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeHaVipsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeHaVipsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeHaVipsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeHaVipsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeHaVipsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeHaVipsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeHaVipsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeHaVipsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeHaVipsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeHaVipsRequest) SetFilter(value string) {
	r.Filter = value
	r.QueryParams.Set("Filter", value)
}
func (r *DescribeHaVipsRequest) GetFilter() string {
	return r.Filter
}

func (r *DescribeHaVipsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeHaVips")
	r.SetProduct(Product)
}

type DescribeHaVipsResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	HaVips     struct {
		HaVip []struct {
			HaVipId             string `xml:"HaVipId" json:"HaVipId"`
			RegionId            string `xml:"RegionId" json:"RegionId"`
			VpcId               string `xml:"VpcId" json:"VpcId"`
			VSwitchId           string `xml:"VSwitchId" json:"VSwitchId"`
			IpAddress           string `xml:"IpAddress" json:"IpAddress"`
			Status              string `xml:"Status" json:"Status"`
			MasterInstanceId    string `xml:"MasterInstanceId" json:"MasterInstanceId"`
			Description         string `xml:"Description" json:"Description"`
			CreateTime          string `xml:"CreateTime" json:"CreateTime"`
			AssociatedInstances struct {
				associatedInstance []string `xml:"associatedInstance" json:"associatedInstance"`
			} `xml:"AssociatedInstances" json:"AssociatedInstances"`
			AssociatedEipAddresses struct {
				associatedEipAddresse []string `xml:"associatedEipAddresse" json:"associatedEipAddresse"`
			} `xml:"AssociatedEipAddresses" json:"AssociatedEipAddresses"`
		} `xml:"HaVip" json:"HaVip"`
	} `xml:"HaVips" json:"HaVips"`
}

func DescribeHaVips(req *DescribeHaVipsRequest, accessId, accessSecret string) (*DescribeHaVipsResponse, error) {
	var pResponse DescribeHaVipsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
