package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDatabasesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	DBStatus             string
	OwnerAccount         string
}

func (r *DescribeDatabasesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDatabasesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDatabasesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDatabasesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDatabasesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDatabasesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDatabasesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDatabasesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDatabasesRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *DescribeDatabasesRequest) GetDBName() string {
	return r.DBName
}
func (r *DescribeDatabasesRequest) SetDBStatus(value string) {
	r.DBStatus = value
	r.QueryParams.Set("DBStatus", value)
}
func (r *DescribeDatabasesRequest) GetDBStatus() string {
	return r.DBStatus
}
func (r *DescribeDatabasesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDatabasesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDatabasesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDatabases")
	r.SetProduct(Product)
}

type DescribeDatabasesResponse struct {
	Databases struct {
		Database []struct {
			DBName           string `xml:"DBName" json:"DBName"`
			DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
			Engine           string `xml:"Engine" json:"Engine"`
			DBStatus         string `xml:"DBStatus" json:"DBStatus"`
			CharacterSetName string `xml:"CharacterSetName" json:"CharacterSetName"`
			DBDescription    string `xml:"DBDescription" json:"DBDescription"`
			Accounts         struct {
				AccountPrivilegeInfo []struct {
					Account          string `xml:"Account" json:"Account"`
					AccountPrivilege string `xml:"AccountPrivilege" json:"AccountPrivilege"`
				} `xml:"AccountPrivilegeInfo" json:"AccountPrivilegeInfo"`
			} `xml:"Accounts" json:"Accounts"`
		} `xml:"Database" json:"Database"`
	} `xml:"Databases" json:"Databases"`
}

func DescribeDatabases(req *DescribeDatabasesRequest, accessId, accessSecret string) (*DescribeDatabasesResponse, error) {
	var pResponse DescribeDatabasesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
