package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCdnMonitorDataRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
	StartTime            string
	EndTime              string
}

func (r *DescribeCdnMonitorDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnMonitorDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCdnMonitorDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCdnMonitorDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCdnMonitorDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnMonitorDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeCdnMonitorDataRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DescribeCdnMonitorDataRequest) GetDomainName() string {
	return r.DomainName
}
func (r *DescribeCdnMonitorDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeCdnMonitorDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeCdnMonitorDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeCdnMonitorDataRequest) GetEndTime() string {
	return r.EndTime
}

func (r *DescribeCdnMonitorDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCdnMonitorData")
	r.SetProduct(Product)
}

type DescribeCdnMonitorDataResponse struct {
	DomainName      string `xml:"DomainName" json:"DomainName"`
	MonitorInterval int    `xml:"MonitorInterval" json:"MonitorInterval"`
	StartTime       string `xml:"StartTime" json:"StartTime"`
	EndTime         string `xml:"EndTime" json:"EndTime"`
	MonitorDatas    struct {
		CDNMonitorData []struct {
			TimeStamp         string `xml:"TimeStamp" json:"TimeStamp"`
			QueryPerSecond    string `xml:"QueryPerSecond" json:"QueryPerSecond"`
			BytesPerSecond    string `xml:"BytesPerSecond" json:"BytesPerSecond"`
			BytesHitRate      string `xml:"BytesHitRate" json:"BytesHitRate"`
			RequestHitRate    string `xml:"RequestHitRate" json:"RequestHitRate"`
			AverageObjectSize string `xml:"AverageObjectSize" json:"AverageObjectSize"`
		} `xml:"CDNMonitorData" json:"CDNMonitorData"`
	} `xml:"MonitorDatas" json:"MonitorDatas"`
}

func DescribeCdnMonitorData(req *DescribeCdnMonitorDataRequest, accessId, accessSecret string) (*DescribeCdnMonitorDataResponse, error) {
	var pResponse DescribeCdnMonitorDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
