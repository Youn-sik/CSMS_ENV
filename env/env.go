package envc

import (
	envType "CSMS_ENV/env/type"
	"CSMS_ENV/logger"
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var tidbClientEnv *sql.DB
var envHttpClient *http.Client
var bid string
var bkey string

func init() {
	tidbClientEnv = NewTidbConnectionEnv()
	envHttpClient = GetHttpClient()
	bid = "XX"
	bkey = "XXXXXXXXXXXXXXXX"
}

func NewTidbConnectionEnv() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("TIDB_USER")+":"+os.Getenv("TIDB_PASS")+"@tcp("+os.Getenv("TIDB_HOST")+":"+os.Getenv("TIDB_PORT")+")/"+os.Getenv("TIDB_DBNM"))
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return nil
	}
	return db
}

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func EnvHttpRequest[TQ envType.EnvHttpRequestInterface, TR envType.EnvHttpResponseInterface](url string, req TQ) (TR, error) {
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
