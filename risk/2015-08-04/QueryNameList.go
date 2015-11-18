package risk

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryNameListRequest struct {
	RpcRequest
	Type      string
	DataType  string
	DataValue string
	QueryLike string
	Extend    string
}

func (r *QueryNameListRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *QueryNameListRequest) GetType() string {
	return r.Type
}
func (r *QueryNameListRequest) SetDataType(value string) {
	r.DataType = value
	r.QueryParams.Set("DataType", value)
}
func (r *QueryNameListRequest) GetDataType() string {
	return r.DataType
}
func (r *QueryNameListRequest) SetDataValue(value string) {
	r.DataValue = value
	r.QueryParams.Set("DataValue", value)
}
func (r *QueryNameListRequest) GetDataValue() string {
	return r.DataValue
}
func (r *QueryNameListRequest) SetQueryLike(value string) {
	r.QueryLike = value
	r.QueryParams.Set("QueryLike", value)
}
func (r *QueryNameListRequest) GetQueryLike() string {
	return r.QueryLike
}
func (r *QueryNameListRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *QueryNameListRequest) GetExtend() string {
	return r.Extend
}

func (r *QueryNameListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryNameList")
	r.SetProduct(Product)
}

type QueryNameListResponse struct {
	Code          string `xml:"Code" json:"Code"`
	RiskNameLists struct {
		RiskName []struct {
			Type      string `xml:"Type" json:"Type"`
			DataType  string `xml:"DataType" json:"DataType"`
			DataValue string `xml:"DataValue" json:"DataValue"`
		} `xml:"RiskName" json:"RiskName"`
	} `xml:"RiskNameLists" json:"RiskNameLists"`
}

func QueryNameList(req *QueryNameListRequest, accessId, accessSecret string) (*QueryNameListResponse, error) {
	var pResponse QueryNameListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
