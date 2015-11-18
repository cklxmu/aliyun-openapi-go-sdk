package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceOnExcessIndexRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnExcessIndexRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceOnExcessIndexRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceOnExcessIndex")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceOnExcessIndexResponse struct {
	TotalRecordsCount int `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceOnExcessIndex []struct {
			DBName     string `xml:"DBName" json:"DBName"`
			TableName  string `xml:"TableName" json:"TableName"`
			IndexCount int    `xml:"IndexCount" json:"IndexCount"`
		} `xml:"AdviceOnExcessIndex" json:"AdviceOnExcessIndex"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceOnExcessIndex(req *DescribeOptimizeAdviceOnExcessIndexRequest, accessId, accessSecret string) (*DescribeOptimizeAdviceOnExcessIndexResponse, error) {
	var pResponse DescribeOptimizeAdviceOnExcessIndexResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
