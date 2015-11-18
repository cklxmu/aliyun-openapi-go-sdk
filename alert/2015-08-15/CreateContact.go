package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateContactRequest struct {
	RoaRequest
	ProjectName string
	Contact     string
}

func (r *CreateContactRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateContactRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateContactRequest) SetContact(value string) {
	r.Contact = value
	r.QueryParams.Set("Contact", value)
}
func (r *CreateContactRequest) GetContact() string {
	return r.Contact
}

func (r *CreateContactRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/contacts"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateContactResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateContact(req *CreateContactRequest, accessId, accessSecret string) (*CreateContactResponse, error) {
	var pResponse CreateContactResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
