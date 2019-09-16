# pulsedive-go
A Golang library for [Pulsedive](https://pulsedive.com/) API

## Usage

```go
package main
import (
    "log"
    "time"
    "github.com/viiftw/pulsedive-go"
)

func main() {
    // Load your Pulsedive API key <YOUR-API-KEY>: string, Pretty: true
    pulsedive.Pulsedive(<YOUR-API-KEY>, true)
    // Indicators
    body, err := pulsedive.IndicatorByValue("pulsedive.com")
    body, err = pulsedive.IndicatorByID(2)
    body, err = pulsedive.IndicatorLinks(2)
    body, err = pulsedive.IndicatorProperties(2)
    // Threats
    body, err = pulsedive.ThreatByID(1)
    body, err = pulsedive.ThreatByName("Zeus")
    body, err = pulsedive.ThreatLink(1)
    body, err = pulsedive.ThreatSummary(1)
    // Feeds
    body, err = pulsedive.FeedID(1)
    body, err = pulsedive.FeedLink(1)
    body, err = pulsedive.FeedName("Zeus Bad Domains", "abuse.ch")
    // Analyze
    qid, err := pulsedive.Analyze("google.com")
    time.Sleep(5 * time.Second) // Wait for analyze
    body, err = pulsedive.AnalyzeResult(qid)
    // Search
    attributes := []string{"http", "443"}
    threats := []string{"zeus", "banjori"}
    feeds := []string{"zeus bad domains", "zeus bad ips"}
    body, err = pulsedive.SearchIndicators("", attributes, threats, feeds)
    body, err = pulsedive.SearchThreat("zeus", attributes, threats, feeds)
    body, err = pulsedive.SearchToCSV("", attributes, threats, feeds)
    body, err = pulsedive.SearchFeed("zeus")
    if err != nil {
      log.Println(err)
    }
    log.Println(string(body))
    // => {"results":[{"fid":1,"name":"Zeus Bad Domains","category":"malware"...
}
```

## License

The pulsedive-go is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).