package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeExecutePlanRequest struct {
	RpcRequest
	Id int
}

func (r *DescribeExecutePlanRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *DescribeExecutePlanRequest) GetId() int {
	return r.Id
}

func (r *DescribeExecutePlanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeExecutePlan")
	r.SetProduct(Product)
}

type DescribeExecutePlanResponse struct {
	Id              int    `xml:"Id" json:"Id"`
	Name            string `xml:"Name" json:"Name"`
	Strategy        int    `xml:"Strategy" json:"Strategy"`
	TimeInterval    int    `xml:"TimeInterval" json:"TimeInterval"`
	StartTime       string `xml:"StartTime" json:"StartTime"`
	TimeUnit        string `xml:"TimeUnit" json:"TimeUnit"`
	IsCreateCluster bool   `xml:"IsCreateCluster" json:"IsCreateCluster"`
	JobInfos        struct {
		JobInfo []struct {
			Id int `xml:"Id" json:"Id"`
		} `xml:"JobInfo" json:"JobInfo"`
	} `xml:"JobInfos" json:"JobInfos"`
	ClusterInfo struct {
		ClusterId         int    `xml:"ClusterId" json:"ClusterId"`
		BizId             string `xml:"BizId" json:"BizId"`
		ClusterName       string `xml:"ClusterName" json:"ClusterName"`
		UtcStartTime      string `xml:"UtcStartTime" json:"UtcStartTime"`
		UtcStopTime       string `xml:"UtcStopTime" json:"UtcStopTime"`
		TimeOutEnable     int    `xml:"TimeOutEnable" json:"TimeOutEnable"`
		TimeOutTime       string `xml:"TimeOutTime" json:"TimeOutTime"`
		TimeOutWarningWay int    `xml:"TimeOutWarningWay" json:"TimeOutWarningWay"`
		TimeOutOperate    int    `xml:"TimeOutOperate" json:"TimeOutOperate"`
		ReleaseSetting    int    `xml:"ReleaseSetting" json:"ReleaseSetting"`
		IsOpenOssLog      bool   `xml:"IsOpenOssLog" json:"IsOpenOssLog"`
		Status            int    `xml:"Status" json:"Status"`
		LastJobStatus     int    `xml:"LastJobStatus" json:"LastJobStatus"`
		RunningTime       int    `xml:"RunningTime" json:"RunningTime"`
		FailReason        string `xml:"FailReason" json:"FailReason"`
		OssPath           string `xml:"OssPath" json:"OssPath"`
		EmrRole4ECS       string `xml:"EmrRole4ECS" json:"EmrRole4ECS"`
		EmrRole4Oss       string `xml:"EmrRole4Oss" json:"EmrRole4Oss"`
		IsOpenPublicIp    bool   `xml:"IsOpenPublicIp" json:"IsOpenPublicIp"`
	} `xml:"ClusterInfo" json:"ClusterInfo"`
	SoftwareInfo struct {
		EmrVer      string `xml:"EmrVer" json:"EmrVer"`
		ClusterType string `xml:"ClusterType" json:"ClusterType"`
		Softwares   string `xml:"Softwares" json:"Softwares"`
	} `xml:"SoftwareInfo" json:"SoftwareInfo"`
	EcsInfo struct {
		RegionId        string `xml:"RegionId" json:"RegionId"`
		ZoneId          string `xml:"ZoneId" json:"ZoneId"`
		ImageId         string `xml:"ImageId" json:"ImageId"`
		ImageName       string `xml:"ImageName" json:"ImageName"`
		SparkVersion    string `xml:"SparkVersion" json:"SparkVersion"`
		SecurityGroupId string `xml:"SecurityGroupId" json:"SecurityGroupId"`
		TotalCount      int    `xml:"TotalCount" json:"TotalCount"`
		MasterCount     int    `xml:"MasterCount" json:"MasterCount"`
		SlaveCount      int    `xml:"SlaveCount" json:"SlaveCount"`
		EcsRoles        struct {
			EcsRole []struct {
				IsMaster       bool   `xml:"IsMaster" json:"IsMaster"`
				InstanceType   string `xml:"InstanceType" json:"InstanceType"`
				Payment        string `xml:"Payment" json:"Payment"`
				CpuCore        string `xml:"CpuCore" json:"CpuCore"`
				MemoryCapacity string `xml:"MemoryCapacity" json:"MemoryCapacity"`
				DiskType       int    `xml:"DiskType" json:"DiskType"`
				DiskCapacity   int    `xml:"DiskCapacity" json:"DiskCapacity"`
				BandWidth      string `xml:"BandWidth" json:"BandWidth"`
				NetPayType     string `xml:"NetPayType" json:"NetPayType"`
				EcsPayType     string `xml:"EcsPayType" json:"EcsPayType"`
				Nodes          struct {
					Node []struct {
						InstanceId string `xml:"InstanceId" json:"InstanceId"`
						PubIp      string `xml:"PubIp" json:"PubIp"`
						InnerIp    string `xml:"InnerIp" json:"InnerIp"`
						DiskInfo   string `xml:"DiskInfo" json:"DiskInfo"`
					} `xml:"Node" json:"Node"`
				} `xml:"Nodes" json:"Nodes"`
			} `xml:"EcsRole" json:"EcsRole"`
		} `xml:"EcsRoles" json:"EcsRoles"`
	} `xml:"EcsInfo" json:"EcsInfo"`
}

func DescribeExecutePlan(req *DescribeExecutePlanRequest, accessId, accessSecret string) (*DescribeExecutePlanResponse, error) {
	var pResponse DescribeExecutePlanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
