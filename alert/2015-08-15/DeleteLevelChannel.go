package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteLevelChannelRequest struct {
	RoaRequest
	ProjectName string
	Level       string
}

func (r *DeleteLevelChannelRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteLevelChannelRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteLevelChannelRequest) SetLevel(value string) {
	r.Level = value
	r.PathParams.Set("Level", value)
}
func (r *DeleteLevelChannelRequest) GetLevel() string {
	return r.Level
}

func (r *DeleteLevelChannelRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/level_channels/Level"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteLevelChannelResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteLevelChannel(req *DeleteLevelChannelRequest, accessId, accessSecret string) (*DeleteLevelChannelResponse, error) {
	var pResponse DeleteLevelChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
