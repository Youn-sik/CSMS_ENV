package envc

import (
	envType "CSMS_ENV/env/type"
	"CSMS_ENV/logger"
	"bytes"
	"crypto/tls"
	"database/sql"
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tidbClientEnv *sql.DB
var envHttpClient *http.Client
var bid string
var bkey string

func init() {
	tidbClientEnv = newTidbConnectionEnv()
	envHttpClient = getHttpClient()
	bid = "XX"
	bkey = "XXXXXXXXXXXXXXXX"
}

func newTidbConnectionEnv() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("TIDB_USER")+":"+os.Getenv("TIDB_PASS")+"@tcp("+os.Getenv("TIDB_HOST")+":"+os.Getenv("TIDB_PORT")+")/"+os.Getenv("TIDB_DBNM"))
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return nil
	}
	return db
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func envHttpRequest[TQ envType.EnvHttpRequestInterface, TR envType.EnvHttpResponseInterface](reqUrl string, req TQ) (TR, error) {
	var res TR

	// reqBody, _ := json.Marshal(req)
	// reqBodyBuff := bytes.NewBuffer(reqBody)

	// request, err := http.NewRequest("POST", "http://10.101.160.34/r2"+reqUrl, reqBodyBuff)
	// if err != nil {
	// 	return res, err
	// }

	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// response, err := envHttpClient.Do(request)
	// if err != nil {
	// 	return res, err
	// }

	reqBody, _ := json.Marshal(req)

	formData := url.Values{}
	formData.Set("messages", string(reqBody))

	log.Println("http://10.101.160.34/r2" + reqUrl)
	log.Println(formData)

	response, err := http.Post(
		"http://10.101.160.34/r2"+reqUrl,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(formData.Encode()),
	)
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return res, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
