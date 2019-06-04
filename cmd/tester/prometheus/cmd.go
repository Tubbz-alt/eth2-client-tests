package prometheus

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/eth2-client-tests/tester/genesis"
	"github.com/ethereum/eth2-client-tests/tester/report"

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
					TestOutputFile,
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

	var message string
	if counter != len(nodes) {
		message = fmt.Sprintf("Not all nodes are reporting being up. %d/%d", counter, len(nodes))
	} else {
		message = fmt.Sprintf("All nodes reporting")
	}
	log.Println(message)

	testReportOutput := ctx.String(TestOutputFile.Name)
	if testReportOutput != "" {
		//testReports []TestReport, stdout string, stderr string, outputFilePath string)
		report.WriteReport("prometheus-up", []report.TestReport{
			report.TestReport{Name: "prometheus-up",
				Message: message,
				Failed:  counter != len(nodes),
				Logs:    ""},
		}, "", "", testReportOutput)
	}
}
