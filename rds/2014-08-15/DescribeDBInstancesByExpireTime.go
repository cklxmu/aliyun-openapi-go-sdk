package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstancesByExpireTimeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	ExpirePeriod         int
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeDBInstancesByExpireTimeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetproxyId() string {
	return r.proxyId
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetExpirePeriod(value int) {
	r.ExpirePeriod = value
	r.QueryParams.Set("ExpirePeriod", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetExpirePeriod() int {
	return r.ExpirePeriod
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeDBInstancesByExpireTimeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstancesByExpireTimeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstancesByExpireTimeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstancesByExpireTime")
	r.SetProduct(Product)
}

type DescribeDBInstancesByExpireTimeResponse struct {
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		DBInstanceExpireTime []struct {
			DBInstanceId          string `xml:"DBInstanceId" json:"DBInstanceId"`
			DBInstanceDescription string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
			ExpireTime            string `xml:"ExpireTime" json:"ExpireTime"`
		} `xml:"DBInstanceExpireTime" json:"DBInstanceExpireTime"`
	} `xml:"Items" json:"Items"`
}

func DescribeDBInstancesByExpireTime(req *DescribeDBInstancesByExpireTimeRequest, accessId, accessSecret string) (*DescribeDBInstancesByExpireTimeResponse, error) {
	var pResponse DescribeDBInstancesByExpireTimeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
