package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceOnMissIndexRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnMissIndexRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceOnMissIndexRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceOnMissIndex")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceOnMissIndexResponse struct {
	DBInstanceId      string `xml:"DBInstanceId" json:"DBInstanceId"`
	TotalRecordsCount int    `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceOnMissIndex []struct {
			DBName      string `xml:"DBName" json:"DBName"`
			TableName   string `xml:"TableName" json:"TableName"`
			QueryColumn string `xml:"QueryColumn" json:"QueryColumn"`
			SQLText     string `xml:"SQLText" json:"SQLText"`
		} `xml:"AdviceOnMissIndex" json:"AdviceOnMissIndex"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceOnMissIndex(req *DescribeOptimizeAdviceOnMissIndexRequest, accessId, accessSecret string) (*DescribeOptimizeAdviceOnMissIndexResponse, error) {
	var pResponse DescribeOptimizeAdviceOnMissIndexResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
