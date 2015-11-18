package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstanceAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DescribeDBInstanceAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstanceAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstanceAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstanceAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstanceAttributeRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstanceAttributeRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstanceAttributeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstanceAttributeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstanceAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstanceAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstanceAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstanceAttribute")
	r.SetProduct(Product)
}

type DescribeDBInstanceAttributeResponse struct {
	Items struct {
		DBInstanceAttribute []struct {
			InsId                       int    `xml:"InsId" json:"InsId"`
			DBInstanceId                string `xml:"DBInstanceId" json:"DBInstanceId"`
			PayType                     string `xml:"PayType" json:"PayType"`
			DBInstanceClassType         string `xml:"DBInstanceClassType" json:"DBInstanceClassType"`
			DBInstanceType              string `xml:"DBInstanceType" json:"DBInstanceType"`
			InsId                       int    `xml:"InsId" json:"InsId"`
			RegionId                    string `xml:"RegionId" json:"RegionId"`
			ConnectionString            string `xml:"ConnectionString" json:"ConnectionString"`
			Port                        string `xml:"Port" json:"Port"`
			Engine                      string `xml:"Engine" json:"Engine"`
			EngineVersion               string `xml:"EngineVersion" json:"EngineVersion"`
			DBInstanceClass             string `xml:"DBInstanceClass" json:"DBInstanceClass"`
			DBInstanceMemory            int    `xml:"DBInstanceMemory" json:"DBInstanceMemory"`
			DBInstanceStorage           int    `xml:"DBInstanceStorage" json:"DBInstanceStorage"`
			DBInstanceNetType           string `xml:"DBInstanceNetType" json:"DBInstanceNetType"`
			DBInstanceStatus            string `xml:"DBInstanceStatus" json:"DBInstanceStatus"`
			DBInstanceDescription       string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
			LockMode                    string `xml:"LockMode" json:"LockMode"`
			LockReason                  string `xml:"LockReason" json:"LockReason"`
			ReadDelayTime               string `xml:"ReadDelayTime" json:"ReadDelayTime"`
			DBMaxQuantity               int    `xml:"DBMaxQuantity" json:"DBMaxQuantity"`
			AccountMaxQuantity          int    `xml:"AccountMaxQuantity" json:"AccountMaxQuantity"`
			CreationTime                string `xml:"CreationTime" json:"CreationTime"`
			ExpireTime                  string `xml:"ExpireTime" json:"ExpireTime"`
			MaintainTime                string `xml:"MaintainTime" json:"MaintainTime"`
			AvailabilityValue           string `xml:"AvailabilityValue" json:"AvailabilityValue"`
			MaxIOPS                     int    `xml:"MaxIOPS" json:"MaxIOPS"`
			MaxConnections              int    `xml:"MaxConnections" json:"MaxConnections"`
			MasterInstanceId            string `xml:"MasterInstanceId" json:"MasterInstanceId"`
			IncrementSourceDBInstanceId string `xml:"IncrementSourceDBInstanceId" json:"IncrementSourceDBInstanceId"`
			GuardDBInstanceId           string `xml:"GuardDBInstanceId" json:"GuardDBInstanceId"`
			TempDBInstanceId            string `xml:"TempDBInstanceId" json:"TempDBInstanceId"`
			SecurityIPList              string `xml:"SecurityIPList" json:"SecurityIPList"`
			ZoneId                      string `xml:"ZoneId" json:"ZoneId"`
			InstanceNetworkType         string `xml:"InstanceNetworkType" json:"InstanceNetworkType"`
			VpcId                       string `xml:"VpcId" json:"VpcId"`
			ConnectionMode              string `xml:"ConnectionMode" json:"ConnectionMode"`
			ReadOnlyDBInstanceIds       struct {
				ReadOnlyDBInstanceId []struct {
					DBInstanceId string `xml:"DBInstanceId" json:"DBInstanceId"`
				} `xml:"ReadOnlyDBInstanceId" json:"ReadOnlyDBInstanceId"`
			} `xml:"ReadOnlyDBInstanceIds" json:"ReadOnlyDBInstanceIds"`
		} `xml:"DBInstanceAttribute" json:"DBInstanceAttribute"`
	} `xml:"Items" json:"Items"`
}

func DescribeDBInstanceAttribute(req *DescribeDBInstanceAttributeRequest, accessId, accessSecret string) (*DescribeDBInstanceAttributeResponse, error) {
	var pResponse DescribeDBInstanceAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
