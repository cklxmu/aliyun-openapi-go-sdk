package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListExecutePlanExecuteRecordNodesRequest struct {
	RpcRequest
	ExecutePlanExecRecordId int
	IsDesc                  bool
	PageNumber              int
	PageSize                int
}

func (r *ListExecutePlanExecuteRecordNodesRequest) SetExecutePlanExecRecordId(value int) {
	r.ExecutePlanExecRecordId = value
	r.QueryParams.Set("ExecutePlanExecRecordId", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordNodesRequest) GetExecutePlanExecRecordId() int {
	return r.ExecutePlanExecRecordId
}
func (r *ListExecutePlanExecuteRecordNodesRequest) SetIsDesc(value bool) {
	r.IsDesc = value
	r.QueryParams.Set("IsDesc", strconv.FormatBool(value))
}
func (r *ListExecutePlanExecuteRecordNodesRequest) GetIsDesc() bool {
	return r.IsDesc
}
func (r *ListExecutePlanExecuteRecordNodesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordNodesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *ListExecutePlanExecuteRecordNodesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListExecutePlanExecuteRecordNodesRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListExecutePlanExecuteRecordNodesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListExecutePlanExecuteRecordNodes")
	r.SetProduct(Product)
}

type ListExecutePlanExecuteRecordNodesResponse struct {
	TotalCount      int `xml:"TotalCount" json:"TotalCount"`
	PageNumber      int `xml:"PageNumber" json:"PageNumber"`
	PageSize        int `xml:"PageSize" json:"PageSize"`
	executePlanNode struct {
		ExecutePlanNodeInfo []struct {
			WorkNodeId   string `xml:"WorkNodeId" json:"WorkNodeId"`
			WorkNodeName string `xml:"WorkNodeName" json:"WorkNodeName"`
			StartTime    string `xml:"StartTime" json:"StartTime"`
			RunTime      int    `xml:"RunTime" json:"RunTime"`
			JobType      int    `xml:"JobType" json:"JobType"`
			JobId        int    `xml:"JobId" json:"JobId"`
			ClusterId    int    `xml:"ClusterId" json:"ClusterId"`
			Status       int    `xml:"Status" json:"Status"`
		} `xml:"ExecutePlanNodeInfo" json:"ExecutePlanNodeInfo"`
	} `xml:"executePlanNode" json:"executePlanNode"`
}

func ListExecutePlanExecuteRecordNodes(req *ListExecutePlanExecuteRecordNodesRequest, accessId, accessSecret string) (*ListExecutePlanExecuteRecordNodesResponse, error) {
	var pResponse ListExecutePlanExecuteRecordNodesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
