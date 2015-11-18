package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListExecutePlansRequest struct {
	RpcRequest
	Strategy   int
	Status     string
	IsDesc     bool
	PageNumber int
	PageSize   int
}

func (r *ListExecutePlansRequest) SetStrategy(value int) {
	r.Strategy = value
	r.QueryParams.Set("Strategy", strconv.Itoa(value))
}
func (r *ListExecutePlansRequest) GetStrategy() int {
	return r.Strategy
}
func (r *ListExecutePlansRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *ListExecutePlansRequest) GetStatus() string {
	return r.Status
}
func (r *ListExecutePlansRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListExecutePlansRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListExecutePlansRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListExecutePlansRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListExecutePlansRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListExecutePlansRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListExecutePlansRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListExecutePlans")
	r.SetProduct(Product)
}

type ListExecutePlansResponse struct {
	TotalCount   int `xml:"TotalCount" json:"TotalCount"`
	PageNumber   int `xml:"PageNumber" json:"PageNumber"`
	PageSize     int `xml:"PageSize" json:"PageSize"`
	ExecutePlans struct {
		ExecutePlanInfo []struct {
			Id              int    `xml:"Id" json:"Id"`
			Name            string `xml:"Name" json:"Name"`
			LastRunStatus   string `xml:"LastRunStatus" json:"LastRunStatus"`
			RunTime         int    `xml:"RunTime" json:"RunTime"`
			ClusterName     string `xml:"ClusterName" json:"ClusterName"`
			IsCreateCluster bool   `xml:"IsCreateCluster" json:"IsCreateCluster"`
			Stragety        int    `xml:"Stragety" json:"Stragety"`
			Status          string `xml:"Status" json:"Status"`
			TimeInterval    int    `xml:"TimeInterval" json:"TimeInterval"`
			StartTime       string `xml:"StartTime" json:"StartTime"`
			TimeUnit        string `xml:"TimeUnit" json:"TimeUnit"`
		} `xml:"ExecutePlanInfo" json:"ExecutePlanInfo"`
	} `xml:"ExecutePlans" json:"ExecutePlans"`
}

func ListExecutePlans(req *ListExecutePlansRequest, accessId, accessSecret string) (*ListExecutePlansResponse, error) {
	var pResponse ListExecutePlansResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
