package pulsedive

import (
	"net/url"
	"strconv"
)

// FeedID implement api Feeds By Feed ID
func FeedID(id int) ([]byte, error) {
	q := url.Values{}
	q.Add("fid", strconv.Itoa(id))
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// FeedName implement api Feeds By Feed Name and Organization
func FeedName(name, organization string) ([]byte, error) {
	q := url.Values{}
	q.Add("feed", name)
	q.Add("organization", organization)
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}

// FeedLink implement api Feeds Linked Indicators
func FeedLink(id int) ([]byte, error) {
	q := url.Values{}
	q.Add("fid", strconv.Itoa(id))
	q.Add("get", "links")
	q.Add("pretty", pretty)
	q.Add("key", apiKey)
	return Get(q.Encode(), "/info.php")
}
