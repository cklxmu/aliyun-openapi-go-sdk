package oms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetUserDataRequest struct {
	RpcRequest
	OwnerId      int
	OwnerAccount string
	ProductName  string
	DataType     string
	StartTime    string
	EndTime      string
	NextToken    string
	MaxResult    int
}

func (r *GetUserDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *GetUserDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *GetUserDataRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GetUserDataRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *GetUserDataRequest) SetProductName(value string) {
	r.ProductName = value
	r.QueryParams.Set("ProductName", value)
}
func (r *GetUserDataRequest) GetProductName() string {
	return r.ProductName
}
func (r *GetUserDataRequest) SetDataType(value string) {
	r.DataType = value
	r.QueryParams.Set("DataType", value)
}
func (r *GetUserDataRequest) GetDataType() string {
	return r.DataType
}
func (r *GetUserDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *GetUserDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *GetUserDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *GetUserDataRequest) GetEndTime() string {
	return r.EndTime
}
func (r *GetUserDataRequest) SetNextToken(value string) {
	r.NextToken = value
	r.QueryParams.Set("NextToken", value)
}
func (r *GetUserDataRequest) GetNextToken() string {
	return r.NextToken
}
func (r *GetUserDataRequest) SetMaxResult(value int) {
	r.MaxResult = value
	r.QueryParams.Set("MaxResult", strconv.Itoa(value))
}
func (r *GetUserDataRequest) GetMaxResult() int {
	return r.MaxResult
}

func (r *GetUserDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetUserData")
	r.SetProduct(Product)
}

type GetUserDataResponse struct {
	ProductName string `xml:"ProductName" json:"ProductName"`
	DataType    string `xml:"DataType" json:"DataType"`
	NextToken   string `xml:"NextToken" json:"NextToken"`
	DataList    struct {
		Data []struct {
			DataItems struct {
				DataItem []struct {
					Name  string `xml:"Name" json:"Name"`
					Value string `xml:"Value" json:"Value"`
				} `xml:"DataItem" json:"DataItem"`
			} `xml:"DataItems" json:"DataItems"`
		} `xml:"Data" json:"Data"`
	} `xml:"DataList" json:"DataList"`
}

func GetUserData(req *GetUserDataRequest, accessId, accessSecret string) (*GetUserDataResponse, error) {
	var pResponse GetUserDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
