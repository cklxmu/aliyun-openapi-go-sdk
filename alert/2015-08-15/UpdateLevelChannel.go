package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateLevelChannelRequest struct {
	RoaRequest
	ProjectName         string
	Level               string
	LevelChannelSetting string
}

func (r *UpdateLevelChannelRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateLevelChannelRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateLevelChannelRequest) SetLevel(value string) {
	r.Level = value
	r.PathParams.Set("Level", value)
}
func (r *UpdateLevelChannelRequest) GetLevel() string {
	return r.Level
}
func (r *UpdateLevelChannelRequest) SetLevelChannelSetting(value string) {
	r.LevelChannelSetting = value
	r.QueryParams.Set("LevelChannelSetting", value)
}
func (r *UpdateLevelChannelRequest) GetLevelChannelSetting() string {
	return r.LevelChannelSetting
}

func (r *UpdateLevelChannelRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/level_channels/Level"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateLevelChannelResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateLevelChannel(req *UpdateLevelChannelRequest, accessId, accessSecret string) (*UpdateLevelChannelResponse, error) {
	var pResponse UpdateLevelChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
