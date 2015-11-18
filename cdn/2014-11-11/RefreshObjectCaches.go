package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RefreshObjectCachesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ObjectPath           string
	ObjectType           string
}

func (r *RefreshObjectCachesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RefreshObjectCachesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RefreshObjectCachesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RefreshObjectCachesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RefreshObjectCachesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RefreshObjectCachesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RefreshObjectCachesRequest) SetObjectPath(value string) {
	r.ObjectPath = value
	r.QueryParams.Set("ObjectPath", value)
}
func (r *RefreshObjectCachesRequest) GetObjectPath() string {
	return r.ObjectPath
}
func (r *RefreshObjectCachesRequest) SetObjectType(value string) {
	r.ObjectType = value
	r.QueryParams.Set("ObjectType", value)
}
func (r *RefreshObjectCachesRequest) GetObjectType() string {
	return r.ObjectType
}

func (r *RefreshObjectCachesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RefreshObjectCaches")
	r.SetProduct(Product)
}

type RefreshObjectCachesResponse struct {
	RefreshTaskId string `xml:"RefreshTaskId" json:"RefreshTaskId"`
}

func RefreshObjectCaches(req *RefreshObjectCachesRequest, accessId, accessSecret string) (*RefreshObjectCachesResponse, error) {
	var pResponse RefreshObjectCachesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
