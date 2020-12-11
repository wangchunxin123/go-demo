/* #
# Author: wangchunxin
# Created Time: 2020/11/21 3:33 下午
# File Name:
# Description:
# */
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const (
	MOTHOD_GET  = "GET"
	MOTHOD_POST = "POST"
)

type RequestParams struct {
	mu  sync.Mutex
	Map map[string]interface{}
}

func ApiRequest(mothod string, path string, paramsByte []byte, headerParams map[string]string, bean interface{}) error {
	var req *http.Request
	var paramsMap RequestParams
	paramsMap.mu.Lock()
	defer paramsMap.mu.Unlock()
	json.Unmarshal(paramsByte, &paramsMap.Map)
	if MOTHOD_GET == mothod {
		req, _ = http.NewRequest(MOTHOD_GET, path, nil)
		q := req.URL.Query()
		for key, value := range paramsMap.Map {
			var valueStr string
			switch value.(type) {
			case float64:
				valueStr = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			case string :
				valueStr = value.(string)
			default:
				vbyte, _ := json.Marshal(value)
				valueStr = string(vbyte)
			}
			q.Add(key, valueStr)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())
	} else if MOTHOD_POST == mothod {
		body := &bytes.Buffer{}
		body.WriteString(string(paramsByte))
		req, _ = http.NewRequest(MOTHOD_POST, path, body)
	} else {
	}

	//头信息
	if len(headerParams) > 0 {
		for hk, hv := range headerParams {
			req.Header.Add(hk, hv)
		}
	}

	client := http.Client{}
	response, err := client.Do(req)
	responseBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseBody, &bean)
	response.Body.Close()
	return err
}
