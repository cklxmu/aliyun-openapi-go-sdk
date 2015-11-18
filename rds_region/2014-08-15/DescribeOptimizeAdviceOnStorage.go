package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceOnStorageRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceOnStorageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceOnStorageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnStorageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceOnStorageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceOnStorage")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceOnStorageResponse struct {
	DBInstanceId      string `xml:"DBInstanceId" json:"DBInstanceId"`
	TotalRecordsCount int    `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceOnStorage []struct {
			DBName        string `xml:"DBName" json:"DBName"`
			TableName     string `xml:"TableName" json:"TableName"`
			CurrentEngine string `xml:"CurrentEngine" json:"CurrentEngine"`
			AdviseEngine  string `xml:"AdviseEngine" json:"AdviseEngine"`
		} `xml:"AdviceOnStorage" json:"AdviceOnStorage"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceOnStorage(req *DescribeOptimizeAdviceOnStorageRequest, accessId, accessSecret string) (*DescribeOptimizeAdviceOnStorageResponse, error) {
	var pResponse DescribeOptimizeAdviceOnStorageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
