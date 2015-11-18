package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListRolesRequest struct {
	RpcRequest
	Marker   string
	MaxItems int
}

func (r *ListRolesRequest) SetMarker(value string) {
	r.Marker = value
	r.QueryParams.Set("Marker", value)
}
func (r *ListRolesRequest) GetMarker() string {
	return r.Marker
}
func (r *ListRolesRequest) SetMaxItems(value int) {
	r.MaxItems = value
	r.QueryParams.Set("MaxItems", strconv.Itoa(value))
}
func (r *ListRolesRequest) GetMaxItems() int {
	return r.MaxItems
}

func (r *ListRolesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListRoles")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListRolesResponse struct {
	IsTruncated bool   `xml:"IsTruncated" json:"IsTruncated"`
	Marker      string `xml:"Marker" json:"Marker"`
	Roles       struct {
		Role []struct {
			RoleId      string `xml:"RoleId" json:"RoleId"`
			RoleName    string `xml:"RoleName" json:"RoleName"`
			Arn         string `xml:"Arn" json:"Arn"`
			Description string `xml:"Description" json:"Description"`
			CreateDate  string `xml:"CreateDate" json:"CreateDate"`
			UpdateDate  string `xml:"UpdateDate" json:"UpdateDate"`
		} `xml:"Role" json:"Role"`
	} `xml:"Roles" json:"Roles"`
}

func ListRoles(req *ListRolesRequest, accessId, accessSecret string) (*ListRolesResponse, error) {
	var pResponse ListRolesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
