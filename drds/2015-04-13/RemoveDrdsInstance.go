package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type RemoveDrdsInstanceRequest struct {
	RpcRequest
	DrdsInstanceId string
}

func (r *RemoveDrdsInstanceRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *RemoveDrdsInstanceRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}

func (r *RemoveDrdsInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveDrdsInstance")
	r.SetProduct(Product)
}

type RemoveDrdsInstanceResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func RemoveDrdsInstance(req *RemoveDrdsInstanceRequest, accessId, accessSecret string) (*RemoveDrdsInstanceResponse, error) {
	var pResponse RemoveDrdsInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
