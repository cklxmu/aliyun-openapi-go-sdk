package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateTransactionRequest struct {
	RpcRequest
	ScriptId        int
	TransactionName string
}

func (r *CreateTransactionRequest) SetScriptId(value int) {
	r.ScriptId = value
	r.QueryParams.Set("ScriptId", strconv.Itoa(value))
}
func (r *CreateTransactionRequest) GetScriptId() int {
	return r.ScriptId
}
func (r *CreateTransactionRequest) SetTransactionName(value string) {
	r.TransactionName = value
	r.QueryParams.Set("TransactionName", value)
}
func (r *CreateTransactionRequest) GetTransactionName() string {
	return r.TransactionName
}

func (r *CreateTransactionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateTransaction")
	r.SetProduct(Product)
}

type CreateTransactionResponse struct {
	TransactionId int `xml:"TransactionId" json:"TransactionId"`
}

func CreateTransaction(req *CreateTransactionRequest, accessId, accessSecret string) (*CreateTransactionResponse, error) {
	var pResponse CreateTransactionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
