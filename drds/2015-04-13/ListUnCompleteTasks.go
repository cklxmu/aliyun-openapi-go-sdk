package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListUnCompleteTasksRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
}

func (r *ListUnCompleteTasksRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *ListUnCompleteTasksRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *ListUnCompleteTasksRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *ListUnCompleteTasksRequest) GetDbName() string {
	return r.DbName
}

func (r *ListUnCompleteTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListUnCompleteTasks")
	r.SetProduct(Product)
}

type ListUnCompleteTasksResponse struct {
	Data struct {
		task []struct {
			RequestId   string `xml:"RequestId" json:"RequestId"`
			TargetId    string `xml:"TargetId" json:"TargetId"`
			TaskDetail  string `xml:"TaskDetail" json:"TaskDetail"`
			taskStatus  int    `xml:"taskStatus" json:"taskStatus"`
			TaskPhase   string `xml:"TaskPhase" json:"TaskPhase"`
			TaskType    int    `xml:"TaskType" json:"TaskType"`
			TaskName    string `xml:"TaskName" json:"TaskName"`
			gmtCreate   int    `xml:"gmtCreate" json:"gmtCreate"`
			AllowCancel string `xml:"AllowCancel" json:"AllowCancel"`
			ErrMsg      string `xml:"ErrMsg" json:"ErrMsg"`
		} `xml:"task" json:"task"`
	} `xml:"Data" json:"Data"`
}

func ListUnCompleteTasks(req *ListUnCompleteTasksRequest, accessId, accessSecret string) (*ListUnCompleteTasksResponse, error) {
	var pResponse ListUnCompleteTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
