package envc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var envHttpClient *http.Client
var bid string
var bkey string

func init() {
	envHttpClient = GetHttpClient()
	bid = "XX"
	bkey = "XXXXXXXXXXXXXXXX"
}

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func EnvHttpRequest[TQ EnvHttpRequestInterface, TR EnvHttpResponseInterface](url string, req TQ) (TR, error) {
	var res TR

	reqBody, _ := json.Marshal(req)
	reqBodyBuff := bytes.NewBuffer(reqBody)
	request, err := http.NewRequest("POST", "http://10.101.160.34/r2"+url, reqBodyBuff)
	if err != nil {
		return res, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := envHttpClient.Do(request)
	if err != nil {
		return res, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
