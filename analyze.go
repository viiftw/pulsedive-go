package pulsedive

import (
	"crypto/tls"
	"encoding/base64"
	// "encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Analyze adding to the Queue for Analyze
func Analyze(ioc string) ([]byte, error) {
	iocEnc := base64.StdEncoding.EncodeToString([]byte(ioc))
	data := url.Values{}
	data.Set("ioc", iocEnc)
	data.Set("enrich", "1")
	data.Set("probe", "1")
	data.Set("pretty", "1")
	data.Set("key", apiKey)

	req, err := http.NewRequest("POST", "https://pulsedive.com/api/analyze.php", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	trt := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: trt}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// var results PDAResponse
	// json.Unmarshal(resBody, &results)
	return resBody, nil
}

// AnalyzeResult get result from Analyze Retrieving the Results api
func AnalyzeResult(qid int) ([]byte, error) {
	q := url.Values{}
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	q.Add("qid", strconv.Itoa(qid))
	return Get(q.Encode(), "/analyze.php")
}
