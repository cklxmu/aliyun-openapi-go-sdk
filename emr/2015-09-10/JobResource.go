package emr

import (
	. "aliyun-openapi-go-sdk/core"
)

type JobResourceRequest struct {
	RpcRequest
	Bucket string
	Path   string
}

func (r *JobResourceRequest) SetBucket(value string) {
	r.Bucket = value
	r.QueryParams.Set("Bucket", value)
}
func (r *JobResourceRequest) GetBucket() string {
	return r.Bucket
}
func (r *JobResourceRequest) SetPath(value string) {
	r.Path = value
	r.QueryParams.Set("Path", value)
}
func (r *JobResourceRequest) GetPath() string {
	return r.Path
}

func (r *JobResourceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("JobResource")
	r.SetProduct(Product)
}

type JobResourceResponse struct {
	JobResourceInfoList struct {
		JobResourceInfo []struct {
			Type     int    `xml:"Type" json:"Type"`
			Filename string `xml:"Filename" json:"Filename"`
			Size     int    `xml:"Size" json:"Size"`
			Region   string `xml:"Region" json:"Region"`
		} `xml:"JobResourceInfo" json:"JobResourceInfo"`
	} `xml:"JobResourceInfoList" json:"JobResourceInfoList"`
}

func JobResource(req *JobResourceRequest, accessId, accessSecret string) (*JobResourceResponse, error) {
	var pResponse JobResourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
