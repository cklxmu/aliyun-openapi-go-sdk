package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteContactGroupRequest struct {
	RoaRequest
	ProjectName string
	GroupName   string
}

func (r *DeleteContactGroupRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteContactGroupRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteContactGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.PathParams.Set("GroupName", value)
}
func (r *DeleteContactGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *DeleteContactGroupRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/groups/GroupName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteContactGroupResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteContactGroup(req *DeleteContactGroupRequest, accessId, accessSecret string) (*DeleteContactGroupResponse, error) {
	var pResponse DeleteContactGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
