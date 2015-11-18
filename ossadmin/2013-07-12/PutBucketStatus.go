package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
)

type PutBucketStatusRequest struct {
	RpcRequest
	uid          string
	bid          string
	Bucket       string
	Status       string
	OwnerAccount string
}

func (r *PutBucketStatusRequest) Setuid(value string) {
	r.uid = value
	r.QueryParams.Set("uid", value)
}
func (r *PutBucketStatusRequest) Getuid() string {
	return r.uid
}
func (r *PutBucketStatusRequest) Setbid(value string) {
	r.bid = value
	r.QueryParams.Set("bid", value)
}
func (r *PutBucketStatusRequest) Getbid() string {
	return r.bid
}
func (r *PutBucketStatusRequest) SetBucket(value string) {
	r.Bucket = value
	r.QueryParams.Set("Bucket", value)
}
func (r *PutBucketStatusRequest) GetBucket() string {
	return r.Bucket
}
func (r *PutBucketStatusRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *PutBucketStatusRequest) GetStatus() string {
	return r.Status
}
func (r *PutBucketStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *PutBucketStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *PutBucketStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutBucketStatus")
	r.SetProduct(Product)
}

type PutBucketStatusResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success bool   `xml:"Success" json:"Success"`
}

func PutBucketStatus(req *PutBucketStatusRequest, accessId, accessSecret string) (*PutBucketStatusResponse, error) {
	var pResponse PutBucketStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
