package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeParametersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DescribeParametersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeParametersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeParametersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeParametersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeParametersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeParametersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeParametersRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeParametersRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeParametersRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeParametersRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeParametersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeParametersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeParametersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeParameters")
	r.SetProduct(Product)
}

type DescribeParametersResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	EngineVersion    string `xml:"EngineVersion" json:"EngineVersion"`
	ConfigParameters struct {
		DBInstanceParameter []struct {
			ParameterName        string `xml:"ParameterName" json:"ParameterName"`
			ParameterValue       string `xml:"ParameterValue" json:"ParameterValue"`
			ParameterDescription string `xml:"ParameterDescription" json:"ParameterDescription"`
		} `xml:"DBInstanceParameter" json:"DBInstanceParameter"`
	} `xml:"ConfigParameters" json:"ConfigParameters"`
	RunningParameters struct {
		DBInstanceParameter []struct {
			ParameterName        string `xml:"ParameterName" json:"ParameterName"`
			ParameterValue       string `xml:"ParameterValue" json:"ParameterValue"`
			ParameterDescription string `xml:"ParameterDescription" json:"ParameterDescription"`
		} `xml:"DBInstanceParameter" json:"DBInstanceParameter"`
	} `xml:"RunningParameters" json:"RunningParameters"`
}

func DescribeParameters(req *DescribeParametersRequest, accessId, accessSecret string) (*DescribeParametersResponse, error) {
	var pResponse DescribeParametersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
