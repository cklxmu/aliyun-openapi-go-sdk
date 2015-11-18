package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PushObjectCacheRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ObjectPath           string
}

func (r *PushObjectCacheRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *PushObjectCacheRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *PushObjectCacheRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *PushObjectCacheRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *PushObjectCacheRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *PushObjectCacheRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *PushObjectCacheRequest) SetObjectPath(value string) {
	r.ObjectPath = value
	r.QueryParams.Set("ObjectPath", value)
}
func (r *PushObjectCacheRequest) GetObjectPath() string {
	return r.ObjectPath
}

func (r *PushObjectCacheRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PushObjectCache")
	r.SetProduct(Product)
}

type PushObjectCacheResponse struct {
	PushTaskId string `xml:"PushTaskId" json:"PushTaskId"`
}

func PushObjectCache(req *PushObjectCacheRequest, accessId, accessSecret string) (*PushObjectCacheResponse, error) {
	var pResponse PushObjectCacheResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
