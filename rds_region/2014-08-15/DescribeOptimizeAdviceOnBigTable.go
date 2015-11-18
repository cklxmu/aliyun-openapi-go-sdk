package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceOnBigTableRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceOnBigTableRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnBigTableRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceOnBigTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceOnBigTable")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceOnBigTableResponse struct {
	TotalRecordsCount int `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceOnBigTable []struct {
			DBName    string `xml:"DBName" json:"DBName"`
			TableName string `xml:"TableName" json:"TableName"`
			TableSize int    `xml:"TableSize" json:"TableSize"`
			DataSize  int    `xml:"DataSize" json:"DataSize"`
			IndexSize int    `xml:"IndexSize" json:"IndexSize"`
		} `xml:"AdviceOnBigTable" json:"AdviceOnBigTable"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceOnBigTable(req *DescribeOptimizeAdviceOnBigTableRequest, accessId, accessSecret string) (*DescribeOptimizeAdviceOnBigTableResponse, error) {
	var pResponse DescribeOptimizeAdviceOnBigTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
