package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateRoleRequest struct {
	RpcRequest
	RoleName                    string
	NewAssumeRolePolicyDocument string
}

func (r *UpdateRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *UpdateRoleRequest) GetRoleName() string {
	return r.RoleName
}
func (r *UpdateRoleRequest) SetNewAssumeRolePolicyDocument(value string) {
	r.NewAssumeRolePolicyDocument = value
	r.QueryParams.Set("NewAssumeRolePolicyDocument", value)
}
func (r *UpdateRoleRequest) GetNewAssumeRolePolicyDocument() string {
	return r.NewAssumeRolePolicyDocument
}

func (r *UpdateRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UpdateRoleResponse struct {
	Role struct {
		RoleId                   string `xml:"RoleId" json:"RoleId"`
		RoleName                 string `xml:"RoleName" json:"RoleName"`
		Arn                      string `xml:"Arn" json:"Arn"`
		Description              string `xml:"Description" json:"Description"`
		AssumeRolePolicyDocument string `xml:"AssumeRolePolicyDocument" json:"AssumeRolePolicyDocument"`
		CreateDate               string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate               string `xml:"UpdateDate" json:"UpdateDate"`
	} `xml:"Role" json:"Role"`
}

func UpdateRole(req *UpdateRoleRequest, accessId, accessSecret string) (*UpdateRoleResponse, error) {
	var pResponse UpdateRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
