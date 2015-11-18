package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type ModifyDrdsDBPasswdRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	NewPasswd      string
}

func (r *ModifyDrdsDBPasswdRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *ModifyDrdsDBPasswdRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *ModifyDrdsDBPasswdRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *ModifyDrdsDBPasswdRequest) GetDbName() string {
	return r.DbName
}
func (r *ModifyDrdsDBPasswdRequest) SetNewPasswd(value string) {
	r.NewPasswd = value
	r.QueryParams.Set("NewPasswd", value)
}
func (r *ModifyDrdsDBPasswdRequest) GetNewPasswd() string {
	return r.NewPasswd
}

func (r *ModifyDrdsDBPasswdRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDrdsDBPasswd")
	r.SetProduct(Product)
}

type ModifyDrdsDBPasswdResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func ModifyDrdsDBPasswd(req *ModifyDrdsDBPasswdRequest, accessId, accessSecret string) (*ModifyDrdsDBPasswdResponse, error) {
	var pResponse ModifyDrdsDBPasswdResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
