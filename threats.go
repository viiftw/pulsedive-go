package pulsedive
import (
	"net/url"
	"strconv"
)

// ThreatByID implement api Threats By Threat ID
func ThreatByID(id int) ([]byte, error){
	q := url.Values{}
	q.Add("tid", strconv.Itoa(id))
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// ThreatByName implement api Threats By Threat Name
func ThreatByName(name string) ([]byte, error){
	q := url.Values{}
	q.Add("threat", name)
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// ThreatSummary implement api Threats By Threats Indicator Summary
func ThreatSummary(id int) ([]byte, error){
	q := url.Values{}
	q.Add("tid", strconv.Itoa(id))
	q.Add("get", "links")
	q.Add("summary", "1")
	q.Add("splitrisk", "1")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// ThreatLink implement api Threats By Threats Linked Indicators
func ThreatLink(id int) ([]byte, error){
	q := url.Values{}
	q.Add("tid", strconv.Itoa(id))
	q.Add("get", "links")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}