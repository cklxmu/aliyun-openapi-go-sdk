package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateGroupRequest struct {
	RpcRequest
	GroupName string
	Comments  string
}

func (r *CreateGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *CreateGroupRequest) GetGroupName() string {
	return r.GroupName
}
func (r *CreateGroupRequest) SetComments(value string) {
	r.Comments = value
	r.QueryParams.Set("Comments", value)
}
func (r *CreateGroupRequest) GetComments() string {
	return r.Comments
}

func (r *CreateGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateGroupResponse struct {
	Group struct {
		GroupName  string `xml:"GroupName" json:"GroupName"`
		Comments   string `xml:"Comments" json:"Comments"`
		CreateDate string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"Group" json:"Group"`
}

func CreateGroup(req *CreateGroupRequest, accessId, accessSecret string) (*CreateGroupResponse, error) {
	var pResponse CreateGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
