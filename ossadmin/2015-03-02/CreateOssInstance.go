package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateOssInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	region               string
	OwnerAccount         string
}

func (r *CreateOssInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateOssInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateOssInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateOssInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateOssInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateOssInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateOssInstanceRequest) Setregion(value string) {
	r.region = value
	r.QueryParams.Set("region", value)
}
func (r *CreateOssInstanceRequest) Getregion() string {
	return r.region
}
func (r *CreateOssInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateOssInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateOssInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateOssInstance")
	r.SetProduct(Product)
}

type CreateOssInstanceResponse struct {
	AliUid         int    `xml:"AliUid" json:"AliUid"`
	InstanceId     string `xml:"InstanceId" json:"InstanceId"`
	InstacneStatus string `xml:"InstacneStatus" json:"InstacneStatus"`
	InstanceName   string `xml:"InstanceName" json:"InstanceName"`
	StartTime      string `xml:"StartTime" json:"StartTime"`
	EndTime        string `xml:"EndTime" json:"EndTime"`
}

func CreateOssInstance(req *CreateOssInstanceRequest, accessId, accessSecret string) (*CreateOssInstanceResponse, error) {
	var pResponse CreateOssInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
