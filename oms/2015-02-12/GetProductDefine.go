package oms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetProductDefineRequest struct {
	RpcRequest
	ProductName string
	DataType    string
}

func (r *GetProductDefineRequest) SetProductName(value string) {
	r.ProductName = value
	r.QueryParams.Set("ProductName", value)
}
func (r *GetProductDefineRequest) GetProductName() string {
	return r.ProductName
}
func (r *GetProductDefineRequest) SetDataType(value string) {
	r.DataType = value
	r.QueryParams.Set("DataType", value)
}
func (r *GetProductDefineRequest) GetDataType() string {
	return r.DataType
}

func (r *GetProductDefineRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetProductDefine")
	r.SetProduct(Product)
}

type GetProductDefineResponse struct {
	ProductName string `xml:"ProductName" json:"ProductName"`
	DataType    string `xml:"DataType" json:"DataType"`
	ProductList struct {
		Product []struct {
			ProductName string `xml:"ProductName" json:"ProductName"`
			TypeList    struct {
				Type []struct {
					DataType string `xml:"DataType" json:"DataType"`
					Fields   struct {
						Field []string `xml:"Field" json:"Field"`
					} `xml:"Fields" json:"Fields"`
				} `xml:"Type" json:"Type"`
			} `xml:"TypeList" json:"TypeList"`
		} `xml:"Product" json:"Product"`
	} `xml:"ProductList" json:"ProductList"`
}

func GetProductDefine(req *GetProductDefineRequest, accessId, accessSecret string) (*GetProductDefineResponse, error) {
	var pResponse GetProductDefineResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
