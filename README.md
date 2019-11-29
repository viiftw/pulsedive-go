# pulsedive-go
A Golang library for [Pulsedive](https://pulsedive.com/) API

## Usage
### Indicator By Value
```go
package main
import (
    "log"
    "github.com/viiftw/pulsedive-go"
)

func main() {
    // Load your Pulsedive API key <YOUR-API-KEY>: string, Pretty: true
    pulsedive.Pulsedive(<YOUR-API-KEY>, true)

    // Indicators
    body, err := pulsedive.IndicatorByValue("pulsedive.com")
    if err != nil {
      log.Println(err)
    }
    log.Println(string(body))
    // => {"iid"=>53929,"type"=>"domain","indicator"=>"pulsedive.com","risk"=>"medium", ...
}
```

## Support API
```go
  // Indicators
  pulsedive.IndicatorByID(id int) ([]byte, error)
  pulsedive.IndicatorByValue(indicatorValue string) ([]byte, error)
  pulsedive.IndicatorLinks(id int) ([]byte, error)
  pulsedive.IndicatorProperties(id int) ([]byte, error)
  // Threats
  pulsedive.ThreatByID(id int) ([]byte, error)
  pulsedive.ThreatByName(name string) ([]byte, error)
  pulsedive.ThreatSummary(id int) ([]byte, error)
  pulsedive.ThreatLink(id int) ([]byte, error)
  // Feeds
  pulsedive.FeedID(id int) ([]byte, error)
  pulsedive.FeedName(name, organization string) ([]byte, error)
  pulsedive.FeedLink(id int) ([]byte, error)
  // Analyze
  pulsedive.Analyze(ioc string) ([]byte, error)
  pulsedive.AnalyzeResult(qid int) ([]byte, error)
  // Search
  pulsedive.SearchIndicators(value string, attributes []string, threats []string, feeds []string) ([]byte, error)
  pulsedive.SearchToCSV(value string, attributes []string, threats []string, feeds []string) ([]byte, error)
  pulsedive.SearchThreat(value string, attributes []string, threats []string, feeds []string) ([]byte, error)
  pulsedive.SearchFeed(value string) ([]byte, error)
```

## License

The pulsedive-go is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).