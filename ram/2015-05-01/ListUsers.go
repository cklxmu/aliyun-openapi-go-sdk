package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListUsersRequest struct {
	RpcRequest
	Marker   string
	MaxItems int
}

func (r *ListUsersRequest) SetMarker(value string) {
	r.Marker = value
	r.QueryParams.Set("Marker", value)
}
func (r *ListUsersRequest) GetMarker() string {
	return r.Marker
}
func (r *ListUsersRequest) SetMaxItems(value int) {
	r.MaxItems = value
	r.QueryParams.Set("MaxItems", strconv.Itoa(value))
}
func (r *ListUsersRequest) GetMaxItems() int {
	return r.MaxItems
}

func (r *ListUsersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListUsers")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListUsersResponse struct {
	IsTruncated bool   `xml:"IsTruncated" json:"IsTruncated"`
	Marker      string `xml:"Marker" json:"Marker"`
	Users       struct {
		User []struct {
			UserId      string `xml:"UserId" json:"UserId"`
			UserName    string `xml:"UserName" json:"UserName"`
			DisplayName string `xml:"DisplayName" json:"DisplayName"`
			MobilePhone string `xml:"MobilePhone" json:"MobilePhone"`
			Email       string `xml:"Email" json:"Email"`
			Comments    string `xml:"Comments" json:"Comments"`
			CreateDate  string `xml:"CreateDate" json:"CreateDate"`
			UpdateDate  string `xml:"UpdateDate" json:"UpdateDate"`
		} `xml:"User" json:"User"`
	} `xml:"Users" json:"Users"`
}

func ListUsers(req *ListUsersRequest, accessId, accessSecret string) (*ListUsersResponse, error) {
	var pResponse ListUsersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
