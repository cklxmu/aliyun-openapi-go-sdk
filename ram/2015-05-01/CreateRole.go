package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateRoleRequest struct {
	RpcRequest
	RoleName                 string
	Description              string
	AssumeRolePolicyDocument string
}

func (r *CreateRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *CreateRoleRequest) GetRoleName() string {
	return r.RoleName
}
func (r *CreateRoleRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateRoleRequest) GetDescription() string {
	return r.Description
}
func (r *CreateRoleRequest) SetAssumeRolePolicyDocument(value string) {
	r.AssumeRolePolicyDocument = value
	r.QueryParams.Set("AssumeRolePolicyDocument", value)
}
func (r *CreateRoleRequest) GetAssumeRolePolicyDocument() string {
	return r.AssumeRolePolicyDocument
}

func (r *CreateRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateRoleResponse struct {
	Role struct {
		RoleId                   string `xml:"RoleId" json:"RoleId"`
		RoleName                 string `xml:"RoleName" json:"RoleName"`
		Arn                      string `xml:"Arn" json:"Arn"`
		Description              string `xml:"Description" json:"Description"`
		AssumeRolePolicyDocument string `xml:"AssumeRolePolicyDocument" json:"AssumeRolePolicyDocument"`
		CreateDate               string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"Role" json:"Role"`
}

func CreateRole(req *CreateRoleRequest, accessId, accessSecret string) (*CreateRoleResponse, error) {
	var pResponse CreateRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
