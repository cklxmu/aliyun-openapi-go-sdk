package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type restartOssRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *restartOssRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *restartOssRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *restartOssRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *restartOssRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *restartOssRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *restartOssRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *restartOssRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *restartOssRequest) Getregion() string {
	return r.region
}
func (r *restartOssRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *restartOssRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *restartOssRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("restartOss")
	r.SetProduct(Product)
}

type restartOssResponse struct {
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

func restartOss(req *restartOssRequest, accessId, accessSecret string) (*restartOssResponse, error) {
	var pResponse restartOssResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
