package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetContactGroupRequest struct {
	RoaRequest
	ProjectName string
	GroupName   string
}

func (r *GetContactGroupRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetContactGroupRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetContactGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.PathParams.Set("GroupName", value)
}
func (r *GetContactGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *GetContactGroupRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/groups/GroupName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetContactGroupResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetContactGroup(req *GetContactGroupRequest, accessId, accessSecret string) (*GetContactGroupResponse, error) {
	var pResponse GetContactGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
