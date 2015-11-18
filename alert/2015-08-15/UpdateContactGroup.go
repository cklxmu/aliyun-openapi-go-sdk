package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateContactGroupRequest struct {
	RoaRequest
	ProjectName  string
	GroupName    string
	ContactGroup string
}

func (r *UpdateContactGroupRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateContactGroupRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateContactGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.PathParams.Set("GroupName", value)
}
func (r *UpdateContactGroupRequest) GetGroupName() string {
	return r.GroupName
}
func (r *UpdateContactGroupRequest) SetContactGroup(value string) {
	r.ContactGroup = value
	r.QueryParams.Set("ContactGroup", value)
}
func (r *UpdateContactGroupRequest) GetContactGroup() string {
	return r.ContactGroup
}

func (r *UpdateContactGroupRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/groups/GroupName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateContactGroupResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateContactGroup(req *UpdateContactGroupRequest, accessId, accessSecret string) (*UpdateContactGroupResponse, error) {
	var pResponse UpdateContactGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
