package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RemoveInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingGroupId       string
	InstanceId_1         string
	InstanceId_2         string
	InstanceId_3         string
	InstanceId_4         string
	InstanceId_5         string
	InstanceId_6         string
	InstanceId_7         string
	InstanceId_8         string
	InstanceId_9         string
	InstanceId_10        string
	InstanceId_11        string
	InstanceId_12        string
	InstanceId_13        string
	InstanceId_14        string
	InstanceId_15        string
	InstanceId_16        string
	InstanceId_17        string
	InstanceId_18        string
	InstanceId_19        string
	InstanceId_20        string
	OwnerAccount         string
}

func (r *RemoveInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RemoveInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RemoveInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RemoveInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RemoveInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RemoveInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RemoveInstancesRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *RemoveInstancesRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *RemoveInstancesRequest) SetInstanceId_1(value string) {
	r.InstanceId_1 = value
	r.QueryParams.Set("InstanceId.1", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_1() string {
	return r.InstanceId_1
}
func (r *RemoveInstancesRequest) SetInstanceId_2(value string) {
	r.InstanceId_2 = value
	r.QueryParams.Set("InstanceId.2", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_2() string {
	return r.InstanceId_2
}
func (r *RemoveInstancesRequest) SetInstanceId_3(value string) {
	r.InstanceId_3 = value
	r.QueryParams.Set("InstanceId.3", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_3() string {
	return r.InstanceId_3
}
func (r *RemoveInstancesRequest) SetInstanceId_4(value string) {
	r.InstanceId_4 = value
	r.QueryParams.Set("InstanceId.4", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_4() string {
	return r.InstanceId_4
}
func (r *RemoveInstancesRequest) SetInstanceId_5(value string) {
	r.InstanceId_5 = value
	r.QueryParams.Set("InstanceId.5", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_5() string {
	return r.InstanceId_5
}
func (r *RemoveInstancesRequest) SetInstanceId_6(value string) {
	r.InstanceId_6 = value
	r.QueryParams.Set("InstanceId.6", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_6() string {
	return r.InstanceId_6
}
func (r *RemoveInstancesRequest) SetInstanceId_7(value string) {
	r.InstanceId_7 = value
	r.QueryParams.Set("InstanceId.7", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_7() string {
	return r.InstanceId_7
}
func (r *RemoveInstancesRequest) SetInstanceId_8(value string) {
	r.InstanceId_8 = value
	r.QueryParams.Set("InstanceId.8", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_8() string {
	return r.InstanceId_8
}
func (r *RemoveInstancesRequest) SetInstanceId_9(value string) {
	r.InstanceId_9 = value
	r.QueryParams.Set("InstanceId.9", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_9() string {
	return r.InstanceId_9
}
func (r *RemoveInstancesRequest) SetInstanceId_10(value string) {
	r.InstanceId_10 = value
	r.QueryParams.Set("InstanceId.10", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_10() string {
	return r.InstanceId_10
}
func (r *RemoveInstancesRequest) SetInstanceId_11(value string) {
	r.InstanceId_11 = value
	r.QueryParams.Set("InstanceId.11", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_11() string {
	return r.InstanceId_11
}
func (r *RemoveInstancesRequest) SetInstanceId_12(value string) {
	r.InstanceId_12 = value
	r.QueryParams.Set("InstanceId.12", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_12() string {
	return r.InstanceId_12
}
func (r *RemoveInstancesRequest) SetInstanceId_13(value string) {
	r.InstanceId_13 = value
	r.QueryParams.Set("InstanceId.13", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_13() string {
	return r.InstanceId_13
}
func (r *RemoveInstancesRequest) SetInstanceId_14(value string) {
	r.InstanceId_14 = value
	r.QueryParams.Set("InstanceId.14", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_14() string {
	return r.InstanceId_14
}
func (r *RemoveInstancesRequest) SetInstanceId_15(value string) {
	r.InstanceId_15 = value
	r.QueryParams.Set("InstanceId.15", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_15() string {
	return r.InstanceId_15
}
func (r *RemoveInstancesRequest) SetInstanceId_16(value string) {
	r.InstanceId_16 = value
	r.QueryParams.Set("InstanceId.16", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_16() string {
	return r.InstanceId_16
}
func (r *RemoveInstancesRequest) SetInstanceId_17(value string) {
	r.InstanceId_17 = value
	r.QueryParams.Set("InstanceId.17", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_17() string {
	return r.InstanceId_17
}
func (r *RemoveInstancesRequest) SetInstanceId_18(value string) {
	r.InstanceId_18 = value
	r.QueryParams.Set("InstanceId.18", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_18() string {
	return r.InstanceId_18
}
func (r *RemoveInstancesRequest) SetInstanceId_19(value string) {
	r.InstanceId_19 = value
	r.QueryParams.Set("InstanceId.19", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_19() string {
	return r.InstanceId_19
}
func (r *RemoveInstancesRequest) SetInstanceId_20(value string) {
	r.InstanceId_20 = value
	r.QueryParams.Set("InstanceId.20", value)
}
func (r *RemoveInstancesRequest) GetInstanceId_20() string {
	return r.InstanceId_20
}
func (r *RemoveInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RemoveInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RemoveInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveInstances")
	r.SetProduct(Product)
}

type RemoveInstancesResponse struct {
	ScalingActivityId string `xml:"ScalingActivityId" json:"ScalingActivityId"`
}

func RemoveInstances(req *RemoveInstancesRequest, accessId, accessSecret string) (*RemoveInstancesResponse, error) {
	var pResponse RemoveInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
