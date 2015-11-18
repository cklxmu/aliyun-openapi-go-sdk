package core

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type RpcRequest struct {
	AcsRequest
}

func (r *RpcRequest) GenSignature(accessId, accessSecret string) string {
	params := r.GetQueryParams()
	method := r.GetMethod()
	//step 1
	canonicalizedQueryString := EncodeParams(params)
	fmt.Println("canonicalizedQueryString:", canonicalizedQueryString)
	//step 2
	stringToSign := strings.ToUpper(method) + "&" + PercentEncode("/") + "&" + PercentEncode(canonicalizedQueryString)
	fmt.Println("stringToSign:", stringToSign)
	//step 3
	h := hmac.New(sha1.New, []byte(accessSecret+"&"))
	h.Write([]byte(stringToSign))
	signData := h.Sum(nil)
	fmt.Println("signData:", hex.EncodeToString(signData))
	signature := base64.StdEncoding.EncodeToString(signData)
	fmt.Println("signature:", signature)

	return signature
}

func (r *RpcRequest) DoAction(accessId, accessSecret string) (*http.Response, error) {
	var err error

	params := r.GetQueryParams()
	method := r.GetMethod()
	protocol := r.GetProtocol()
	domain := r.GetDomain()
	if domain == "" {
		domain, err = FindProductDomain(r.GetProduct(), r.GetRegionId())
		if err != nil {
			return nil, err
		}
	}

	//set common params
	if _, exist := params["SignatureNonce"]; !exist {
		params.Set("SignatureNonce", GenGuid())
	}
	if _, exist := params["TimeStamp"]; !exist {
		params.Set("TimeStamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	}
	if _, exist := params["SignatureMethod"]; !exist {
		params.Set("SignatureMethod", "HMAC-SHA1")
	}
	if _, exist := params["SignatureVersion"]; !exist {
		params.Set("SignatureVersion", "1.0")
	}
	if _, exist := params["Format"]; !exist {
		params.Set("Format", FORMAT_XML)
	}
	if _, exist := params["AccessKeyId"]; !exist {
		params.Set("AccessKeyId", accessId)
	}

	if method == "" {
		method = METHOD_GET
	}
	if protocol == "" {
		protocol = PROTOCOL_HTTP
	}

	//generate and set signature
	signature := r.GenSignature(accessId, accessSecret)
	params.Set("Signature", signature)

	//send http request
	var req *http.Request
	requestUrl := protocol + "://" + domain + "/"
	if METHOD_POST == strings.ToUpper(method) {
		req, err = http.NewRequest(method, requestUrl, strings.NewReader(EncodeParams(params)))
		if err != nil {
			return nil, err
		}
	} else {
		if len(params) > 0 {
			requestUrl += "?" + EncodeParams(params)
		}
		req, err = http.NewRequest(method, requestUrl, nil)
		if err != nil {
			return nil, err
		}
	}
	return http.DefaultClient.Do(req)
}
