package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOperationLogsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	SearchKey            string
	DBInstanceUseType    string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOperationLogsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOperationLogsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOperationLogsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOperationLogsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOperationLogsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOperationLogsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOperationLogsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOperationLogsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOperationLogsRequest) SetSearchKey(value string) {
	r.SearchKey = value
	r.QueryParams.Set("SearchKey", value)
}
func (r *DescribeOperationLogsRequest) GetSearchKey() string {
	return r.SearchKey
}
func (r *DescribeOperationLogsRequest) SetDBInstanceUseType(value string) {
	r.DBInstanceUseType = value
	r.QueryParams.Set("DBInstanceUseType", value)
}
func (r *DescribeOperationLogsRequest) GetDBInstanceUseType() string {
	return r.DBInstanceUseType
}
func (r *DescribeOperationLogsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeOperationLogsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeOperationLogsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeOperationLogsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeOperationLogsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOperationLogsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOperationLogsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOperationLogsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOperationLogsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOperationLogsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOperationLogsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOperationLogs")
	r.SetProduct(Product)
}

type DescribeOperationLogsResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		Operationlog []struct {
			DBInstanceDescription string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
			DBInstanceId          string `xml:"DBInstanceId" json:"DBInstanceId"`
			OperationSource       string `xml:"OperationSource" json:"OperationSource"`
			OperationItem         string `xml:"OperationItem" json:"OperationItem"`
			executionTime         string `xml:"executionTime" json:"executionTime"`
		} `xml:"Operationlog" json:"Operationlog"`
	} `xml:"Items" json:"Items"`
}

func DescribeOperationLogs(req *DescribeOperationLogsRequest, accessId, accessSecret string) (*DescribeOperationLogsResponse, error) {
	var pResponse DescribeOperationLogsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
