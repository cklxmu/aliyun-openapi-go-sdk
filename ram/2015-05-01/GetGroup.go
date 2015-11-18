package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetGroupRequest struct {
	RpcRequest
	GroupName string
}

func (r *GetGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *GetGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *GetGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetGroupResponse struct {
	Group struct {
		GroupName  string `xml:"GroupName" json:"GroupName"`
		Comments   string `xml:"Comments" json:"Comments"`
		CreateDate string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate string `xml:"UpdateDate" json:"UpdateDate"`
	} `xml:"Group" json:"Group"`
}

func GetGroup(req *GetGroupRequest, accessId, accessSecret string) (*GetGroupResponse, error) {
	var pResponse GetGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
