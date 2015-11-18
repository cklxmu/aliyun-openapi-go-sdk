package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeUserDomainsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageSize             int
	PageNumber           int
}

func (r *DescribeUserDomainsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeUserDomainsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeUserDomainsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeUserDomainsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeUserDomainsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeUserDomainsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeUserDomainsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeUserDomainsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeUserDomainsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeUserDomainsRequest) GetPageNumber() int {
	return r.PageNumber
}

func (r *DescribeUserDomainsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeUserDomains")
	r.SetProduct(Product)
}

type DescribeUserDomainsResponse struct {
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	Domains    struct {
		PageData []struct {
			DomainName   string `xml:"DomainName" json:"DomainName"`
			Cname        string `xml:"Cname" json:"Cname"`
			CdnType      string `xml:"CdnType" json:"CdnType"`
			DomainStatus string `xml:"DomainStatus" json:"DomainStatus"`
			GmtCreated   string `xml:"GmtCreated" json:"GmtCreated"`
			GmtModified  string `xml:"GmtModified" json:"GmtModified"`
		} `xml:"PageData" json:"PageData"`
	} `xml:"Domains" json:"Domains"`
}

func DescribeUserDomains(req *DescribeUserDomainsRequest, accessId, accessSecret string) (*DescribeUserDomainsResponse, error) {
	var pResponse DescribeUserDomainsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
