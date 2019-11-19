package pulsedive

import (
	"net/url"
)

// SearchIndicators implement api Search Indicators
func SearchIndicators(value string, attributes []string, threats []string, feeds []string) ([]byte, error) {
	q := url.Values{}
	q.Add("value", value)
	for _, indicatorType := range indicatorTypes {
		q.Add("type[]", indicatorType)
	}
	for _, risk := range risks {
		q.Add("risk[]", risk)
	}
	for _, attribute := range attributes {
		q.Add("attribute[]", attribute)
	}
	q.Add("property", "content-type:text/html")
	for _, threat := range threats {
		q.Add("threat[]", threat)
	}
	for _, feed := range feeds {
		q.Add("feed[]", feed)
	}
	q.Add("limit", "hundred")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/search.php")
}

// SearchToCSV implement api Search Exporting to CSV
func SearchToCSV(value string, attributes []string, threats []string, feeds []string) ([]byte, error) {
	q := url.Values{}
	q.Add("value", value)
	q.Add("property", "content-type:text/html")
	for _, indicatorType := range indicatorTypes {
		q.Add("type[]", indicatorType)
	}
	for _, risk := range risks {
		q.Add("risk[]", risk)
	}
	for _, attribute := range attributes {
		q.Add("attribute[]", attribute)
	}
	q.Add("property", "content-type:text/html")
	for _, threat := range threats {
		q.Add("threat[]", threat)
	}
	for _, feed := range feeds {
		q.Add("feed[]", feed)
	}
	q.Add("limit", "hundred")
	q.Add("export", "1")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/search.php")
}

// SearchThreat implement api Search Threats
func SearchThreat(value string, attributes []string, threats []string, feeds []string) ([]byte, error) {
	q := url.Values{}
	q.Add("search", "threat")
	q.Add("value", value)
	for _, category := range categories {
		q.Add("category[]", category)
	}
	for _, risk := range risks {
		q.Add("risk[]", risk)
	}
	for _, attribute := range attributes {
		q.Add("attribute[]", attribute)
	}
	q.Add("property", "content-type:text/html")
	for _, threat := range threats {
		q.Add("threat[]", threat)
	}
	for _, feed := range feeds {
		q.Add("feed[]", feed)
	}
	q.Add("property", "content-type:text/html")
	q.Add("limit", "hundred")
	q.Add("splitrisk", "1")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/search.php")
}

// SearchFeed implement api Search Feeds
func SearchFeed(value string) ([]byte, error) {
	q := url.Values{}
	q.Add("search", "feed")
	q.Add("value", value)
	for _, category := range categories {
		q.Add("category[]", category)
	}
	q.Add("splitrisk", "1")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/search.php")
}
