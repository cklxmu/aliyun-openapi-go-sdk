package ubsms_inner

import (
	. "aliyun-openapi-go-sdk/core"
)

type NotifyUserBusinessCommandRequest struct {
	RpcRequest
	Uid         string
	ServiceCode string
	Cmd         string
	Region      string
	InstanceId  string
	ClientToken string
	Password    string
}

func (r *NotifyUserBusinessCommandRequest) SetUid(value string) {
	r.Uid = value
	r.QueryParams.Set("Uid", value)
}
func (r *NotifyUserBusinessCommandRequest) GetUid() string {
	return r.Uid
}
func (r *NotifyUserBusinessCommandRequest) SetServiceCode(value string) {
	r.ServiceCode = value
	r.QueryParams.Set("ServiceCode", value)
}
func (r *NotifyUserBusinessCommandRequest) GetServiceCode() string {
	return r.ServiceCode
}
func (r *NotifyUserBusinessCommandRequest) SetCmd(value string) {
	r.Cmd = value
	r.QueryParams.Set("Cmd", value)
}
func (r *NotifyUserBusinessCommandRequest) GetCmd() string {
	return r.Cmd
}
func (r *NotifyUserBusinessCommandRequest) SetRegion(value string) {
	r.Region = value
	r.QueryParams.Set("Region", value)
}
func (r *NotifyUserBusinessCommandRequest) GetRegion() string {
	return r.Region
}
func (r *NotifyUserBusinessCommandRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *NotifyUserBusinessCommandRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *NotifyUserBusinessCommandRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *NotifyUserBusinessCommandRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *NotifyUserBusinessCommandRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *NotifyUserBusinessCommandRequest) GetPassword() string {
	return r.Password
}

func (r *NotifyUserBusinessCommandRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("NotifyUserBusinessCommand")
	r.SetProduct(Product)
}

type NotifyUserBusinessCommandResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func NotifyUserBusinessCommand(req *NotifyUserBusinessCommandRequest, accessId, accessSecret string) (*NotifyUserBusinessCommandResponse, error) {
	var pResponse NotifyUserBusinessCommandResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
