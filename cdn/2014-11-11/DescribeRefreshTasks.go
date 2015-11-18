package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRefreshTasksRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	TaskId               string
	ObjectPath           string
	PageNumber           int
	PageSize             int
}

func (r *DescribeRefreshTasksRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeRefreshTasksRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeRefreshTasksRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeRefreshTasksRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeRefreshTasksRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeRefreshTasksRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeRefreshTasksRequest) SetTaskId(value string) {
	r.TaskId = value
	r.QueryParams.Set("TaskId", value)
}
func (r *DescribeRefreshTasksRequest) GetTaskId() string {
	return r.TaskId
}
func (r *DescribeRefreshTasksRequest) SetObjectPath(value string) {
	r.ObjectPath = value
	r.QueryParams.Set("ObjectPath", value)
}
func (r *DescribeRefreshTasksRequest) GetObjectPath() string {
	return r.ObjectPath
}
func (r *DescribeRefreshTasksRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeRefreshTasksRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeRefreshTasksRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeRefreshTasksRequest) GetPageSize() int {
	return r.PageSize
}

func (r *DescribeRefreshTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeRefreshTasks")
	r.SetProduct(Product)
}

type DescribeRefreshTasksResponse struct {
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	Tasks      struct {
		CDNTask []struct {
			TaskId       string `xml:"TaskId" json:"TaskId"`
			ObjectPath   string `xml:"ObjectPath" json:"ObjectPath"`
			Status       string `xml:"Status" json:"Status"`
			CreationTime string `xml:"CreationTime" json:"CreationTime"`
		} `xml:"CDNTask" json:"CDNTask"`
	} `xml:"Tasks" json:"Tasks"`
}

func DescribeRefreshTasks(req *DescribeRefreshTasksRequest, accessId, accessSecret string) (*DescribeRefreshTasksResponse, error) {
	var pResponse DescribeRefreshTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
