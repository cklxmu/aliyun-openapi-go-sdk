package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDDLTaskRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	TaskId         string
}

func (r *DescribeDDLTaskRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeDDLTaskRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DescribeDDLTaskRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DescribeDDLTaskRequest) GetDbName() string {
	return r.DbName
}
func (r *DescribeDDLTaskRequest) SetTaskId(value string) {
	r.TaskId = value
	r.QueryParams.Set("TaskId", value)
}
func (r *DescribeDDLTaskRequest) GetTaskId() string {
	return r.TaskId
}

func (r *DescribeDDLTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDDLTask")
	r.SetProduct(Product)
}

type DescribeDDLTaskResponse struct {
	Data struct {
		RequestId   string `xml:"RequestId" json:"RequestId"`
		TargetId    string `xml:"TargetId" json:"TargetId"`
		TaskDetail  string `xml:"TaskDetail" json:"TaskDetail"`
		TaskStatus  int    `xml:"TaskStatus" json:"TaskStatus"`
		TaskPhase   string `xml:"TaskPhase" json:"TaskPhase"`
		TaskType    int    `xml:"TaskType" json:"TaskType"`
		TaskName    string `xml:"TaskName" json:"TaskName"`
		GmtCreate   int    `xml:"GmtCreate" json:"GmtCreate"`
		AllowCancel string `xml:"AllowCancel" json:"AllowCancel"`
		ErrMsg      string `xml:"ErrMsg" json:"ErrMsg"`
	} `xml:"Data" json:"Data"`
}

func DescribeDDLTask(req *DescribeDDLTaskRequest, accessId, accessSecret string) (*DescribeDDLTaskResponse, error) {
	var pResponse DescribeDDLTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
