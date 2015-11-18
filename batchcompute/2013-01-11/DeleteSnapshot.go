package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteSnapshotRequest struct {
	RoaRequest
	ResourceName string
}

func (r *DeleteSnapshotRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *DeleteSnapshotRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *DeleteSnapshotRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/snapshots/ResourceName"
	r.SetMethod("DELETE")
	r.SetProduct(Product)
}

type DeleteSnapshotResponse struct {
}

func DeleteSnapshot(req *DeleteSnapshotRequest, accessId, accessSecret string) (*DeleteSnapshotResponse, error) {
	var pResponse DeleteSnapshotResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
