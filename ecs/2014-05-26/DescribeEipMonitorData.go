package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeEipMonitorDataRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AllocationId         string
	StartTime            string
	EndTime              string
	Period               int
	OwnerAccount         string
}

func (r *DescribeEipMonitorDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeEipMonitorDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeEipMonitorDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeEipMonitorDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeEipMonitorDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeEipMonitorDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeEipMonitorDataRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *DescribeEipMonitorDataRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *DescribeEipMonitorDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeEipMonitorDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeEipMonitorDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeEipMonitorDataRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeEipMonitorDataRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *DescribeEipMonitorDataRequest) GetPeriod() int {
	return r.Period
}
func (r *DescribeEipMonitorDataRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeEipMonitorDataRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeEipMonitorDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeEipMonitorData")
	r.SetProduct(Product)
}

type DescribeEipMonitorDataResponse struct {
	EipMonitorDatas struct {
		EipMonitorData []struct {
			EipRX        int    `xml:"EipRX" json:"EipRX"`
			EipTX        int    `xml:"EipTX" json:"EipTX"`
			EipFlow      int    `xml:"EipFlow" json:"EipFlow"`
			EipBandwidth int    `xml:"EipBandwidth" json:"EipBandwidth"`
			EipPackets   int    `xml:"EipPackets" json:"EipPackets"`
			TimeStamp    string `xml:"TimeStamp" json:"TimeStamp"`
		} `xml:"EipMonitorData" json:"EipMonitorData"`
	} `xml:"EipMonitorDatas" json:"EipMonitorDatas"`
}

func DescribeEipMonitorData(req *DescribeEipMonitorDataRequest, accessId, accessSecret string) (*DescribeEipMonitorDataResponse, error) {
	var pResponse DescribeEipMonitorDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
