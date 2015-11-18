package crm

import (
	. "aliyun-openapi-go-sdk/core"
)

type AddLabelRequest struct {
	RpcRequest
	PK          string
	LabelName   string
	LabelSeries string
	EndTime     string
}

func (r *AddLabelRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *AddLabelRequest) GetPK() string {
	return r.PK
}
func (r *AddLabelRequest) SetLabelName(value string) {
	r.LabelName = value
	r.QueryParams.Set("LabelName", value)
}
func (r *AddLabelRequest) GetLabelName() string {
	return r.LabelName
}
func (r *AddLabelRequest) SetLabelSeries(value string) {
	r.LabelSeries = value
	r.QueryParams.Set("LabelSeries", value)
}
func (r *AddLabelRequest) GetLabelSeries() string {
	return r.LabelSeries
}
func (r *AddLabelRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *AddLabelRequest) GetEndTime() string {
	return r.EndTime
}

func (r *AddLabelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddLabel")
	r.SetProduct(Product)
}

type AddLabelResponse struct {
	Result string `xml:"Result" json:"Result"`
}

func AddLabel(req *AddLabelRequest, accessId, accessSecret string) (*AddLabelResponse, error) {
	var pResponse AddLabelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
