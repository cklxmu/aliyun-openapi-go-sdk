package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PutBucketLimitRequest struct {
	RpcRequest
	uid          string
	bid          string
	BucketLimit  int
	OwnerAccount string
}

func (r *PutBucketLimitRequest) Setuid(value string) {
	r.uid = value
	r.QueryParams.Set("uid", value)
}
func (r *PutBucketLimitRequest) Getuid() string {
	return r.uid
}
func (r *PutBucketLimitRequest) Setbid(value string) {
	r.bid = value
	r.QueryParams.Set("bid", value)
}
func (r *PutBucketLimitRequest) Getbid() string {
	return r.bid
}
func (r *PutBucketLimitRequest) SetBucketLimit(value int) {
	r.BucketLimit = value
	r.QueryParams.Set("BucketLimit", strconv.Itoa(value))
}
func (r *PutBucketLimitRequest) GetBucketLimit() int {
	return r.BucketLimit
}
func (r *PutBucketLimitRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *PutBucketLimitRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *PutBucketLimitRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutBucketLimit")
	r.SetProduct(Product)
}

type PutBucketLimitResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success bool   `xml:"Success" json:"Success"`
}

func PutBucketLimit(req *PutBucketLimitRequest, accessId, accessSecret string) (*PutBucketLimitResponse, error) {
	var pResponse PutBucketLimitResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
