package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetContactRequest struct {
	RoaRequest
	ProjectName string
	ContactName string
}

func (r *GetContactRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetContactRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetContactRequest) SetContactName(value string) {
	r.ContactName = value
	r.PathParams.Set("ContactName", value)
}
func (r *GetContactRequest) GetContactName() string {
	return r.ContactName
}

func (r *GetContactRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/contacts/ContactName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetContactResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetContact(req *GetContactRequest, accessId, accessSecret string) (*GetContactResponse, error) {
	var pResponse GetContactResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
