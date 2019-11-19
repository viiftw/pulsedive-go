package pulsedive

import (
	"net/url"
	"strconv"
)

// IndicatorByID implement api Indicators By Indicator ID
func IndicatorByID(id int) ([]byte, error) {
	q := url.Values{}
	q.Add("schema", schema)
	q.Add("iid", strconv.Itoa(id))
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// IndicatorByValue implement api Indicators By Value
func IndicatorByValue(indicatorValue string) ([]byte, error) {
	q := url.Values{}
	q.Add("indicator", indicatorValue)
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// IndicatorLinks implement api Indicators Links
func IndicatorLinks(id int) ([]byte, error) {
	q := url.Values{}
	q.Add("iid", strconv.Itoa(id))
	q.Add("get", "links")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// IndicatorProperties implement api Indicators Properties
func IndicatorProperties(id int) ([]byte, error) {
	q := url.Values{}
	q.Add("iid", strconv.Itoa(id))
	q.Add("get", "properties")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}
