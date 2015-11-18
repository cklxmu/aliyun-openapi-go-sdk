package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateExecutePlanWithClusterRequest struct {
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
	Name              string
	Strategy          int
	TimeInterval      int
	StartTime         string
	TimeUnit          string
	JobId             string
}

func (r *CreateExecutePlanWithClusterRequest) SetClusterName(value string) {
	r.ClusterName = value
	r.QueryParams.Set("ClusterName", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetClusterName() string {
	return r.ClusterName
}
func (r *CreateExecutePlanWithClusterRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateExecutePlanWithClusterRequest) SetLogEnable(value bool) {
	r.LogEnable = value
	r.QueryParams.Set("LogEnable", strconv.FormatBool(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetLogEnable() bool {
	return r.LogEnable
}
func (r *CreateExecutePlanWithClusterRequest) SetLogPath(value string) {
	r.LogPath = value
	r.QueryParams.Set("LogPath", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetLogPath() string {
	return r.LogPath
}
func (r *CreateExecutePlanWithClusterRequest) SetSecurityGroup(value string) {
	r.SecurityGroup = value
	r.QueryParams.Set("SecurityGroup", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetSecurityGroup() string {
	return r.SecurityGroup
}
func (r *CreateExecutePlanWithClusterRequest) SetIsOpenPublicIp(value bool) {
	r.IsOpenPublicIp = value
	r.QueryParams.Set("IsOpenPublicIp", strconv.FormatBool(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetIsOpenPublicIp() bool {
	return r.IsOpenPublicIp
}
func (r *CreateExecutePlanWithClusterRequest) SetSecurityGroupName(value string) {
	r.SecurityGroupName = value
	r.QueryParams.Set("SecurityGroupName", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetSecurityGroupName() string {
	return r.SecurityGroupName
}
func (r *CreateExecutePlanWithClusterRequest) SetEmrVer(value string) {
	r.EmrVer = value
	r.QueryParams.Set("EmrVer", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetEmrVer() string {
	return r.EmrVer
}
func (r *CreateExecutePlanWithClusterRequest) SetClusterType(value string) {
	r.ClusterType = value
	r.QueryParams.Set("ClusterType", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetClusterType() string {
	return r.ClusterType
}
func (r *CreateExecutePlanWithClusterRequest) SetInstall(value string) {
	r.Install = value
	r.QueryParams.Set("Install", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetInstall() string {
	return r.Install
}
func (r *CreateExecutePlanWithClusterRequest) SetMasterIndex(value int) {
	r.MasterIndex = value
	r.QueryParams.Set("MasterIndex", strconv.Itoa(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetMasterIndex() int {
	return r.MasterIndex
}
func (r *CreateExecutePlanWithClusterRequest) SetEcsOrder(value string) {
	r.EcsOrder = value
	r.QueryParams.Set("EcsOrder", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetEcsOrder() string {
	return r.EcsOrder
}
func (r *CreateExecutePlanWithClusterRequest) SetEmrRole4ECS(value string) {
	r.EmrRole4ECS = value
	r.QueryParams.Set("EmrRole4ECS", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetEmrRole4ECS() string {
	return r.EmrRole4ECS
}
func (r *CreateExecutePlanWithClusterRequest) SetEmrRole4Oss(value string) {
	r.EmrRole4Oss = value
	r.QueryParams.Set("EmrRole4Oss", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetEmrRole4Oss() string {
	return r.EmrRole4Oss
}
func (r *CreateExecutePlanWithClusterRequest) SetPayType(value int) {
	r.PayType = value
	r.QueryParams.Set("PayType", strconv.Itoa(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetPayType() int {
	return r.PayType
}
func (r *CreateExecutePlanWithClusterRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetPeriod() int {
	return r.Period
}
func (r *CreateExecutePlanWithClusterRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetName() string {
	return r.Name
}
func (r *CreateExecutePlanWithClusterRequest) SetStrategy(value int) {
	r.Strategy = value
	r.QueryParams.Set("Strategy", strconv.Itoa(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetStrategy() int {
	return r.Strategy
}
func (r *CreateExecutePlanWithClusterRequest) SetTimeInterval(value int) {
	r.TimeInterval = value
	r.QueryParams.Set("TimeInterval", strconv.Itoa(value))
}
func (r *CreateExecutePlanWithClusterRequest) GetTimeInterval() int {
	return r.TimeInterval
}
func (r *CreateExecutePlanWithClusterRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetStartTime() string {
	return r.StartTime
}
func (r *CreateExecutePlanWithClusterRequest) SetTimeUnit(value string) {
	r.TimeUnit = value
	r.QueryParams.Set("TimeUnit", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetTimeUnit() string {
	return r.TimeUnit
}
func (r *CreateExecutePlanWithClusterRequest) SetJobId(value string) {
	r.JobId = value
	r.QueryParams.Set("JobId", value)
}
func (r *CreateExecutePlanWithClusterRequest) GetJobId() string {
	return r.JobId
}

func (r *CreateExecutePlanWithClusterRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateExecutePlanWithCluster")
	r.SetProduct(Product)
}

type CreateExecutePlanWithClusterResponse struct {
}

func CreateExecutePlanWithCluster(req *CreateExecutePlanWithClusterRequest, accessId, accessSecret string) (*CreateExecutePlanWithClusterResponse, error) {
	var pResponse CreateExecutePlanWithClusterResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
