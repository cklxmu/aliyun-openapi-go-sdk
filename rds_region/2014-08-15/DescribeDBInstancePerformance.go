package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstancePerformanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	Key                  string
	StartTime            string
	EndTime              string
	OwnerAccount         string
}

func (r *DescribeDBInstancePerformanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancePerformanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstancePerformanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstancePerformanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancePerformanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstancePerformanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstancePerformanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstancePerformanceRequest) SetKey(value string) {
	r.Key = value
	r.QueryParams.Set("Key", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetKey() string {
	return r.Key
}
func (r *DescribeDBInstancePerformanceRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeDBInstancePerformanceRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeDBInstancePerformanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstancePerformanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstancePerformanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstancePerformance")
	r.SetProduct(Product)
}

type DescribeDBInstancePerformanceResponse struct {
	DBInstanceId    string `xml:"DBInstanceId" json:"DBInstanceId"`
	Engine          string `xml:"Engine" json:"Engine"`
	StartTime       string `xml:"StartTime" json:"StartTime"`
	EndTime         string `xml:"EndTime" json:"EndTime"`
	PerformanceKeys struct {
		PerformanceKey []struct {
			Key         string `xml:"Key" json:"Key"`
			Unit        string `xml:"Unit" json:"Unit"`
			ValueFormat string `xml:"ValueFormat" json:"ValueFormat"`
			Values      struct {
				PerformanceValue []struct {
					Value string `xml:"Value" json:"Value"`
					Date  string `xml:"Date" json:"Date"`
				} `xml:"PerformanceValue" json:"PerformanceValue"`
			} `xml:"Values" json:"Values"`
		} `xml:"PerformanceKey" json:"PerformanceKey"`
	} `xml:"PerformanceKeys" json:"PerformanceKeys"`
}

func DescribeDBInstancePerformance(req *DescribeDBInstancePerformanceRequest, accessId, accessSecret string) (*DescribeDBInstancePerformanceResponse, error) {
	var pResponse DescribeDBInstancePerformanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
