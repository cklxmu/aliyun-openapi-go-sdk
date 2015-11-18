package crm

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteLabelRequest struct {
	RpcRequest
	PK          string
	LabelName   string
	LabelSeries string
}

func (r *DeleteLabelRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *DeleteLabelRequest) GetPK() string {
	return r.PK
}
func (r *DeleteLabelRequest) SetLabelName(value string) {
	r.LabelName = value
	r.QueryParams.Set("LabelName", value)
}
func (r *DeleteLabelRequest) GetLabelName() string {
	return r.LabelName
}
func (r *DeleteLabelRequest) SetLabelSeries(value string) {
	r.LabelSeries = value
	r.QueryParams.Set("LabelSeries", value)
}
func (r *DeleteLabelRequest) GetLabelSeries() string {
	return r.LabelSeries
}

func (r *DeleteLabelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteLabel")
	r.SetProduct(Product)
}

type DeleteLabelResponse struct {
	Result string `xml:"Result" json:"Result"`
}

func DeleteLabel(req *DeleteLabelRequest, accessId, accessSecret string) (*DeleteLabelResponse, error) {
	var pResponse DeleteLabelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
