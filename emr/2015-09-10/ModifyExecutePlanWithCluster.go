package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyExecutePlanWithClusterRequest struct {
	RpcRequest
	ClusterName       string
	ZoneId            string
	LogEnable         bool
	LogPath           string
	SecurityGroup     string
	IsOpenPublicIp    bool
	SecurityGroupName string
	EmrVer            string
	ClusterType       string
	Install           string
	MasterIndex       int
	EcsOrder          string
	EmrRole4ECS       string
	EmrRole4Oss       string
	PayType           int
	Period            int
	ExecutePlanId     int
	Name              string
	Strategy          int
	TimeInterval      int
	StartTime         string
	TimeUnit          string
	JobId             string
}

func (r *ModifyExecutePlanWithClusterRequest) SetClusterName(value string) {
	r.ClusterName = value
	r.QueryParams.Set("ClusterName", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetClusterName() string {
	return r.ClusterName
}
func (r *ModifyExecutePlanWithClusterRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *ModifyExecutePlanWithClusterRequest) SetLogEnable(value bool) {
	r.LogEnable = value
	r.QueryParams.Set("LogEnable", strconv.FormatBool(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetLogEnable() bool {
	return r.LogEnable
}
func (r *ModifyExecutePlanWithClusterRequest) SetLogPath(value string) {
	r.LogPath = value
	r.QueryParams.Set("LogPath", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetLogPath() string {
	return r.LogPath
}
func (r *ModifyExecutePlanWithClusterRequest) SetSecurityGroup(value string) {
	r.SecurityGroup = value
	r.QueryParams.Set("SecurityGroup", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetSecurityGroup() string {
	return r.SecurityGroup
}
func (r *ModifyExecutePlanWithClusterRequest) SetIsOpenPublicIp(value bool) {
	r.IsOpenPublicIp = value
	r.QueryParams.Set("IsOpenPublicIp", strconv.FormatBool(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetIsOpenPublicIp() bool {
	return r.IsOpenPublicIp
}
func (r *ModifyExecutePlanWithClusterRequest) SetSecurityGroupName(value string) {
	r.SecurityGroupName = value
	r.QueryParams.Set("SecurityGroupName", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetSecurityGroupName() string {
	return r.SecurityGroupName
}
func (r *ModifyExecutePlanWithClusterRequest) SetEmrVer(value string) {
	r.EmrVer = value
	r.QueryParams.Set("EmrVer", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetEmrVer() string {
	return r.EmrVer
}
func (r *ModifyExecutePlanWithClusterRequest) SetClusterType(value string) {
	r.ClusterType = value
	r.QueryParams.Set("ClusterType", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetClusterType() string {
	return r.ClusterType
}
func (r *ModifyExecutePlanWithClusterRequest) SetInstall(value string) {
	r.Install = value
	r.QueryParams.Set("Install", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetInstall() string {
	return r.Install
}
func (r *ModifyExecutePlanWithClusterRequest) SetMasterIndex(value int) {
	r.MasterIndex = value
	r.QueryParams.Set("MasterIndex", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetMasterIndex() int {
	return r.MasterIndex
}
func (r *ModifyExecutePlanWithClusterRequest) SetEcsOrder(value string) {
	r.EcsOrder = value
	r.QueryParams.Set("EcsOrder", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetEcsOrder() string {
	return r.EcsOrder
}
func (r *ModifyExecutePlanWithClusterRequest) SetEmrRole4ECS(value string) {
	r.EmrRole4ECS = value
	r.QueryParams.Set("EmrRole4ECS", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetEmrRole4ECS() string {
	return r.EmrRole4ECS
}
func (r *ModifyExecutePlanWithClusterRequest) SetEmrRole4Oss(value string) {
	r.EmrRole4Oss = value
	r.QueryParams.Set("EmrRole4Oss", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetEmrRole4Oss() string {
	return r.EmrRole4Oss
}
func (r *ModifyExecutePlanWithClusterRequest) SetPayType(value int) {
	r.PayType = value
	r.QueryParams.Set("PayType", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetPayType() int {
	return r.PayType
}
func (r *ModifyExecutePlanWithClusterRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetPeriod() int {
	return r.Period
}
func (r *ModifyExecutePlanWithClusterRequest) SetExecutePlanId(value int) {
	r.ExecutePlanId = value
	r.QueryParams.Set("ExecutePlanId", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetExecutePlanId() int {
	return r.ExecutePlanId
}
func (r *ModifyExecutePlanWithClusterRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetName() string {
	return r.Name
}
func (r *ModifyExecutePlanWithClusterRequest) SetStrategy(value int) {
	r.Strategy = value
	r.QueryParams.Set("Strategy", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetStrategy() int {
	return r.Strategy
}
func (r *ModifyExecutePlanWithClusterRequest) SetTimeInterval(value int) {
	r.TimeInterval = value
	r.QueryParams.Set("TimeInterval", strconv.Itoa(value))
}
func (r *ModifyExecutePlanWithClusterRequest) GetTimeInterval() int {
	return r.TimeInterval
}
func (r *ModifyExecutePlanWithClusterRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetStartTime() string {
	return r.StartTime
}
func (r *ModifyExecutePlanWithClusterRequest) SetTimeUnit(value string) {
	r.TimeUnit = value
	r.QueryParams.Set("TimeUnit", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetTimeUnit() string {
	return r.TimeUnit
}
func (r *ModifyExecutePlanWithClusterRequest) SetJobId(value string) {
	r.JobId = value
	r.QueryParams.Set("JobId", value)
}
func (r *ModifyExecutePlanWithClusterRequest) GetJobId() string {
	return r.JobId
}

func (r *ModifyExecutePlanWithClusterRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyExecutePlanWithCluster")
	r.SetProduct(Product)
}

type ModifyExecutePlanWithClusterResponse struct {
}

func ModifyExecutePlanWithCluster(req *ModifyExecutePlanWithClusterRequest, accessId, accessSecret string) (*ModifyExecutePlanWithClusterResponse, error) {
	var pResponse ModifyExecutePlanWithClusterResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
