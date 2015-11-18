package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListExecutePlanExecuteRecordsRequest struct {
	RpcRequest
	ExecutePlanId int
	IsDesc        bool
	Status        string
	PageNumber    int
	PageSize      int
}

func (r *ListExecutePlanExecuteRecordsRequest) SetExecutePlanId(value int) {
	r.ExecutePlanId = value
	r.QueryParams.Set("ExecutePlanId", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordsRequest) GetExecutePlanId() int {
	return r.ExecutePlanId
}
func (r *ListExecutePlanExecuteRecordsRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListExecutePlanExecuteRecordsRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListExecutePlanExecuteRecordsRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *ListExecutePlanExecuteRecordsRequest) GetStatus() string {
	return r.Status
}
func (r *ListExecutePlanExecuteRecordsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListExecutePlanExecuteRecordsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListExecutePlanExecuteRecordsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListExecutePlanExecuteRecords")
	r.SetProduct(Product)
}

type ListExecutePlanExecuteRecordsResponse struct {
	TotalCount            int `xml:"TotalCount" json:"TotalCount"`
	PageNumber            int `xml:"PageNumber" json:"PageNumber"`
	PageSize              int `xml:"PageSize" json:"PageSize"`
	ExecutePlanExecRecord struct {
		ExecutePlanRecordInfo []struct {
			Id              int    `xml:"Id" json:"Id"`
			ExecutePlanId   int    `xml:"ExecutePlanId" json:"ExecutePlanId"`
			ExecutePlanName string `xml:"ExecutePlanName" json:"ExecutePlanName"`
			StartTime       string `xml:"StartTime" json:"StartTime"`
			RunTime         int    `xml:"RunTime" json:"RunTime"`
			ClusterId       int    `xml:"ClusterId" json:"ClusterId"`
			ClusterName     string `xml:"ClusterName" json:"ClusterName"`
			ClusterType     int    `xml:"ClusterType" json:"ClusterType"`
			Status          int    `xml:"Status" json:"Status"`
			LogEnable       bool   `xml:"LogEnable" json:"LogEnable"`
			LogPath         string `xml:"LogPath" json:"LogPath"`
		} `xml:"ExecutePlanRecordInfo" json:"ExecutePlanRecordInfo"`
	} `xml:"ExecutePlanExecRecord" json:"ExecutePlanExecRecord"`
}

func ListExecutePlanExecuteRecords(req *ListExecutePlanExecuteRecordsRequest, accessId, accessSecret string) (*ListExecutePlanExecuteRecordsResponse, error) {
	var pResponse ListExecutePlanExecuteRecordsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
