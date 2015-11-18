package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListJobsRequest struct {
	RpcRequest
	IsDesc     bool
	PageNumber int
	PageSize   int
}

func (r *ListJobsRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListJobsRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListJobsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListJobsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListJobsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListJobsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListJobsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListJobs")
	r.SetProduct(Product)
}

type ListJobsResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Jobs       struct {
		JobInfo []struct {
			JobId      int    `xml:"JobId" json:"JobId"`
			JobName    string `xml:"JobName" json:"JobName"`
			JobType    int    `xml:"JobType" json:"JobType"`
			EnvConfig  string `xml:"EnvConfig" json:"EnvConfig"`
			JobFailAct int    `xml:"JobFailAct" json:"JobFailAct"`
		} `xml:"JobInfo" json:"JobInfo"`
	} `xml:"Jobs" json:"Jobs"`
}

func ListJobs(req *ListJobsRequest, accessId, accessSecret string) (*ListJobsResponse, error) {
	var pResponse ListJobsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
