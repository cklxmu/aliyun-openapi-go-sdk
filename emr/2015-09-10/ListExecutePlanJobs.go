package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListExecutePlanJobsRequest struct {
	RpcRequest
	ExecutePlanExecuteNodeId string
	ExecutePlanExecRecordId  string
	IsDesc                   bool
	PageNumber               int
	PageSize                 int
}

func (r *ListExecutePlanJobsRequest) SetExecutePlanExecuteNodeId(value string) {
	r.ExecutePlanExecuteNodeId = value
	r.QueryParams.Set("ExecutePlanExecuteNodeId", value)
}
func (r *ListExecutePlanJobsRequest) GetExecutePlanExecuteNodeId() string {
	return r.ExecutePlanExecuteNodeId
}
func (r *ListExecutePlanJobsRequest) SetExecutePlanExecRecordId(value string) {
	r.ExecutePlanExecRecordId = value
	r.QueryParams.Set("ExecutePlanExecRecordId", value)
}
func (r *ListExecutePlanJobsRequest) GetExecutePlanExecRecordId() string {
	return r.ExecutePlanExecRecordId
}
func (r *ListExecutePlanJobsRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListExecutePlanJobsRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListExecutePlanJobsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListExecutePlanJobsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListExecutePlanJobsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListExecutePlanJobsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListExecutePlanJobsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListExecutePlanJobs")
	r.SetProduct(Product)
}

type ListExecutePlanJobsResponse struct {
	TotalCount      int `xml:"TotalCount" json:"TotalCount"`
	PageNumber      int `xml:"PageNumber" json:"PageNumber"`
	PageSize        int `xml:"PageSize" json:"PageSize"`
	ExecutePlanJobs struct {
		ExecutePlanJobInfo []struct {
			id            string `xml:"id" json:"id"`
			isMaster      string `xml:"isMaster" json:"isMaster"`
			LastRunStatus string `xml:"LastRunStatus" json:"LastRunStatus"`
			RunTime       string `xml:"RunTime" json:"RunTime"`
			ClusterName   string `xml:"ClusterName" json:"ClusterName"`
			Status        string `xml:"Status" json:"Status"`
		} `xml:"ExecutePlanJobInfo" json:"ExecutePlanJobInfo"`
	} `xml:"ExecutePlanJobs" json:"ExecutePlanJobs"`
}

func ListExecutePlanJobs(req *ListExecutePlanJobsRequest, accessId, accessSecret string) (*ListExecutePlanJobsResponse, error) {
	var pResponse ListExecutePlanJobsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
