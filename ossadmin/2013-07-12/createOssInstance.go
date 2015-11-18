package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type createOssInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	aliUid               int
	OwnerAccount         string
}

func (r *createOssInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *createOssInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *createOssInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *createOssInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *createOssInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *createOssInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *createOssInstanceRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *createOssInstanceRequest) Getregion() string {
	return r.region
}
func (r *createOssInstanceRequest) SetaliUid(value int) {
	r.aliUid = value
	r.QueryParams.Set("aliUid", strconv.Itoa(value))
}
func (r *createOssInstanceRequest) GetaliUid() int {
	return r.aliUid
}
func (r *createOssInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *createOssInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *createOssInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("createOssInstance")
	r.SetProduct(Product)
}

type createOssInstanceResponse struct {
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

func createOssInstance(req *createOssInstanceRequest, accessId, accessSecret string) (*createOssInstanceResponse, error) {
	var pResponse createOssInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
