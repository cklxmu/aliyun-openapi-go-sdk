package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	Engine               string
	DBInstanceStatus     string
	SearchKey            string
	DBInstanceId         string
	DBInstanceType       string
	PageSize             int
	PageNumber           int
	InstanceNetworkType  string
	ConnectionMode       string
	OwnerAccount         string
}

func (r *DescribeDBInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstancesRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstancesRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstancesRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *DescribeDBInstancesRequest) GetproxyId() string {
	return r.proxyId
}
func (r *DescribeDBInstancesRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *DescribeDBInstancesRequest) GetEngine() string {
	return r.Engine
}
func (r *DescribeDBInstancesRequest) SetDBInstanceStatus(value string) {
	r.DBInstanceStatus = value
	r.QueryParams.Set("DBInstanceStatus", value)
}
func (r *DescribeDBInstancesRequest) GetDBInstanceStatus() string {
	return r.DBInstanceStatus
}
func (r *DescribeDBInstancesRequest) SetSearchKey(value string) {
	r.SearchKey = value
	r.QueryParams.Set("SearchKey", value)
}
func (r *DescribeDBInstancesRequest) GetSearchKey() string {
	return r.SearchKey
}
func (r *DescribeDBInstancesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstancesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstancesRequest) SetDBInstanceType(value string) {
	r.DBInstanceType = value
	r.QueryParams.Set("DBInstanceType", value)
}
func (r *DescribeDBInstancesRequest) GetDBInstanceType() string {
	return r.DBInstanceType
}
func (r *DescribeDBInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeDBInstancesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeDBInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeDBInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeDBInstancesRequest) SetInstanceNetworkType(value string) {
	r.InstanceNetworkType = value
	r.QueryParams.Set("InstanceNetworkType", value)
}
func (r *DescribeDBInstancesRequest) GetInstanceNetworkType() string {
	return r.InstanceNetworkType
}
func (r *DescribeDBInstancesRequest) SetConnectionMode(value string) {
	r.ConnectionMode = value
	r.QueryParams.Set("ConnectionMode", value)
}
func (r *DescribeDBInstancesRequest) GetConnectionMode() string {
	return r.ConnectionMode
}
func (r *DescribeDBInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstances")
	r.SetProduct(Product)
}

type DescribeDBInstancesResponse struct {
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		DBInstance []struct {
			InsId                 int    `xml:"InsId" json:"InsId"`
			DBInstanceId          string `xml:"DBInstanceId" json:"DBInstanceId"`
			DBInstanceDescription string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
			PayType               string `xml:"PayType" json:"PayType"`
			DBInstanceType        string `xml:"DBInstanceType" json:"DBInstanceType"`
			RegionId              string `xml:"RegionId" json:"RegionId"`
			ExpireTime            string `xml:"ExpireTime" json:"ExpireTime"`
			DBInstanceStatus      string `xml:"DBInstanceStatus" json:"DBInstanceStatus"`
			Engine                string `xml:"Engine" json:"Engine"`
			DBInstanceNetType     string `xml:"DBInstanceNetType" json:"DBInstanceNetType"`
			ConnectionMode        string `xml:"ConnectionMode" json:"ConnectionMode"`
			LockMode              string `xml:"LockMode" json:"LockMode"`
			InstanceNetworkType   string `xml:"InstanceNetworkType" json:"InstanceNetworkType"`
			LockReason            string `xml:"LockReason" json:"LockReason"`
			ZoneId                string `xml:"ZoneId" json:"ZoneId"`
			MutriORsignle         bool   `xml:"MutriORsignle" json:"MutriORsignle"`
			CreateTime            string `xml:"CreateTime" json:"CreateTime"`
			EngineVersion         string `xml:"EngineVersion" json:"EngineVersion"`
			GuardDBInstanceId     string `xml:"GuardDBInstanceId" json:"GuardDBInstanceId"`
			TempDBInstanceId      string `xml:"TempDBInstanceId" json:"TempDBInstanceId"`
			MasterInstanceId      string `xml:"MasterInstanceId" json:"MasterInstanceId"`
			VpcId                 string `xml:"VpcId" json:"VpcId"`
			ReadOnlyDBInstanceIds struct {
				ReadOnlyDBInstanceId []struct {
					DBInstanceId string `xml:"DBInstanceId" json:"DBInstanceId"`
				} `xml:"ReadOnlyDBInstanceId" json:"ReadOnlyDBInstanceId"`
			} `xml:"ReadOnlyDBInstanceIds" json:"ReadOnlyDBInstanceIds"`
		} `xml:"DBInstance" json:"DBInstance"`
	} `xml:"Items" json:"Items"`
}

func DescribeDBInstances(req *DescribeDBInstancesRequest, accessId, accessSecret string) (*DescribeDBInstancesResponse, error) {
	var pResponse DescribeDBInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
