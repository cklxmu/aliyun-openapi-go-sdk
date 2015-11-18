package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateGroupRequest struct {
	RpcRequest
	GroupName    string
	NewGroupName string
	NewComments  string
}

func (r *UpdateGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *UpdateGroupRequest) GetGroupName() string {
	return r.GroupName
}
func (r *UpdateGroupRequest) SetNewGroupName(value string) {
	r.NewGroupName = value
	r.QueryParams.Set("NewGroupName", value)
}
func (r *UpdateGroupRequest) GetNewGroupName() string {
	return r.NewGroupName
}
func (r *UpdateGroupRequest) SetNewComments(value string) {
	r.NewComments = value
	r.QueryParams.Set("NewComments", value)
}
func (r *UpdateGroupRequest) GetNewComments() string {
	return r.NewComments
}

func (r *UpdateGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UpdateGroupResponse struct {
	Group struct {
		GroupName  string `xml:"GroupName" json:"GroupName"`
		Comments   string `xml:"Comments" json:"Comments"`
		CreateDate string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate string `xml:"UpdateDate" json:"UpdateDate"`
	} `xml:"Group" json:"Group"`
}

func UpdateGroup(req *UpdateGroupRequest, accessId, accessSecret string) (*UpdateGroupResponse, error) {
	var pResponse UpdateGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
