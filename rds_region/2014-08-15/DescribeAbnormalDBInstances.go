package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeAbnormalDBInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	DBInstanceId         string
	PageSize             int
	PageNumber           int
	Tag_1_key            string
	Tag_2_key            string
	Tag_3_key            string
	Tag_4_key            string
	Tag_5_key            string
	Tag_1_value          string
	Tag_2_value          string
	Tag_3_value          string
	Tag_4_value          string
	Tag_5_value          string
	OwnerAccount         string
}

func (r *DescribeAbnormalDBInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeAbnormalDBInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeAbnormalDBInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeAbnormalDBInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeAbnormalDBInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeAbnormalDBInstancesRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeAbnormalDBInstancesRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetproxyId() string {
	return r.proxyId
}
func (r *DescribeAbnormalDBInstancesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeAbnormalDBInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeAbnormalDBInstancesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeAbnormalDBInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeAbnormalDBInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_1_key(value string) {
	r.Tag_1_key = value
	r.QueryParams.Set("Tag.1.key", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_1_key() string {
	return r.Tag_1_key
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_2_key(value string) {
	r.Tag_2_key = value
	r.QueryParams.Set("Tag.2.key", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_2_key() string {
	return r.Tag_2_key
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_3_key(value string) {
	r.Tag_3_key = value
	r.QueryParams.Set("Tag.3.key", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_3_key() string {
	return r.Tag_3_key
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_4_key(value string) {
	r.Tag_4_key = value
	r.QueryParams.Set("Tag.4.key", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_4_key() string {
	return r.Tag_4_key
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_5_key(value string) {
	r.Tag_5_key = value
	r.QueryParams.Set("Tag.5.key", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_5_key() string {
	return r.Tag_5_key
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_1_value(value string) {
	r.Tag_1_value = value
	r.QueryParams.Set("Tag.1.value", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_1_value() string {
	return r.Tag_1_value
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_2_value(value string) {
	r.Tag_2_value = value
	r.QueryParams.Set("Tag.2.value", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_2_value() string {
	return r.Tag_2_value
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_3_value(value string) {
	r.Tag_3_value = value
	r.QueryParams.Set("Tag.3.value", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_3_value() string {
	return r.Tag_3_value
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_4_value(value string) {
	r.Tag_4_value = value
	r.QueryParams.Set("Tag.4.value", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_4_value() string {
	return r.Tag_4_value
}
func (r *DescribeAbnormalDBInstancesRequest) SetTag_5_value(value string) {
	r.Tag_5_value = value
	r.QueryParams.Set("Tag.5.value", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetTag_5_value() string {
	return r.Tag_5_value
}
func (r *DescribeAbnormalDBInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeAbnormalDBInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeAbnormalDBInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeAbnormalDBInstances")
	r.SetProduct(Product)
}

type DescribeAbnormalDBInstancesResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		InstanceResult []struct {
			DBInstanceDescription string `xml:"DBInstanceDescription" json:"DBInstanceDescription"`
			DBInstanceId          string `xml:"DBInstanceId" json:"DBInstanceId"`
			AbnormalItems         struct {
				AbnormalItem []struct {
					CheckTime      string `xml:"CheckTime" json:"CheckTime"`
					CheckItem      string `xml:"CheckItem" json:"CheckItem"`
					AbnormalReason string `xml:"AbnormalReason" json:"AbnormalReason"`
					AbnormalValue  string `xml:"AbnormalValue" json:"AbnormalValue"`
					AbnormalDetail string `xml:"AbnormalDetail" json:"AbnormalDetail"`
					AdviceKey      string `xml:"AdviceKey" json:"AdviceKey"`
					AdviseValue    struct {
						String []string `xml:"String" json:"String"`
					} `xml:"AdviseValue" json:"AdviseValue"`
				} `xml:"AbnormalItem" json:"AbnormalItem"`
			} `xml:"AbnormalItems" json:"AbnormalItems"`
		} `xml:"InstanceResult" json:"InstanceResult"`
	} `xml:"Items" json:"Items"`
}

func DescribeAbnormalDBInstances(req *DescribeAbnormalDBInstancesRequest, accessId, accessSecret string) (*DescribeAbnormalDBInstancesResponse, error) {
	var pResponse DescribeAbnormalDBInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
