package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScalingRulesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageNumber           int
	PageSize             int
	ScalingGroupId       string
	ScalingRuleId_1      string
	ScalingRuleId_2      string
	ScalingRuleId_3      string
	ScalingRuleId_4      string
	ScalingRuleId_5      string
	ScalingRuleId_6      string
	ScalingRuleId_7      string
	ScalingRuleId_8      string
	ScalingRuleId_9      string
	ScalingRuleId_10     string
	ScalingRuleName_1    string
	ScalingRuleName_2    string
	ScalingRuleName_3    string
	ScalingRuleName_4    string
	ScalingRuleName_5    string
	ScalingRuleName_6    string
	ScalingRuleName_7    string
	ScalingRuleName_8    string
	ScalingRuleName_9    string
	ScalingRuleName_10   string
	ScalingRuleAri_1     string
	ScalingRuleAri_2     string
	ScalingRuleAri_3     string
	ScalingRuleAri_4     string
	ScalingRuleAri_5     string
	ScalingRuleAri_6     string
	ScalingRuleAri_7     string
	ScalingRuleAri_8     string
	ScalingRuleAri_9     string
	ScalingRuleAri_10    string
	OwnerAccount         string
}

func (r *DescribeScalingRulesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingRulesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScalingRulesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScalingRulesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScalingRulesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingRulesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScalingRulesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScalingRulesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScalingRulesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScalingRulesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScalingRulesRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DescribeScalingRulesRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_1(value string) {
	r.ScalingRuleId_1 = value
	r.QueryParams.Set("ScalingRuleId.1", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_1() string {
	return r.ScalingRuleId_1
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_2(value string) {
	r.ScalingRuleId_2 = value
	r.QueryParams.Set("ScalingRuleId.2", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_2() string {
	return r.ScalingRuleId_2
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_3(value string) {
	r.ScalingRuleId_3 = value
	r.QueryParams.Set("ScalingRuleId.3", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_3() string {
	return r.ScalingRuleId_3
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_4(value string) {
	r.ScalingRuleId_4 = value
	r.QueryParams.Set("ScalingRuleId.4", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_4() string {
	return r.ScalingRuleId_4
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_5(value string) {
	r.ScalingRuleId_5 = value
	r.QueryParams.Set("ScalingRuleId.5", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_5() string {
	return r.ScalingRuleId_5
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_6(value string) {
	r.ScalingRuleId_6 = value
	r.QueryParams.Set("ScalingRuleId.6", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_6() string {
	return r.ScalingRuleId_6
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_7(value string) {
	r.ScalingRuleId_7 = value
	r.QueryParams.Set("ScalingRuleId.7", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_7() string {
	return r.ScalingRuleId_7
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_8(value string) {
	r.ScalingRuleId_8 = value
	r.QueryParams.Set("ScalingRuleId.8", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_8() string {
	return r.ScalingRuleId_8
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_9(value string) {
	r.ScalingRuleId_9 = value
	r.QueryParams.Set("ScalingRuleId.9", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_9() string {
	return r.ScalingRuleId_9
}
func (r *DescribeScalingRulesRequest) SetScalingRuleId_10(value string) {
	r.ScalingRuleId_10 = value
	r.QueryParams.Set("ScalingRuleId.10", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleId_10() string {
	return r.ScalingRuleId_10
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_1(value string) {
	r.ScalingRuleName_1 = value
	r.QueryParams.Set("ScalingRuleName.1", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_1() string {
	return r.ScalingRuleName_1
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_2(value string) {
	r.ScalingRuleName_2 = value
	r.QueryParams.Set("ScalingRuleName.2", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_2() string {
	return r.ScalingRuleName_2
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_3(value string) {
	r.ScalingRuleName_3 = value
	r.QueryParams.Set("ScalingRuleName.3", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_3() string {
	return r.ScalingRuleName_3
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_4(value string) {
	r.ScalingRuleName_4 = value
	r.QueryParams.Set("ScalingRuleName.4", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_4() string {
	return r.ScalingRuleName_4
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_5(value string) {
	r.ScalingRuleName_5 = value
	r.QueryParams.Set("ScalingRuleName.5", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_5() string {
	return r.ScalingRuleName_5
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_6(value string) {
	r.ScalingRuleName_6 = value
	r.QueryParams.Set("ScalingRuleName.6", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_6() string {
	return r.ScalingRuleName_6
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_7(value string) {
	r.ScalingRuleName_7 = value
	r.QueryParams.Set("ScalingRuleName.7", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_7() string {
	return r.ScalingRuleName_7
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_8(value string) {
	r.ScalingRuleName_8 = value
	r.QueryParams.Set("ScalingRuleName.8", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_8() string {
	return r.ScalingRuleName_8
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_9(value string) {
	r.ScalingRuleName_9 = value
	r.QueryParams.Set("ScalingRuleName.9", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_9() string {
	return r.ScalingRuleName_9
}
func (r *DescribeScalingRulesRequest) SetScalingRuleName_10(value string) {
	r.ScalingRuleName_10 = value
	r.QueryParams.Set("ScalingRuleName.10", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleName_10() string {
	return r.ScalingRuleName_10
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_1(value string) {
	r.ScalingRuleAri_1 = value
	r.QueryParams.Set("ScalingRuleAri.1", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_1() string {
	return r.ScalingRuleAri_1
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_2(value string) {
	r.ScalingRuleAri_2 = value
	r.QueryParams.Set("ScalingRuleAri.2", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_2() string {
	return r.ScalingRuleAri_2
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_3(value string) {
	r.ScalingRuleAri_3 = value
	r.QueryParams.Set("ScalingRuleAri.3", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_3() string {
	return r.ScalingRuleAri_3
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_4(value string) {
	r.ScalingRuleAri_4 = value
	r.QueryParams.Set("ScalingRuleAri.4", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_4() string {
	return r.ScalingRuleAri_4
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_5(value string) {
	r.ScalingRuleAri_5 = value
	r.QueryParams.Set("ScalingRuleAri.5", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_5() string {
	return r.ScalingRuleAri_5
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_6(value string) {
	r.ScalingRuleAri_6 = value
	r.QueryParams.Set("ScalingRuleAri.6", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_6() string {
	return r.ScalingRuleAri_6
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_7(value string) {
	r.ScalingRuleAri_7 = value
	r.QueryParams.Set("ScalingRuleAri.7", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_7() string {
	return r.ScalingRuleAri_7
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_8(value string) {
	r.ScalingRuleAri_8 = value
	r.QueryParams.Set("ScalingRuleAri.8", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_8() string {
	return r.ScalingRuleAri_8
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_9(value string) {
	r.ScalingRuleAri_9 = value
	r.QueryParams.Set("ScalingRuleAri.9", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_9() string {
	return r.ScalingRuleAri_9
}
func (r *DescribeScalingRulesRequest) SetScalingRuleAri_10(value string) {
	r.ScalingRuleAri_10 = value
	r.QueryParams.Set("ScalingRuleAri.10", value)
}
func (r *DescribeScalingRulesRequest) GetScalingRuleAri_10() string {
	return r.ScalingRuleAri_10
}
func (r *DescribeScalingRulesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScalingRulesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScalingRulesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScalingRules")
	r.SetProduct(Product)
}

type DescribeScalingRulesResponse struct {
	TotalCount   int `xml:"TotalCount" json:"TotalCount"`
	PageNumber   int `xml:"PageNumber" json:"PageNumber"`
	PageSize     int `xml:"PageSize" json:"PageSize"`
	ScalingRules struct {
		ScalingRule []struct {
			ScalingRuleId   string `xml:"ScalingRuleId" json:"ScalingRuleId"`
			ScalingGroupId  string `xml:"ScalingGroupId" json:"ScalingGroupId"`
			ScalingRuleName string `xml:"ScalingRuleName" json:"ScalingRuleName"`
			Cooldown        int    `xml:"Cooldown" json:"Cooldown"`
			AdjustmentType  string `xml:"AdjustmentType" json:"AdjustmentType"`
			AdjustmentValue int    `xml:"AdjustmentValue" json:"AdjustmentValue"`
			MinSize         int    `xml:"MinSize" json:"MinSize"`
			MaxSize         int    `xml:"MaxSize" json:"MaxSize"`
			ScalingRuleAri  string `xml:"ScalingRuleAri" json:"ScalingRuleAri"`
		} `xml:"ScalingRule" json:"ScalingRule"`
	} `xml:"ScalingRules" json:"ScalingRules"`
}

func DescribeScalingRules(req *DescribeScalingRulesRequest, accessId, accessSecret string) (*DescribeScalingRulesResponse, error) {
	var pResponse DescribeScalingRulesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
