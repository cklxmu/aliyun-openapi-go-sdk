package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeHistoryMonitorValuesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	StartTime            string
	EndTime              string
	IntervalForHistory   string
	MonitorKeys          string
}

func (r *DescribeHistoryMonitorValuesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeHistoryMonitorValuesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeHistoryMonitorValuesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeHistoryMonitorValuesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeHistoryMonitorValuesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeHistoryMonitorValuesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeHistoryMonitorValuesRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeHistoryMonitorValuesRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeHistoryMonitorValuesRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeHistoryMonitorValuesRequest) SetIntervalForHistory(value string) {
	r.IntervalForHistory = value
	r.QueryParams.Set("IntervalForHistory", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetIntervalForHistory() string {
	return r.IntervalForHistory
}
func (r *DescribeHistoryMonitorValuesRequest) SetMonitorKeys(value string) {
	r.MonitorKeys = value
	r.QueryParams.Set("MonitorKeys", value)
}
func (r *DescribeHistoryMonitorValuesRequest) GetMonitorKeys() string {
	return r.MonitorKeys
}

func (r *DescribeHistoryMonitorValuesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeHistoryMonitorValues")
	r.SetProduct(Product)
}

type DescribeHistoryMonitorValuesResponse struct {
	MonitorHistory string `xml:"MonitorHistory" json:"MonitorHistory"`
}

func DescribeHistoryMonitorValues(req *DescribeHistoryMonitorValuesRequest, accessId, accessSecret string) (*DescribeHistoryMonitorValuesResponse, error) {
	var pResponse DescribeHistoryMonitorValuesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
