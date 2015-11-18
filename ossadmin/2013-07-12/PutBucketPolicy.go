package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PutBucketPolicyRequest struct {
	RpcRequest
	uid                string
	bid                string
	BucketName         string
	IamPolicy          string
	DisallowEmptyRefer bool
	EnableDualCluster  bool
	ErrorFile          string
	IndexFile          string
	Location           string
	LogBucket          string
	LogPrefix          string
	WhiteReferList     string
	OwnerAccount       string
}

func (r *PutBucketPolicyRequest) Setuid(value string) {
	r.uid = value
	r.QueryParams.Set("uid", value)
}
func (r *PutBucketPolicyRequest) Getuid() string {
	return r.uid
}
func (r *PutBucketPolicyRequest) Setbid(value string) {
	r.bid = value
	r.QueryParams.Set("bid", value)
}
func (r *PutBucketPolicyRequest) Getbid() string {
	return r.bid
}
func (r *PutBucketPolicyRequest) SetBucketName(value string) {
	r.BucketName = value
	r.QueryParams.Set("BucketName", value)
}
func (r *PutBucketPolicyRequest) GetBucketName() string {
	return r.BucketName
}
func (r *PutBucketPolicyRequest) SetIamPolicy(value string) {
	r.IamPolicy = value
	r.QueryParams.Set("IamPolicy", value)
}
func (r *PutBucketPolicyRequest) GetIamPolicy() string {
	return r.IamPolicy
}
func (r *PutBucketPolicyRequest) SetDisallowEmptyRefer(value bool) {
	r.DisallowEmptyRefer = value
	r.QueryParams.Set("DisallowEmptyRefer", strconv.FormatBool(value))
}
func (r *PutBucketPolicyRequest) GetDisallowEmptyRefer() bool {
	return r.DisallowEmptyRefer
}
func (r *PutBucketPolicyRequest) SetEnableDualCluster(value bool) {
	r.EnableDualCluster = value
	r.QueryParams.Set("EnableDualCluster", strconv.FormatBool(value))
}
func (r *PutBucketPolicyRequest) GetEnableDualCluster() bool {
	return r.EnableDualCluster
}
func (r *PutBucketPolicyRequest) SetErrorFile(value string) {
	r.ErrorFile = value
	r.QueryParams.Set("ErrorFile", value)
}
func (r *PutBucketPolicyRequest) GetErrorFile() string {
	return r.ErrorFile
}
func (r *PutBucketPolicyRequest) SetIndexFile(value string) {
	r.IndexFile = value
	r.QueryParams.Set("IndexFile", value)
}
func (r *PutBucketPolicyRequest) GetIndexFile() string {
	return r.IndexFile
}
func (r *PutBucketPolicyRequest) SetLocation(value string) {
	r.Location = value
	r.QueryParams.Set("Location", value)
}
func (r *PutBucketPolicyRequest) GetLocation() string {
	return r.Location
}
func (r *PutBucketPolicyRequest) SetLogBucket(value string) {
	r.LogBucket = value
	r.QueryParams.Set("LogBucket", value)
}
func (r *PutBucketPolicyRequest) GetLogBucket() string {
	return r.LogBucket
}
func (r *PutBucketPolicyRequest) SetLogPrefix(value string) {
	r.LogPrefix = value
	r.QueryParams.Set("LogPrefix", value)
}
func (r *PutBucketPolicyRequest) GetLogPrefix() string {
	return r.LogPrefix
}
func (r *PutBucketPolicyRequest) SetWhiteReferList(value string) {
	r.WhiteReferList = value
	r.QueryParams.Set("WhiteReferList", value)
}
func (r *PutBucketPolicyRequest) GetWhiteReferList() string {
	return r.WhiteReferList
}
func (r *PutBucketPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *PutBucketPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *PutBucketPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutBucketPolicy")
	r.SetProduct(Product)
}

type PutBucketPolicyResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success bool   `xml:"Success" json:"Success"`
}

func PutBucketPolicy(req *PutBucketPolicyRequest, accessId, accessSecret string) (*PutBucketPolicyResponse, error) {
	var pResponse PutBucketPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
