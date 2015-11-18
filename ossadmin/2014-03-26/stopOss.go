package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type stopOssRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *stopOssRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *stopOssRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *stopOssRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *stopOssRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *stopOssRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *stopOssRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *stopOssRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *stopOssRequest) Getregion() string {
	return r.region
}
func (r *stopOssRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *stopOssRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *stopOssRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("stopOss")
	r.SetProduct(Product)
}

type stopOssResponse struct {
	Code           string `xml:"Code" json:"Code"`
	Message        string `xml:"Message" json:"Message"`
	Success        bool   `xml:"Success" json:"Success"`
	aliUid         int    `xml:"aliUid" json:"aliUid"`
	instanceId     string `xml:"instanceId" json:"instanceId"`
	instacneStatus string `xml:"instacneStatus" json:"instacneStatus"`
	instanceName   string `xml:"instanceName" json:"instanceName"`
	startTime      string `xml:"startTime" json:"startTime"`
	endTime        string `xml:"endTime" json:"endTime"`
}

func stopOss(req *stopOssRequest, accessId, accessSecret string) (*stopOssResponse, error) {
	var pResponse stopOssResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
