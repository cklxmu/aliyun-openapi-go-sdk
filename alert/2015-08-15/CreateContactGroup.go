package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateContactGroupRequest struct {
	RoaRequest
	ProjectName  string
	ContactGroup string
}

func (r *CreateContactGroupRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateContactGroupRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateContactGroupRequest) SetContactGroup(value string) {
	r.ContactGroup = value
	r.QueryParams.Set("ContactGroup", value)
}
func (r *CreateContactGroupRequest) GetContactGroup() string {
	return r.ContactGroup
}

func (r *CreateContactGroupRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/groups"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateContactGroupResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateContactGroup(req *CreateContactGroupRequest, accessId, accessSecret string) (*CreateContactGroupResponse, error) {
	var pResponse CreateContactGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
