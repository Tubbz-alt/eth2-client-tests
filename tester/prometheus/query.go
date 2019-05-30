package prometheus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const prometheusServer = "http://eth2jenkins.themachine.io:9090/api/v1/query?"

/**
{
        "metric": {
          "__name__": "p2p_peer_count",
          "blockchain": "prysm",
          "instance": "10.1.0.18:8088",
          "ip": "10.1.0.18",
          "job": "prysm-6adfb3b2-1f3b-475a-be72-b16ce5987232-10.1.0.18",
          "testnet": "6750b16d-a21e-4294-aeb6-523b77af7b11"
        },
        "value": [
          1559196629.3,
          "3"
        ]
      }
*/

type ResultItem struct {
	Metric map[string]string
	Value  []interface{}
}

type ResponseData struct {
	ResultType string
	Result     []ResultItem
}

type PrometheusQueryResponse struct {
	Status string
	Data   ResponseData
}

func Query(expr string) PrometheusQueryResponse {
	params := url.Values{}
	params.Add("query", expr)

	resp, err := http.Get(prometheusServer + params.Encode())
	if err != nil {
		log.Fatalf("Error communicating with Prometheus %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body %v", err)
	}

	if resp.StatusCode == 400 {
		log.Fatalf("Invalid query %s: %s", expr, body)
	}
	if resp.StatusCode == 422 {
		log.Fatalf("Cannot execute query %s: %s", expr, body)
	}
	if resp.StatusCode == 503 {
		log.Fatalf("Query timed out %s: %s", expr, body)
	}
	if resp.StatusCode == 200 {
		var response PrometheusQueryResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Error unmarshaling body %v", err)
		}
		if response.Status != "success" {
			log.Fatalf("Prometheus response status %s", response.Status)
		}

		return response
	}
	log.Fatalf("Response with error status code: %d", resp.StatusCode)
	return PrometheusQueryResponse{}
}
