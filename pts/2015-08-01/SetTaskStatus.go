package pts

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetTaskStatusRequest struct {
	RpcRequest
	Wskey  string
	Status string
}

func (r *SetTaskStatusRequest) SetWskey(value string) {
	r.Wskey = value
	r.QueryParams.Set("Wskey", value)
}
func (r *SetTaskStatusRequest) GetWskey() string {
	return r.Wskey
}
func (r *SetTaskStatusRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *SetTaskStatusRequest) GetStatus() string {
	return r.Status
}

func (r *SetTaskStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetTaskStatus")
	r.SetProduct(Product)
}

type SetTaskStatusResponse struct {
}

func SetTaskStatus(req *SetTaskStatusRequest, accessId, accessSecret string) (*SetTaskStatusResponse, error) {
	var pResponse SetTaskStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
