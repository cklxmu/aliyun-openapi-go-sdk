package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteBackDoorFileRequest struct {
	RpcRequest
	JstOwnerId int
	InstanceId string
	Path       string
}

func (r *DeleteBackDoorFileRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *DeleteBackDoorFileRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *DeleteBackDoorFileRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DeleteBackDoorFileRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DeleteBackDoorFileRequest) SetPath(value string) {
	r.Path = value
	r.QueryParams.Set("Path", value)
}
func (r *DeleteBackDoorFileRequest) GetPath() string {
	return r.Path
}

func (r *DeleteBackDoorFileRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteBackDoorFile")
	r.SetProduct(Product)
}

type DeleteBackDoorFileResponse struct {
}

func DeleteBackDoorFile(req *DeleteBackDoorFileRequest, accessId, accessSecret string) (*DeleteBackDoorFileResponse, error) {
	var pResponse DeleteBackDoorFileResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
