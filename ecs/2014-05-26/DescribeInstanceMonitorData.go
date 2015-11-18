package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceMonitorDataRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	StartTime            string
	EndTime              string
	Period               int
	OwnerAccount         string
}

func (r *DescribeInstanceMonitorDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceMonitorDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceMonitorDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceMonitorDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceMonitorDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceMonitorDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceMonitorDataRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeInstanceMonitorDataRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeInstanceMonitorDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeInstanceMonitorDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeInstanceMonitorDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeInstanceMonitorDataRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeInstanceMonitorDataRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *DescribeInstanceMonitorDataRequest) GetPeriod() int {
	return r.Period
}
func (r *DescribeInstanceMonitorDataRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceMonitorDataRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceMonitorDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceMonitorData")
	r.SetProduct(Product)
}

type DescribeInstanceMonitorDataResponse struct {
	MonitorData struct {
		InstanceMonitorData []struct {
			InstanceId        string `xml:"InstanceId" json:"InstanceId"`
			CPU               int    `xml:"CPU" json:"CPU"`
			IntranetRX        int    `xml:"IntranetRX" json:"IntranetRX"`
			IntranetTX        int    `xml:"IntranetTX" json:"IntranetTX"`
			IntranetBandwidth int    `xml:"IntranetBandwidth" json:"IntranetBandwidth"`
			InternetRX        int    `xml:"InternetRX" json:"InternetRX"`
			InternetTX        int    `xml:"InternetTX" json:"InternetTX"`
			InternetBandwidth int    `xml:"InternetBandwidth" json:"InternetBandwidth"`
			IOPSRead          int    `xml:"IOPSRead" json:"IOPSRead"`
			IOPSWrite         int    `xml:"IOPSWrite" json:"IOPSWrite"`
			BPSRead           int    `xml:"BPSRead" json:"BPSRead"`
			BPSWrite          int    `xml:"BPSWrite" json:"BPSWrite"`
			TimeStamp         string `xml:"TimeStamp" json:"TimeStamp"`
		} `xml:"InstanceMonitorData" json:"InstanceMonitorData"`
	} `xml:"MonitorData" json:"MonitorData"`
}

func DescribeInstanceMonitorData(req *DescribeInstanceMonitorDataRequest, accessId, accessSecret string) (*DescribeInstanceMonitorDataResponse, error) {
	var pResponse DescribeInstanceMonitorDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
