package crm

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryCustomerLabelRequest struct {
	RpcRequest
	LabelSeries string
}

func (r *QueryCustomerLabelRequest) SetLabelSeries(value string) {
	r.LabelSeries = value
	r.QueryParams.Set("LabelSeries", value)
}
func (r *QueryCustomerLabelRequest) GetLabelSeries() string {
	return r.LabelSeries
}

func (r *QueryCustomerLabelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryCustomerLabel")
	r.SetProduct(Product)
}

type QueryCustomerLabelResponse struct {
	Success bool   `xml:"Success" json:"Success"`
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Data    struct {
		CustomerLabel []struct {
			Label       string `xml:"Label" json:"Label"`
			LabelSeries string `xml:"LabelSeries" json:"LabelSeries"`
		} `xml:"CustomerLabel" json:"CustomerLabel"`
	} `xml:"Data" json:"Data"`
}

func QueryCustomerLabel(req *QueryCustomerLabelRequest, accessId, accessSecret string) (*QueryCustomerLabelResponse, error) {
	var pResponse QueryCustomerLabelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
