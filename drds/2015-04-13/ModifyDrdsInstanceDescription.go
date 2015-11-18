package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type ModifyDrdsInstanceDescriptionRequest struct {
	RpcRequest
	DrdsInstanceId string
	Description    string
}

func (r *ModifyDrdsInstanceDescriptionRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *ModifyDrdsInstanceDescriptionRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *ModifyDrdsInstanceDescriptionRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyDrdsInstanceDescriptionRequest) GetDescription() string {
	return r.Description
}

func (r *ModifyDrdsInstanceDescriptionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDrdsInstanceDescription")
	r.SetProduct(Product)
}

type ModifyDrdsInstanceDescriptionResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func ModifyDrdsInstanceDescription(req *ModifyDrdsInstanceDescriptionRequest, accessId, accessSecret string) (*ModifyDrdsInstanceDescriptionResponse, error) {
	var pResponse ModifyDrdsInstanceDescriptionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
