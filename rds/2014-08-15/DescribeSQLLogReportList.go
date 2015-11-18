package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSQLLogReportListRequest struct {
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

func (r *DescribeSQLLogReportListRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportListRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSQLLogReportListRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSQLLogReportListRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSQLLogReportListRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportListRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSQLLogReportListRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSQLLogReportListRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSQLLogReportListRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSQLLogReportListRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSQLLogReportListRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSQLLogReportListRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSQLLogReportListRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportListRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSQLLogReportListRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportListRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSQLLogReportListRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSQLLogReportListRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSQLLogReportListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSQLLogReportList")
	r.SetProduct(Product)
}

type DescribeSQLLogReportListResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		Item []struct {
			ReportTime       string `xml:"ReportTime" json:"ReportTime"`
			LatencyTopNItems struct {
				LatencyTopNItem []struct {
					SQLText         string `xml:"SQLText" json:"SQLText"`
					AvgLatency      int    `xml:"AvgLatency" json:"AvgLatency"`
					SQLExecuteTimes int    `xml:"SQLExecuteTimes" json:"SQLExecuteTimes"`
				} `xml:"LatencyTopNItem" json:"LatencyTopNItem"`
			} `xml:"LatencyTopNItems" json:"LatencyTopNItems"`
			QPSTopNItems struct {
				QPSTopNItem []struct {
					SQLText         string `xml:"SQLText" json:"SQLText"`
					SQLExecuteTimes int    `xml:"SQLExecuteTimes" json:"SQLExecuteTimes"`
				} `xml:"QPSTopNItem" json:"QPSTopNItem"`
			} `xml:"QPSTopNItems" json:"QPSTopNItems"`
		} `xml:"Item" json:"Item"`
	} `xml:"Items" json:"Items"`
}

func DescribeSQLLogReportList(req *DescribeSQLLogReportListRequest, accessId, accessSecret string) (*DescribeSQLLogReportListResponse, error) {
	var pResponse DescribeSQLLogReportListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
