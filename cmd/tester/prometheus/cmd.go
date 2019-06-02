package prometheus

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/eth2-client-tests/tester/genesis"
	"log"

	"github.com/ethereum/eth2-client-tests/tester/prometheus"
	"github.com/urfave/cli"
)

var (
	PrometheusCommand = cli.Command{
		Name:        "prometheus",
		Usage:       "Commands to query Prometheus",
		Description: `Commands to query and interpret Prometheus data`,
		Subcommands: []cli.Command{
			cli.Command{
				Name:        "query",
				Usage:       "Queries Prometheus",
				Description: `Queries Prometheus data`,
				Action:      queryPrometheus,
				Flags: []cli.Flag{
					TestnetName,
					Query,
					Pretty,
				},
			},
			cli.Command{
				Name:        "up",
				Usage:       "check nodes in the testnet",
				Description: "Queries Prometheus to check nodes in the testnet are reporting",
				Action:      checkPrometheusUp,
				Flags: []cli.Flag{
					TestnetName,
				},
			},
		},
	}
)

func queryPrometheus(ctx *cli.Context) {
	query := ctx.String(Query.Name)
	testnet := ctx.String(TestnetName.Name)
	if testnet != "" {
		query += fmt.Sprintf(`%s AND {testnet="%s"}`, query, testnet)
	}

	resp := prometheus.Query(query)
	var bytes []byte
	if ctx.Bool(Pretty.Name) {
		bytes, _ = json.MarshalIndent(resp, "", "  ")
	} else {
		bytes, _ = json.Marshal(resp)
	}
	fmt.Println(string(bytes))
}

func checkPrometheusUp(ctx *cli.Context) {
	testnet := ctx.String(TestnetName.Name)
	resp := prometheus.Query(fmt.Sprintf(`{testnet="%s"}`, testnet))
	counter := 0
	for _, item := range resp.Data.Result {
		val := item.Value[1]
		if val == 1 {
			counter++
		}
	}
	nodes := genesis.GetNodes(testnet)
	if counter != len(nodes) {
		log.Fatalf("Not all nodes are reporting being up. %d/%d", counter, len(nodes))
	} else {
		log.Println("All nodes reporting")
	}
}