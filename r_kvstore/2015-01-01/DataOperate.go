package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DataOperateRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Command              string
}

func (r *DataOperateRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DataOperateRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DataOperateRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DataOperateRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DataOperateRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DataOperateRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DataOperateRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DataOperateRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DataOperateRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DataOperateRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DataOperateRequest) SetCommand(value string) {
	r.Command = value
	r.QueryParams.Set("Command", value)
}
func (r *DataOperateRequest) GetCommand() string {
	return r.Command
}

func (r *DataOperateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DataOperate")
	r.SetProduct(Product)
}

type DataOperateResponse struct {
	CommandResult string `xml:"CommandResult" json:"CommandResult"`
}

func DataOperate(req *DataOperateRequest, accessId, accessSecret string) (*DataOperateResponse, error) {
	var pResponse DataOperateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
