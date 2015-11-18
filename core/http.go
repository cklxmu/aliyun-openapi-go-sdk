package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

func ApiHttpRequest(accessId, accessSecret string, req ApiRequest) ([]byte, error) {
	response, err := req.DoAction(accessId, accessSecret)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.StatusCode >= 400 {
		apiErr, err := ErrorUnmarshal(req.GetQueryParams().Get("Format"), body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if apiErr != nil {
			fmt.Println("body:", string(body))
			fmt.Println("RequestID:", apiErr.RequestId)
			fmt.Println("HostID:", apiErr.HostId)
			fmt.Println("Code:", apiErr.Code)
			fmt.Println("Message:", apiErr.Message)
			return nil, apiErr
		}
	}
	return body, nil
}

func ApiUnmarshalResponse(format string, body []byte, pResponse interface{}) error {
	fmt.Println("UnmarshalResponse: body:", string(body))
	var err error
	if strings.ToUpper(format) == FORMAT_XML {
		err = xml.Unmarshal(body, pResponse)
	} else {
		err = json.Unmarshal(body, pResponse)
	}
	if err != nil {
		return err
	}
	return nil

}
