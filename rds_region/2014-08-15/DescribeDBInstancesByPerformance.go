package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstancesByPerformanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	SortMethod           string
	SortKey              string
	Tag_1_key            string
	Tag_2_key            string
	Tag_3_key            string
	Tag_4_key            string
	Tag_5_key            string
	Tag_1_value          string
	Tag_2_value          string
	Tag_3_value          string
	Tag_4_value          string
	Tag_5_value          string
	OwnerAccount         string
}

func (r *DescribeDBInstancesByPerformanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByPerformanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstancesByPerformanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstancesByPerformanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByPerformanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstancesByPerformanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstancesByPerformanceRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetproxyId() string {
	return r.proxyId
}
func (r *DescribeDBInstancesByPerformanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstancesByPerformanceRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByPerformanceRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeDBInstancesByPerformanceRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByPerformanceRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeDBInstancesByPerformanceRequest) SetSortMethod(value string) {
	r.SortMethod = value
	r.QueryParams.Set("SortMethod", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetSortMethod() string {
	return r.SortMethod
}
func (r *DescribeDBInstancesByPerformanceRequest) SetSortKey(value string) {
	r.SortKey = value
	r.QueryParams.Set("SortKey", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetSortKey() string {
	return r.SortKey
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_1_key(value string) {
	r.Tag_1_key = value
	r.QueryParams.Set("Tag.1.key", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_1_key() string {
	return r.Tag_1_key
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_2_key(value string) {
	r.Tag_2_key = value
	r.QueryParams.Set("Tag.2.key", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_2_key() string {
	return r.Tag_2_key
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_3_key(value string) {
	r.Tag_3_key = value
	r.QueryParams.Set("Tag.3.key", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_3_key() string {
	return r.Tag_3_key
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_4_key(value string) {
	r.Tag_4_key = value
	r.QueryParams.Set("Tag.4.key", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_4_key() string {
	return r.Tag_4_key
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_5_key(value string) {
	r.Tag_5_key = value
	r.QueryParams.Set("Tag.5.key", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_5_key() string {
	return r.Tag_5_key
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_1_value(value string) {
	r.Tag_1_value = value
	r.QueryParams.Set("Tag.1.value", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_1_value() string {
	return r.Tag_1_value
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_2_value(value string) {
	r.Tag_2_value = value
	r.QueryParams.Set("Tag.2.value", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_2_value() string {
	return r.Tag_2_value
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_3_value(value string) {
	r.Tag_3_value = value
	r.QueryParams.Set("Tag.3.value", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_3_value() string {
	return r.Tag_3_value
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_4_value(value string) {
	r.Tag_4_value = value
	r.QueryParams.Set("Tag.4.value", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_4_value() string {
	return r.Tag_4_value
}
func (r *DescribeDBInstancesByPerformanceRequest) SetTag_5_value(value string) {
	r.Tag_5_value = value
	r.QueryParams.Set("Tag.5.value", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetTag_5_value() string {
	return r.Tag_5_value
}
func (r *DescribeDBInstancesByPerformanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstancesByPerformanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstancesByPerformanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstancesByPerformance")
	r.SetProduct(Product)
}

type DescribeDBInstancesByPerformanceResponse struct {
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		DBInstancePerformance []struct {
			CPUUsage              string `xml:"CPUUsage" json:"CPUUsage"`
			IOPSUsage             string `xml:"IOPSUsage" json:"IOPSUsage"`
			DiskUsage             string `xml:"DiskUsage" json:"DiskUsage"`
			SessionUsage          string `xml:"SessionUsage" json:"SessionUsage"`
			DBInstanceId          string `xml:"DBInstanceId" json:"DBInstanceId"`
			DBInstanceDescription string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
		} `xml:"DBInstancePerformance" json:"DBInstancePerformance"`
	} `xml:"Items" json:"Items"`
}

func DescribeDBInstancesByPerformance(req *DescribeDBInstancesByPerformanceRequest, accessId, accessSecret string) (*DescribeDBInstancesByPerformanceResponse, error) {
	var pResponse DescribeDBInstancesByPerformanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
