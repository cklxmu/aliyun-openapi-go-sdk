package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeMonitorValuesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceIds          string
	MonitorKeys          string
}

func (r *DescribeMonitorValuesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeMonitorValuesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeMonitorValuesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeMonitorValuesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeMonitorValuesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeMonitorValuesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeMonitorValuesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeMonitorValuesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeMonitorValuesRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *DescribeMonitorValuesRequest) GetInstanceIds() string {
	return r.InstanceIds
}
func (r *DescribeMonitorValuesRequest) SetMonitorKeys(value string) {
	r.MonitorKeys = value
	r.QueryParams.Set("MonitorKeys", value)
}
func (r *DescribeMonitorValuesRequest) GetMonitorKeys() string {
	return r.MonitorKeys
}

func (r *DescribeMonitorValuesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMonitorValues")
	r.SetProduct(Product)
}

type DescribeMonitorValuesResponse struct {
	Date        string `xml:"Date" json:"Date"`
	InstanceIds struct {
		OcsInstanceMonitor []struct {
			InstanceId  string `xml:"InstanceId" json:"InstanceId"`
			MonitorKeys struct {
				OcsMonitorKey []struct {
					MonitorKey string `xml:"MonitorKey" json:"MonitorKey"`
					Value      string `xml:"Value" json:"Value"`
					Unit       string `xml:"Unit" json:"Unit"`
				} `xml:"OcsMonitorKey" json:"OcsMonitorKey"`
			} `xml:"MonitorKeys" json:"MonitorKeys"`
		} `xml:"OcsInstanceMonitor" json:"OcsInstanceMonitor"`
	} `xml:"InstanceIds" json:"InstanceIds"`
}

func DescribeMonitorValues(req *DescribeMonitorValuesRequest, accessId, accessSecret string) (*DescribeMonitorValuesResponse, error) {
	var pResponse DescribeMonitorValuesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
