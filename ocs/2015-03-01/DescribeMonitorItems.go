package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeMonitorItemsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
}

func (r *DescribeMonitorItemsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeMonitorItemsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeMonitorItemsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeMonitorItemsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeMonitorItemsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeMonitorItemsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeMonitorItemsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeMonitorItemsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeMonitorItemsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMonitorItems")
	r.SetProduct(Product)
}

type DescribeMonitorItemsResponse struct {
	MonitorItems struct {
		OcsMonitorItem []struct {
			MonitorKey string `xml:"MonitorKey" json:"MonitorKey"`
			Unit       string `xml:"Unit" json:"Unit"`
		} `xml:"OcsMonitorItem" json:"OcsMonitorItem"`
	} `xml:"MonitorItems" json:"MonitorItems"`
}

func DescribeMonitorItems(req *DescribeMonitorItemsRequest, accessId, accessSecret string) (*DescribeMonitorItemsResponse, error) {
	var pResponse DescribeMonitorItemsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
