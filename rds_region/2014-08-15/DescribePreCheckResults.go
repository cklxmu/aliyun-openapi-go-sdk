package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribePreCheckResultsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	PreCheckId           string
	OwnerAccount         string
}

func (r *DescribePreCheckResultsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribePreCheckResultsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribePreCheckResultsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribePreCheckResultsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribePreCheckResultsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribePreCheckResultsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribePreCheckResultsRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribePreCheckResultsRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribePreCheckResultsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribePreCheckResultsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribePreCheckResultsRequest) SetPreCheckId(value string) {
	r.PreCheckId = value
	r.QueryParams.Set("PreCheckId", value)
}
func (r *DescribePreCheckResultsRequest) GetPreCheckId() string {
	return r.PreCheckId
}
func (r *DescribePreCheckResultsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribePreCheckResultsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribePreCheckResultsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribePreCheckResults")
	r.SetProduct(Product)
}

type DescribePreCheckResultsResponse struct {
	DBInstanceId string `xml:"DBInstanceId" json:"DBInstanceId"`
	Items        struct {
		PreCheckResult []struct {
			PreCheckName   string `xml:"PreCheckName" json:"PreCheckName"`
			PreCheckResult string `xml:"PreCheckResult" json:"PreCheckResult"`
			FailReasion    string `xml:"FailReasion" json:"FailReasion"`
			RepairMethod   string `xml:"RepairMethod" json:"RepairMethod"`
		} `xml:"PreCheckResult" json:"PreCheckResult"`
	} `xml:"Items" json:"Items"`
}

func DescribePreCheckResults(req *DescribePreCheckResultsRequest, accessId, accessSecret string) (*DescribePreCheckResultsResponse, error) {
	var pResponse DescribePreCheckResultsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
