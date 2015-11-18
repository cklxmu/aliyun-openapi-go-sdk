package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetSnapshotRequest struct {
	RoaRequest
	ResourceName string
}

func (r *GetSnapshotRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *GetSnapshotRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *GetSnapshotRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/snapshots/ResourceName"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetSnapshotResponse struct {
}

func GetSnapshot(req *GetSnapshotRequest, accessId, accessSecret string) (*GetSnapshotResponse, error) {
	var pResponse GetSnapshotResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
