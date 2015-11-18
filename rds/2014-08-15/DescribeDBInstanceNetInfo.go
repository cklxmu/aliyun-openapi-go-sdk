package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstanceNetInfoRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	Flag                 string
	OwnerAccount         string
}

func (r *DescribeDBInstanceNetInfoRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceNetInfoRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstanceNetInfoRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstanceNetInfoRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceNetInfoRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstanceNetInfoRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstanceNetInfoRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstanceNetInfoRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstanceNetInfoRequest) SetFlag(value string) {
	r.Flag = value
	r.QueryParams.Set("Flag", value)
}
func (r *DescribeDBInstanceNetInfoRequest) GetFlag() string {
	return r.Flag
}
func (r *DescribeDBInstanceNetInfoRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstanceNetInfoRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstanceNetInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstanceNetInfo")
	r.SetProduct(Product)
}

type DescribeDBInstanceNetInfoResponse struct {
	InstanceNetworkType string `xml:"InstanceNetworkType" json:"InstanceNetworkType"`
	DBInstanceNetInfos  struct {
		DBInstanceNetInfo []struct {
			ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
			IPAddress        string `xml:"IPAddress" json:"IPAddress"`
			IPType           string `xml:"IPType" json:"IPType"`
			Port             string `xml:"Port" json:"Port"`
			VPCId            string `xml:"VPCId" json:"VPCId"`
			VSwitchId        string `xml:"VSwitchId" json:"VSwitchId"`
			SecurityIPGroups struct {
				securityIPGroups []struct {
					SecurityIPGroupName string `xml:"SecurityIPGroupName" json:"SecurityIPGroupName"`
					SecurityIPs         string `xml:"SecurityIPs" json:"SecurityIPs"`
				} `xml:"securityIPGroups" json:"securityIPGroups"`
			} `xml:"SecurityIPGroups" json:"SecurityIPGroups"`
		} `xml:"DBInstanceNetInfo" json:"DBInstanceNetInfo"`
	} `xml:"DBInstanceNetInfos" json:"DBInstanceNetInfos"`
}

func DescribeDBInstanceNetInfo(req *DescribeDBInstanceNetInfoRequest, accessId, accessSecret string) (*DescribeDBInstanceNetInfoResponse, error) {
	var pResponse DescribeDBInstanceNetInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
