package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

type SdkError string

func (e SdkError) Error() string {
	return fmt.Sprintf("SDK Error: %s", e)
}

type ApiError struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	HostId    string `json:"HostId" xml:"HostId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API Error: %s: %s", e.Code, e.Message)
}

func ErrorUnmarshal(format string, data []byte) (*ApiError, error) {
	var apiErr ApiError
	var err error
	if strings.ToUpper(format) == FORMAT_XML {
		err = xml.Unmarshal(data, &apiErr)
	} else {
		err = json.Unmarshal(data, &apiErr)
	}
	if err != nil {
		return nil, err
	}

	return &apiErr, nil
}
