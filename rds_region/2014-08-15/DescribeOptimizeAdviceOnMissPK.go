package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOptimizeAdviceOnMissPKRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeOptimizeAdviceOnMissPKRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeOptimizeAdviceOnMissPKRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeOptimizeAdviceOnMissPKRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOptimizeAdviceOnMissPK")
	r.SetProduct(Product)
}

type DescribeOptimizeAdviceOnMissPKResponse struct {
	TotalRecordsCount int `xml:"TotalRecordsCount" json:"TotalRecordsCount"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount   int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items             struct {
		AdviceOnMissPK []struct {
			DBName    string `xml:"DBName" json:"DBName"`
			TableName string `xml:"TableName" json:"TableName"`
		} `xml:"AdviceOnMissPK" json:"AdviceOnMissPK"`
	} `xml:"Items" json:"Items"`
}

func DescribeOptimizeAdviceOnMissPK(req *DescribeOptimizeAdviceOnMissPKRequest, accessId, accessSecret string) (*DescribeOptimizeAdviceOnMissPKResponse, error) {
	var pResponse DescribeOptimizeAdviceOnMissPKResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
