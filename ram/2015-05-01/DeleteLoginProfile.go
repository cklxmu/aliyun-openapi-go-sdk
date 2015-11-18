package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteLoginProfileRequest struct {
	RpcRequest
	UserName string
}

func (r *DeleteLoginProfileRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *DeleteLoginProfileRequest) GetUserName() string {
	return r.UserName
}

func (r *DeleteLoginProfileRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteLoginProfile")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteLoginProfileResponse struct {
}

func DeleteLoginProfile(req *DeleteLoginProfileRequest, accessId, accessSecret string) (*DeleteLoginProfileResponse, error) {
	var pResponse DeleteLoginProfileResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
