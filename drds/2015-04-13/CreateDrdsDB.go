package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateDrdsDBRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	Encode         string
	Password       string
	RdsInstances   string
}

func (r *CreateDrdsDBRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *CreateDrdsDBRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *CreateDrdsDBRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *CreateDrdsDBRequest) GetDbName() string {
	return r.DbName
}
func (r *CreateDrdsDBRequest) SetEncode(value string) {
	r.Encode = value
	r.QueryParams.Set("Encode", value)
}
func (r *CreateDrdsDBRequest) GetEncode() string {
	return r.Encode
}
func (r *CreateDrdsDBRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *CreateDrdsDBRequest) GetPassword() string {
	return r.Password
}
func (r *CreateDrdsDBRequest) SetRdsInstances(value string) {
	r.RdsInstances = value
	r.QueryParams.Set("RdsInstances", value)
}
func (r *CreateDrdsDBRequest) GetRdsInstances() string {
	return r.RdsInstances
}

func (r *CreateDrdsDBRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDrdsDB")
	r.SetProduct(Product)
}

type CreateDrdsDBResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func CreateDrdsDB(req *CreateDrdsDBRequest, accessId, accessSecret string) (*CreateDrdsDBResponse, error) {
	var pResponse CreateDrdsDBResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
