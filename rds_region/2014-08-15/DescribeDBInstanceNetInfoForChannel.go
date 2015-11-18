package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstanceNetInfoForChannelRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	Flag                 string
	OwnerAccount         string
}

func (r *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetFlag(value string) {
	r.Flag = value
	r.QueryParams.Set("Flag", value)
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetFlag() string {
	return r.Flag
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstanceNetInfoForChannelRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstanceNetInfoForChannelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstanceNetInfoForChannel")
	r.SetProduct(Product)
}

type DescribeDBInstanceNetInfoForChannelResponse struct {
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

func DescribeDBInstanceNetInfoForChannel(req *DescribeDBInstanceNetInfoForChannelRequest, accessId, accessSecret string) (*DescribeDBInstanceNetInfoForChannelResponse, error) {
	var pResponse DescribeDBInstanceNetInfoForChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
