package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListSnapshotsRequest struct {
	RoaRequest
}

func (r *ListSnapshotsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/snapshots"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type ListSnapshotsResponse struct {
}

func ListSnapshots(req *ListSnapshotsRequest, accessId, accessSecret string) (*ListSnapshotsResponse, error) {
	var pResponse ListSnapshotsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
