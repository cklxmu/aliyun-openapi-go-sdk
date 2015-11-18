package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListInstanceInfosRequest struct {
	RpcRequest
	JstOwnerId   int
	PageNumber   int
	PageSize     int
	Region       string
	EventType    string
	InstanceName string
	InstanceType string
	InstanceIds  string
}

func (r *ListInstanceInfosRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *ListInstanceInfosRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *ListInstanceInfosRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListInstanceInfosRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListInstanceInfosRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListInstanceInfosRequest) GetPageSize() int {
	return r.PageSize
}
func (r *ListInstanceInfosRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *ListInstanceInfosRequest) GetRegion() string {
	return r.Region
}
func (r *ListInstanceInfosRequest) SetEventType(value string) {
	r.EventType = value
	r.QueryParams.Set("EventType", value)
}
func (r *ListInstanceInfosRequest) GetEventType() string {
	return r.EventType
}
func (r *ListInstanceInfosRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *ListInstanceInfosRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *ListInstanceInfosRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *ListInstanceInfosRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *ListInstanceInfosRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *ListInstanceInfosRequest) GetInstanceIds() string {
	return r.InstanceIds
}

func (r *ListInstanceInfosRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListInstanceInfos")
	r.SetProduct(Product)
}

type ListInstanceInfosResponse struct {
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	InfosList  struct {
		InstanceInfo []struct {
			Region       string `xml:"Region" json:"Region"`
			RegionName   string `xml:"RegionName" json:"RegionName"`
			RegionEnName string `xml:"RegionEnName" json:"RegionEnName"`
			InstanceName string `xml:"InstanceName" json:"InstanceName"`
			InstanceId   string `xml:"InstanceId" json:"InstanceId"`
			Ip           string `xml:"Ip" json:"Ip"`
			InternetIp   string `xml:"InternetIp" json:"InternetIp"`
			IntranetIp   string `xml:"IntranetIp" json:"IntranetIp"`
			Ddos         int    `xml:"Ddos" json:"Ddos"`
			HostEvent    int    `xml:"HostEvent" json:"HostEvent"`
			SecureCheck  int    `xml:"SecureCheck" json:"SecureCheck"`
			AegisStatus  int    `xml:"AegisStatus" json:"AegisStatus"`
			Waf          int    `xml:"Waf" json:"Waf"`
			IsLock       bool   `xml:"IsLock" json:"IsLock"`
			LockType     string `xml:"LockType" json:"LockType"`
			UnLockTimes  int    `xml:"UnLockTimes" json:"UnLockTimes"`
			TriggerTime  string `xml:"TriggerTime" json:"TriggerTime"`
		} `xml:"InstanceInfo" json:"InstanceInfo"`
	} `xml:"InfosList" json:"InfosList"`
}

func ListInstanceInfos(req *ListInstanceInfosRequest, accessId, accessSecret string) (*ListInstanceInfosResponse, error) {
	var pResponse ListInstanceInfosResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
