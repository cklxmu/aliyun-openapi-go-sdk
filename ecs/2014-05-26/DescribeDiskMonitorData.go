package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDiskMonitorDataRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	StartTime            string
	EndTime              string
	Period               int
	OwnerAccount         string
}

func (r *DescribeDiskMonitorDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDiskMonitorDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDiskMonitorDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDiskMonitorDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDiskMonitorDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDiskMonitorDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDiskMonitorDataRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *DescribeDiskMonitorDataRequest) GetDiskId() string {
	return r.DiskId
}
func (r *DescribeDiskMonitorDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeDiskMonitorDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeDiskMonitorDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeDiskMonitorDataRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeDiskMonitorDataRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *DescribeDiskMonitorDataRequest) GetPeriod() int {
	return r.Period
}
func (r *DescribeDiskMonitorDataRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDiskMonitorDataRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDiskMonitorDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDiskMonitorData")
	r.SetProduct(Product)
}

type DescribeDiskMonitorDataResponse struct {
	TotalCount  int `xml:"TotalCount" json:"TotalCount"`
	MonitorData struct {
		DiskMonitorData []struct {
			DiskId    string `xml:"DiskId" json:"DiskId"`
			IOPSRead  int    `xml:"IOPSRead" json:"IOPSRead"`
			IOPSWrite int    `xml:"IOPSWrite" json:"IOPSWrite"`
			IOPSTotal int    `xml:"IOPSTotal" json:"IOPSTotal"`
			BPSRead   int    `xml:"BPSRead" json:"BPSRead"`
			BPSWrite  int    `xml:"BPSWrite" json:"BPSWrite"`
			BPSTotal  int    `xml:"BPSTotal" json:"BPSTotal"`
			TimeStamp string `xml:"TimeStamp" json:"TimeStamp"`
		} `xml:"DiskMonitorData" json:"DiskMonitorData"`
	} `xml:"MonitorData" json:"MonitorData"`
}

func DescribeDiskMonitorData(req *DescribeDiskMonitorDataRequest, accessId, accessSecret string) (*DescribeDiskMonitorDataResponse, error) {
	var pResponse DescribeDiskMonitorDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
