package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeCdnDomainDetailRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
}

func (r *DescribeCdnDomainDetailRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainDetailRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeCdnDomainDetailRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeCdnDomainDetailRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeCdnDomainDetailRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeCdnDomainDetailRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeCdnDomainDetailRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DescribeCdnDomainDetailRequest) GetDomainName() string {
	return r.DomainName
}

func (r *DescribeCdnDomainDetailRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCdnDomainDetail")
	r.SetProduct(Product)
}

type DescribeCdnDomainDetailResponse struct {
	GetDomainDetailModel struct {
		GmtCreated   string `xml:"GmtCreated" json:"GmtCreated"`
		GmtModified  string `xml:"GmtModified" json:"GmtModified"`
		SourceType   string `xml:"SourceType" json:"SourceType"`
		DomainStatus string `xml:"DomainStatus" json:"DomainStatus"`
		SourceType   string `xml:"SourceType" json:"SourceType"`
		CdnType      string `xml:"CdnType" json:"CdnType"`
		Cname        string `xml:"Cname" json:"Cname"`
		HttpsCname   string `xml:"HttpsCname" json:"HttpsCname"`
		Cname        string `xml:"Cname" json:"Cname"`
		DomainName   string `xml:"DomainName" json:"DomainName"`
		Remark       string `xml:"Remark" json:"Remark"`
		Sources      struct {
			Source []string `xml:"Source" json:"Source"`
		} `xml:"Sources" json:"Sources"`
	} `xml:"GetDomainDetailModel" json:"GetDomainDetailModel"`
}

func DescribeCdnDomainDetail(req *DescribeCdnDomainDetailRequest, accessId, accessSecret string) (*DescribeCdnDomainDetailResponse, error) {
	var pResponse DescribeCdnDomainDetailResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
