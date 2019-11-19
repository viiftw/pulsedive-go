package pulsedive
import (
	"io/ioutil"
	"net/http"
	"crypto/tls"
	"time"
	"strconv"
)

var (
	apiKey = ""
	pretty = "1"
	schema = "1"
	pulsediveURL = "https://pulsedive.com/api"
	risks = []string{"unknown", "none", "low", "medium", "high", "critical", "retired"}
	indicatorTypes = []string{"ip", "ipv6", "url", "domain", "artifact"}
	categories = []string{"general", "abuse", "apt", "attack", "botnet", "crime",
              "exploitkit", "fraud", "group", "malware", "proxy", "pup",
							"reconnaissance", "spam", "terrorism", "phishing", "vulnerability"}
)

// PDAResponse define for pulsedive api queue response
type PDAResponse struct {
	Success string `json:"success"`
	Qid     int    `json:"qid"`
}

// Get create a get request
func Get(query, path string) ([]byte, error) {
	req, err := http.NewRequest("GET", pulsediveURL + path, nil)
	if err != nil {
			return nil, err
	}
	req.URL.RawQuery = query
	trt := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

// Pulsedive set API KEY and pretty
func Pulsedive(key string, isPretty bool) {
	apiKey = key
	if !isPretty {
		pretty = strconv.Itoa(0)
	}
}