package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetRoleRequest struct {
	RpcRequest
	RoleName string
}

func (r *GetRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *GetRoleRequest) GetRoleName() string {
	return r.RoleName
}

func (r *GetRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetRoleResponse struct {
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

func GetRole(req *GetRoleRequest, accessId, accessSecret string) (*GetRoleResponse, error) {
	var pResponse GetRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
