package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateClusterRequest struct {
	RpcRequest
	Name              string
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
	PayType           int
	Period            int
	EmrRole4ECS       string
	EmrRole4Oss       string
}

func (r *CreateClusterRequest) SetName(value string) {
	r.Name = value
	r.QueryParams.Set("Name", value)
}
func (r *CreateClusterRequest) GetName() string {
	return r.Name
}
func (r *CreateClusterRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateClusterRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateClusterRequest) SetLogEnable(value bool) {
	r.LogEnable = value
	r.QueryParams.Set("LogEnable", strconv.FormatBool(value))
}
func (r *CreateClusterRequest) GetLogEnable() bool {
	return r.LogEnable
}
func (r *CreateClusterRequest) SetLogPath(value string) {
	r.LogPath = value
	r.QueryParams.Set("LogPath", value)
}
func (r *CreateClusterRequest) GetLogPath() string {
	return r.LogPath
}
func (r *CreateClusterRequest) SetSecurityGroup(value string) {
	r.SecurityGroup = value
	r.QueryParams.Set("SecurityGroup", value)
}
func (r *CreateClusterRequest) GetSecurityGroup() string {
	return r.SecurityGroup
}
func (r *CreateClusterRequest) SetIsOpenPublicIp(value bool) {
	r.IsOpenPublicIp = value
	r.QueryParams.Set("IsOpenPublicIp", strconv.FormatBool(value))
}
func (r *CreateClusterRequest) GetIsOpenPublicIp() bool {
	return r.IsOpenPublicIp
}
func (r *CreateClusterRequest) SetSecurityGroupName(value string) {
	r.SecurityGroupName = value
	r.QueryParams.Set("SecurityGroupName", value)
}
func (r *CreateClusterRequest) GetSecurityGroupName() string {
	return r.SecurityGroupName
}
func (r *CreateClusterRequest) SetEmrVer(value string) {
	r.EmrVer = value
	r.QueryParams.Set("EmrVer", value)
}
func (r *CreateClusterRequest) GetEmrVer() string {
	return r.EmrVer
}
func (r *CreateClusterRequest) SetClusterType(value string) {
	r.ClusterType = value
	r.QueryParams.Set("ClusterType", value)
}
func (r *CreateClusterRequest) GetClusterType() string {
	return r.ClusterType
}
func (r *CreateClusterRequest) SetInstall(value string) {
	r.Install = value
	r.QueryParams.Set("Install", value)
}
func (r *CreateClusterRequest) GetInstall() string {
	return r.Install
}
func (r *CreateClusterRequest) SetMasterIndex(value int) {
	r.MasterIndex = value
	r.QueryParams.Set("MasterIndex", strconv.Itoa(value))
}
func (r *CreateClusterRequest) GetMasterIndex() int {
	return r.MasterIndex
}
func (r *CreateClusterRequest) SetEcsOrder(value string) {
	r.EcsOrder = value
	r.QueryParams.Set("EcsOrder", value)
}
func (r *CreateClusterRequest) GetEcsOrder() string {
	return r.EcsOrder
}
func (r *CreateClusterRequest) SetPayType(value int) {
	r.PayType = value
	r.QueryParams.Set("PayType", strconv.Itoa(value))
}
func (r *CreateClusterRequest) GetPayType() int {
	return r.PayType
}
func (r *CreateClusterRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *CreateClusterRequest) GetPeriod() int {
	return r.Period
}
func (r *CreateClusterRequest) SetEmrRole4ECS(value string) {
	r.EmrRole4ECS = value
	r.QueryParams.Set("EmrRole4ECS", value)
}
func (r *CreateClusterRequest) GetEmrRole4ECS() string {
	return r.EmrRole4ECS
}
func (r *CreateClusterRequest) SetEmrRole4Oss(value string) {
	r.EmrRole4Oss = value
	r.QueryParams.Set("EmrRole4Oss", value)
}
func (r *CreateClusterRequest) GetEmrRole4Oss() string {
	return r.EmrRole4Oss
}

func (r *CreateClusterRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateCluster")
	r.SetProduct(Product)
}

type CreateClusterResponse struct {
}

func CreateCluster(req *CreateClusterRequest, accessId, accessSecret string) (*CreateClusterResponse, error) {
	var pResponse CreateClusterResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
