package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetBucketVipsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	BucketName           string
}

func (r *GetBucketVipsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *GetBucketVipsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *GetBucketVipsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *GetBucketVipsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *GetBucketVipsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *GetBucketVipsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *GetBucketVipsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GetBucketVipsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *GetBucketVipsRequest) SetBucketName(value string) {
	r.BucketName = value
	r.QueryParams.Set("BucketName", value)
}
func (r *GetBucketVipsRequest) GetBucketName() string {
	return r.BucketName
}

func (r *GetBucketVipsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetBucketVips")
	r.SetProduct(Product)
}

type GetBucketVipsResponse struct {
	requestId string `xml:"requestId" json:"requestId"`
	vipList   struct {
		vip []string `xml:"vip" json:"vip"`
	} `xml:"vipList" json:"vipList"`
}

func GetBucketVips(req *GetBucketVipsRequest, accessId, accessSecret string) (*GetBucketVipsResponse, error) {
	var pResponse GetBucketVipsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
