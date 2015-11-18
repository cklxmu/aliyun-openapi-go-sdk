package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsClusterListRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Cluster      string
}

func (r *OnsClusterListRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsClusterListRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsClusterListRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsClusterListRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsClusterListRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsClusterListRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsClusterListRequest) SetCluster(value string) {
	r.Cluster = value
	r.QueryParams.Set("Cluster", value)
}
func (r *OnsClusterListRequest) GetCluster() string {
	return r.Cluster
}

func (r *OnsClusterListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsClusterList")
	r.SetProduct(Product)
}

type OnsClusterListResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		ClusterInfoDataDo []struct {
			ClusterName    string `xml:"ClusterName" json:"ClusterName"`
			BrokerInfoList struct {
				BrokerInfoDataDo []struct {
					ClusterName   string  `xml:"ClusterName" json:"ClusterName"`
					BrokerName    string  `xml:"BrokerName" json:"BrokerName"`
					BrokerId      int     `xml:"BrokerId" json:"BrokerId"`
					BrokerAddr    string  `xml:"BrokerAddr" json:"BrokerAddr"`
					BrokerIp      string  `xml:"BrokerIp" json:"BrokerIp"`
					Version       string  `xml:"Version" json:"Version"`
					InTPS         float32 `xml:"InTPS" json:"InTPS"`
					OutTPS        float32 `xml:"OutTPS" json:"OutTPS"`
					InTotalYest   float32 `xml:"InTotalYest" json:"InTotalYest"`
					OutTotalYest  float32 `xml:"OutTotalYest" json:"OutTotalYest"`
					InTotalToday  float32 `xml:"InTotalToday" json:"InTotalToday"`
					OutTotalToday float32 `xml:"OutTotalToday" json:"OutTotalToday"`
				} `xml:"BrokerInfoDataDo" json:"BrokerInfoDataDo"`
			} `xml:"BrokerInfoList" json:"BrokerInfoList"`
		} `xml:"ClusterInfoDataDo" json:"ClusterInfoDataDo"`
	} `xml:"Data" json:"Data"`
}

func OnsClusterList(req *OnsClusterListRequest, accessId, accessSecret string) (*OnsClusterListResponse, error) {
	var pResponse OnsClusterListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
