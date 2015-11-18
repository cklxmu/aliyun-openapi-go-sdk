package core

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RoaRequest struct {
	AcsRequest
	Headers     http.Header
	PathPattern string
	PathParams  url.Values
}

func (r *RoaRequest) GetPath() (string, error) {
	var path string = r.PathPattern
	for param, _ := range r.PathParams {
		if r.PathParams.Get(param) == "" {
			return "", SdkError(fmt.Sprintf("Path Param %s for Roa Request not set.", param))
		}
		path = strings.Replace(path, param, r.PathParams.Get(param), -1)
	}
	if path == "" {
		path = "/"
	}
	return path, nil
}

func (r *RoaRequest) Init() {
	r.AcsRequest.Init()
	r.Headers = make(http.Header)
	r.PathParams = make(url.Values)
}

func (r *RoaRequest) GenSignature(accessId, accessSecret string) string {
	stringToSign := strings.ToUpper(r.GetMethod()) + "\n"
	if r.Headers.Get("Accept") != "" {
		stringToSign = stringToSign + r.Headers.Get("Accept")
	}
	stringToSign += "\n"
	if r.Headers.Get("Content-Md5") != "" {
		stringToSign = stringToSign + r.Headers.Get("Content-Md5")
	}
	stringToSign += "\n"
	if r.Headers.Get("Content-Type") != "" {
		stringToSign = stringToSign + r.Headers.Get("Content-Type")
	}
	stringToSign += "\n"
	stringToSign = stringToSign + time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT") + "\n" //TODO
	stringToSign = stringToSign + "x-acs-signature-method:HMAC-SHA1\n"
	stringToSign = stringToSign + "x-acs-signature-version:1.0\n"
	stringToSign = stringToSign + "x-acs-version:" + r.GetVersion() + "\n"
	path, _ := r.GetPath()
	stringToSign = stringToSign + path
	sep := "?"
	for queryParam, _ := range r.GetQueryParams() {
		stringToSign = stringToSign + sep + queryParam + "=" + r.GetQueryParams().Get(queryParam)
		sep = "&"
	}

	h := hmac.New(sha1.New, []byte(accessSecret))
	h.Write([]byte(stringToSign))
	signData := h.Sum(nil)
	fmt.Println("signData:", hex.EncodeToString(signData))
	signature := base64.StdEncoding.EncodeToString(signData)
	fmt.Println("signature:", signature)

	return signature
}

func (r *RoaRequest) DoAction(accessId, accessSecret string) (*http.Response, error) {
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

	if method == "" {
		method = METHOD_GET
	}
	if protocol == "" {
		protocol = PROTOCOL_HTTP
	}

	path, err := r.GetPath()
	if err != nil {
		return nil, err
	}
	fmt.Println("path:", path)

	signature := r.GenSignature(accessId, accessSecret)

	var req *http.Request
	requestUrl := protocol + "://" + domain + path
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
	req.Header.Set("Authorization", "acs "+accessId+":"+signature)
	req.Header.Set("x-acs-signature-method", "HMAC-SHA1")
	req.Header.Set("x-acs-signature-version", "1.0")
	req.Header.Set("x-acs-version", r.GetVersion())
	req.Header.Set("Date", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	return http.DefaultClient.Do(req)
}
