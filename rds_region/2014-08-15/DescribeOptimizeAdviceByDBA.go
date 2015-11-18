package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceByDBARequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceByDBARequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceByDBARequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceByDBARequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceByDBARequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceByDBARequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceByDBARequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceByDBARequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceByDBARequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceByDBARequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceByDBARequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceByDBARequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceByDBARequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceByDBARequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceByDBARequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceByDBARequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceByDBA")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceByDBAResponse struct {
	TotalRecordsCount int `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceByDBA []struct {
			OptimizationText string `xml:"OptimizationText" json:"OptimizationText"`
		} `xml:"AdviceByDBA" json:"AdviceByDBA"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceByDBA(req *DescribeOptimizeAdviceByDBARequest, accessId, accessSecret string) (*DescribeOptimizeAdviceByDBAResponse, error) {
	var pResponse DescribeOptimizeAdviceByDBAResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
