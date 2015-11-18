package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListGroupsRequest struct {
	RpcRequest
	Marker   string
	MaxItems int
}

func (r *ListGroupsRequest) SetMarker(value string) {
	r.Marker = value
	r.QueryParams.Set("Marker", value)
}
func (r *ListGroupsRequest) GetMarker() string {
	return r.Marker
}
func (r *ListGroupsRequest) SetMaxItems(value int) {
	r.MaxItems = value
	r.QueryParams.Set("MaxItems", strconv.Itoa(value))
}
func (r *ListGroupsRequest) GetMaxItems() int {
	return r.MaxItems
}

func (r *ListGroupsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListGroups")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListGroupsResponse struct {
	IsTruncated bool   `xml:"IsTruncated" json:"IsTruncated"`
	Marker      string `xml:"Marker" json:"Marker"`
	Groups      struct {
		Group []struct {
			GroupName  string `xml:"GroupName" json:"GroupName"`
			Comments   string `xml:"Comments" json:"Comments"`
			CreateDate string `xml:"CreateDate" json:"CreateDate"`
			UpdateDate string `xml:"UpdateDate" json:"UpdateDate"`
		} `xml:"Group" json:"Group"`
	} `xml:"Groups" json:"Groups"`
}

func ListGroups(req *ListGroupsRequest, accessId, accessSecret string) (*ListGroupsResponse, error) {
	var pResponse ListGroupsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
