package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCdnDomainBaseDetailRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
}

func (r *DescribeCdnDomainBaseDetailRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainBaseDetailRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCdnDomainBaseDetailRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCdnDomainBaseDetailRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCdnDomainBaseDetailRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainBaseDetailRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeCdnDomainBaseDetailRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DescribeCdnDomainBaseDetailRequest) GetDomainName() string {
	return r.DomainName
}

func (r *DescribeCdnDomainBaseDetailRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCdnDomainBaseDetail")
	r.SetProduct(Product)
}

type DescribeCdnDomainBaseDetailResponse struct {
	DomainBaseDetailModel struct {
		Cname        string `xml:"Cname" json:"Cname"`
		CdnType      string `xml:"CdnType" json:"CdnType"`
		DomainStatus string `xml:"DomainStatus" json:"DomainStatus"`
		SourceType   string `xml:"SourceType" json:"SourceType"`
		DomainName   string `xml:"DomainName" json:"DomainName"`
		Remark       string `xml:"Remark" json:"Remark"`
		GmtModified  string `xml:"GmtModified" json:"GmtModified"`
		GmtCreated   string `xml:"GmtCreated" json:"GmtCreated"`
		Sources      struct {
			Source []string `xml:"Source" json:"Source"`
		} `xml:"Sources" json:"Sources"`
	} `xml:"DomainBaseDetailModel" json:"DomainBaseDetailModel"`
}

func DescribeCdnDomainBaseDetail(req *DescribeCdnDomainBaseDetailRequest, accessId, accessSecret string) (*DescribeCdnDomainBaseDetailResponse, error) {
	var pResponse DescribeCdnDomainBaseDetailResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
