package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateContactRequest struct {
	RoaRequest
	ProjectName string
	ContactName string
	Contact     string
}

func (r *UpdateContactRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateContactRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateContactRequest) SetContactName(value string) {
	r.ContactName = value
	r.PathParams.Set("ContactName", value)
}
func (r *UpdateContactRequest) GetContactName() string {
	return r.ContactName
}
func (r *UpdateContactRequest) SetContact(value string) {
	r.Contact = value
	r.QueryParams.Set("Contact", value)
}
func (r *UpdateContactRequest) GetContact() string {
	return r.Contact
}

func (r *UpdateContactRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/contacts/ContactName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateContactResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateContact(req *UpdateContactRequest, accessId, accessSecret string) (*UpdateContactResponse, error) {
	var pResponse UpdateContactResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
