package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDBInstanceIPArrayListRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DescribeDBInstanceIPArrayListRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceIPArrayListRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDBInstanceIPArrayListRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDBInstanceIPArrayListRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDBInstanceIPArrayListRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDBInstanceIPArrayListRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeDBInstanceIPArrayListRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeDBInstanceIPArrayListRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDBInstanceIPArrayListRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeDBInstanceIPArrayListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDBInstanceIPArrayList")
	r.SetProduct(Product)
}

type DescribeDBInstanceIPArrayListResponse struct {
	Items struct {
		DBInstanceIPArray []struct {
			DBInstanceIPArrayName      string `xml:"DBInstanceIPArrayName" json:"DBInstanceIPArrayName"`
			DBInstanceIPArrayAttribute string `xml:"DBInstanceIPArrayAttribute" json:"DBInstanceIPArrayAttribute"`
			SecurityIPList             string `xml:"SecurityIPList" json:"SecurityIPList"`
		} `xml:"DBInstanceIPArray" json:"DBInstanceIPArray"`
	} `xml:"Items" json:"Items"`
}

func DescribeDBInstanceIPArrayList(req *DescribeDBInstanceIPArrayListRequest, accessId, accessSecret string) (*DescribeDBInstanceIPArrayListResponse, error) {
	var pResponse DescribeDBInstanceIPArrayListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
