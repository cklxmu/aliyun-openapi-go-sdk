package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListClustersRequest struct {
	RpcRequest
	ClusterType string
	PayType     string
	Status      string
	IsDesc      bool
	PageNumber  int
	PageSize    int
}

func (r *ListClustersRequest) SetClusterType(value string) {
	r.ClusterType = value
	r.QueryParams.Set("ClusterType", value)
}
func (r *ListClustersRequest) GetClusterType() string {
	return r.ClusterType
}
func (r *ListClustersRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *ListClustersRequest) GetPayType() string {
	return r.PayType
}
func (r *ListClustersRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *ListClustersRequest) GetStatus() string {
	return r.Status
}
func (r *ListClustersRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListClustersRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListClustersRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListClustersRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListClustersRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListClustersRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListClustersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListClusters")
	r.SetProduct(Product)
}

type ListClustersResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Clusters   struct {
		ClusterInfo []struct {
			Id          int    `xml:"Id" json:"Id"`
			Name        string `xml:"Name" json:"Name"`
			PayType     int    `xml:"PayType" json:"PayType"`
			Type        int    `xml:"Type" json:"Type"`
			CreateTime  string `xml:"CreateTime" json:"CreateTime"`
			RunningTime string `xml:"RunningTime" json:"RunningTime"`
			Status      int    `xml:"Status" json:"Status"`
			FailReason  string `xml:"FailReason" json:"FailReason"`
		} `xml:"ClusterInfo" json:"ClusterInfo"`
	} `xml:"Clusters" json:"Clusters"`
}

func ListClusters(req *ListClustersRequest, accessId, accessSecret string) (*ListClustersResponse, error) {
	var pResponse ListClustersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
