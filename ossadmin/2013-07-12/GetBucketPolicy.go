package ossadmin

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetBucketPolicyRequest struct {
	RpcRequest
	uid          string
	bid          string
	BucketName   string
	OwnerAccount string
}

func (r *GetBucketPolicyRequest) Setuid(value string) {
	r.uid = value
	r.QueryParams.Set("uid", value)
}
func (r *GetBucketPolicyRequest) Getuid() string {
	return r.uid
}
func (r *GetBucketPolicyRequest) Setbid(value string) {
	r.bid = value
	r.QueryParams.Set("bid", value)
}
func (r *GetBucketPolicyRequest) Getbid() string {
	return r.bid
}
func (r *GetBucketPolicyRequest) SetBucketName(value string) {
	r.BucketName = value
	r.QueryParams.Set("BucketName", value)
}
func (r *GetBucketPolicyRequest) GetBucketName() string {
	return r.BucketName
}
func (r *GetBucketPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GetBucketPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *GetBucketPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetBucketPolicy")
	r.SetProduct(Product)
}

type GetBucketPolicyResponse struct {
	DisallowEmptyRefer bool   `xml:"DisallowEmptyRefer" json:"DisallowEmptyRefer"`
	EnableDualCluster  bool   `xml:"EnableDualCluster" json:"EnableDualCluster"`
	ErrorFile          string `xml:"ErrorFile" json:"ErrorFile"`
	IndexFile          string `xml:"IndexFile" json:"IndexFile"`
	Location           string `xml:"Location" json:"Location"`
	LogBucket          string `xml:"LogBucket" json:"LogBucket"`
	LogPrefix          string `xml:"LogPrefix" json:"LogPrefix"`
	IamPolicy          string `xml:"IamPolicy" json:"IamPolicy"`
	WhiteReferList     struct {
		ReferList struct {
			string []string `xml:"string" json:"string"`
		} `xml:"ReferList" json:"ReferList"`
	} `xml:"WhiteReferList" json:"WhiteReferList"`
}

func GetBucketPolicy(req *GetBucketPolicyRequest, accessId, accessSecret string) (*GetBucketPolicyResponse, error) {
	var pResponse GetBucketPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
