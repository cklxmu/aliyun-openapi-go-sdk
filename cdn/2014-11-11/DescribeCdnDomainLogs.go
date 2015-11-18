package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCdnDomainLogsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
	LogDay               string
}

func (r *DescribeCdnDomainLogsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainLogsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCdnDomainLogsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCdnDomainLogsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCdnDomainLogsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainLogsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeCdnDomainLogsRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DescribeCdnDomainLogsRequest) GetDomainName() string {
	return r.DomainName
}
func (r *DescribeCdnDomainLogsRequest) SetLogDay(value string) {
	r.LogDay = value
	r.QueryParams.Set("LogDay", value)
}
func (r *DescribeCdnDomainLogsRequest) GetLogDay() string {
	return r.LogDay
}

func (r *DescribeCdnDomainLogsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCdnDomainLogs")
	r.SetProduct(Product)
}

type DescribeCdnDomainLogsResponse struct {
	DomainLogModel struct {
		DomainName       string `xml:"DomainName" json:"DomainName"`
		DomainLogDetails struct {
			DomainLogDetail []struct {
				LogName   string `xml:"LogName" json:"LogName"`
				LogPath   string `xml:"LogPath" json:"LogPath"`
				LogSize   int    `xml:"LogSize" json:"LogSize"`
				StartTime string `xml:"StartTime" json:"StartTime"`
				EndTime   string `xml:"EndTime" json:"EndTime"`
			} `xml:"DomainLogDetail" json:"DomainLogDetail"`
		} `xml:"DomainLogDetails" json:"DomainLogDetails"`
	} `xml:"DomainLogModel" json:"DomainLogModel"`
}

func DescribeCdnDomainLogs(req *DescribeCdnDomainLogsRequest, accessId, accessSecret string) (*DescribeCdnDomainLogsResponse, error) {
	var pResponse DescribeCdnDomainLogsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
