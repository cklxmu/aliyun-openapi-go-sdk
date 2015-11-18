package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeAccountsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	OwnerAccount         string
}

func (r *DescribeAccountsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeAccountsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeAccountsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeAccountsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeAccountsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeAccountsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeAccountsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeAccountsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeAccountsRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *DescribeAccountsRequest) GetAccountName() string {
	return r.AccountName
}
func (r *DescribeAccountsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeAccountsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeAccountsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeAccounts")
	r.SetProduct(Product)
}

type DescribeAccountsResponse struct {
	Accounts struct {
		DBInstanceAccount []struct {
			DBInstanceId       string `xml:"DBInstanceId" json:"DBInstanceId"`
			AccountName        string `xml:"AccountName" json:"AccountName"`
			AccountStatus      string `xml:"AccountStatus" json:"AccountStatus"`
			AccountDescription string `xml:"AccountDescription" json:"AccountDescription"`
			DatabasePrivileges struct {
				DatabasePrivilege []struct {
					DBName           string `xml:"DBName" json:"DBName"`
					AccountPrivilege string `xml:"AccountPrivilege" json:"AccountPrivilege"`
				} `xml:"DatabasePrivilege" json:"DatabasePrivilege"`
			} `xml:"DatabasePrivileges" json:"DatabasePrivileges"`
		} `xml:"DBInstanceAccount" json:"DBInstanceAccount"`
	} `xml:"Accounts" json:"Accounts"`
}

func DescribeAccounts(req *DescribeAccountsRequest, accessId, accessSecret string) (*DescribeAccountsResponse, error) {
	var pResponse DescribeAccountsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
