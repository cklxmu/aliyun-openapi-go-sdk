package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSQLInjectionInfosRequest struct {
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

func (r *DescribeSQLInjectionInfosRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLInjectionInfosRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSQLInjectionInfosRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSQLInjectionInfosRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSQLInjectionInfosRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLInjectionInfosRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSQLInjectionInfosRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSQLInjectionInfosRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSQLInjectionInfosRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSQLInjectionInfosRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSQLInjectionInfosRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSQLInjectionInfosRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSQLInjectionInfosRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSQLInjectionInfosRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSQLInjectionInfosRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSQLInjectionInfosRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSQLInjectionInfosRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSQLInjectionInfosRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSQLInjectionInfosRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSQLInjectionInfos")
	r.SetProduct(Product)
}

type DescribeSQLInjectionInfosResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLInjectionInfo []struct {
			DBName         string `xml:"DBName" json:"DBName"`
			SQLText        string `xml:"SQLText" json:"SQLText"`
			LatencyTime    string `xml:"LatencyTime" json:"LatencyTime"`
			HostAddress    string `xml:"HostAddress" json:"HostAddress"`
			ExecuteTime    string `xml:"ExecuteTime" json:"ExecuteTime"`
			AccountName    string `xml:"AccountName" json:"AccountName"`
			EffectRowCount string `xml:"EffectRowCount" json:"EffectRowCount"`
		} `xml:"SQLInjectionInfo" json:"SQLInjectionInfo"`
	} `xml:"Items" json:"Items"`
}

func DescribeSQLInjectionInfos(req *DescribeSQLInjectionInfosRequest, accessId, accessSecret string) (*DescribeSQLInjectionInfosResponse, error) {
	var pResponse DescribeSQLInjectionInfosResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
