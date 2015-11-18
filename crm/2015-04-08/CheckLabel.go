package crm

import (
	. "aliyun-openapi-go-sdk/core"
)

type CheckLabelRequest struct {
	RpcRequest
	PK          string
	LabelName   string
	LabelSeries string
}

func (r *CheckLabelRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *CheckLabelRequest) GetPK() string {
	return r.PK
}
func (r *CheckLabelRequest) SetLabelName(value string) {
	r.LabelName = value
	r.QueryParams.Set("LabelName", value)
}
func (r *CheckLabelRequest) GetLabelName() string {
	return r.LabelName
}
func (r *CheckLabelRequest) SetLabelSeries(value string) {
	r.LabelSeries = value
	r.QueryParams.Set("LabelSeries", value)
}
func (r *CheckLabelRequest) GetLabelSeries() string {
	return r.LabelSeries
}

func (r *CheckLabelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CheckLabel")
	r.SetProduct(Product)
}

type CheckLabelResponse struct {
	Result bool `xml:"Result" json:"Result"`
}

func CheckLabel(req *CheckLabelRequest, accessId, accessSecret string) (*CheckLabelResponse, error) {
	var pResponse CheckLabelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
