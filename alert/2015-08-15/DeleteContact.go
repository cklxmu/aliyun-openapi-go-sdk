package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteContactRequest struct {
	RoaRequest
	ProjectName string
	ContactName string
}

func (r *DeleteContactRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteContactRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteContactRequest) SetContactName(value string) {
	r.ContactName = value
	r.PathParams.Set("ContactName", value)
}
func (r *DeleteContactRequest) GetContactName() string {
	return r.ContactName
}

func (r *DeleteContactRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/contacts/ContactName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteContactResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteContact(req *DeleteContactRequest, accessId, accessSecret string) (*DeleteContactResponse, error) {
	var pResponse DeleteContactResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
