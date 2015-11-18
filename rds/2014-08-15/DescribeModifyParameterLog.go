package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeModifyParameterLogRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeModifyParameterLogRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeModifyParameterLogRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeModifyParameterLogRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeModifyParameterLogRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeModifyParameterLogRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeModifyParameterLogRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeModifyParameterLogRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeModifyParameterLogRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeModifyParameterLogRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeModifyParameterLogRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeModifyParameterLogRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeModifyParameterLogRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeModifyParameterLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeModifyParameterLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeModifyParameterLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeModifyParameterLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeModifyParameterLogRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeModifyParameterLogRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeModifyParameterLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeModifyParameterLog")
	r.SetProduct(Product)
}

type DescribeModifyParameterLogResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	EngineVersion    string `xml:"EngineVersion" json:"EngineVersion"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		ParameterChangeLog []struct {
			ModifyTime        string `xml:"ModifyTime" json:"ModifyTime"`
			OldParameterValue string `xml:"OldParameterValue" json:"OldParameterValue"`
			NewParameterValue string `xml:"NewParameterValue" json:"NewParameterValue"`
			ParameterName     string `xml:"ParameterName" json:"ParameterName"`
			Status            string `xml:"Status" json:"Status"`
		} `xml:"ParameterChangeLog" json:"ParameterChangeLog"`
	} `xml:"Items" json:"Items"`
}

func DescribeModifyParameterLog(req *DescribeModifyParameterLogRequest, accessId, accessSecret string) (*DescribeModifyParameterLogResponse, error) {
	var pResponse DescribeModifyParameterLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
