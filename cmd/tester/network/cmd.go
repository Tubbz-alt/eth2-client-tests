package network

import (
	"fmt"
	"log"

	"github.com/ethereum/eth2-client-tests/tester/report"

	"github.com/ethereum/eth2-client-tests/tester/docker"

	"github.com/ethereum/eth2-client-tests/tester/genesis"
	"github.com/urfave/cli"
)

var (
	NetworkCommand = cli.Command{
		Name:        "network",
		Usage:       "Tests network capacity",
		Description: `Tests that all the hosts in the testnet are available`,
		Action:      touchNetwork,
		Flags: []cli.Flag{
			TestnetName,
			Port,
			TestOutputFile,
		},
	}
)

func touchNetwork(ctx *cli.Context) {
	testNet := ctx.String(TestnetName.Name)
	port := ctx.Int(Port.Name)
	nodes := genesis.GetNodes(testNet)
	reports := []report.TestReport{}
	completeStdout := ""
	completeStderr := ""
	fail := false
	for _, node := range nodes {
		stdout, stderr, err := docker.Exec(fmt.Sprintf("whiteblock-node%d", node.LocalId), []string{"lsof", "-i", fmt.Sprintf(":%d", port)})
		//stdout, stderr, err  := docker.ExecScript(fmt.Sprintf("whiteblock-node%d", node.LocalId), "network.sh", testNetwork, "bash network.sh")
		completeStdout += stdout
		completeStderr += stderr
		fail = fail || err != nil
		reports = append(reports, report.TestReport{Name: fmt.Sprintf("whiteblock-node%d up", node.LocalId),
			Message: "",
			Failed:  err != nil,
			Logs:    fmt.Sprintf("%v", err)})
	}

	testReportOutput := ctx.String(TestOutputFile.Name)
	if testReportOutput != "" {
		report.WriteReport("network-up", reports, completeStdout, completeStderr, testReportOutput)
	} else {
		fmt.Printf("STDOUT: %s", completeStdout)
		fmt.Printf("STDERR: %s", completeStderr)
		if fail {
			log.Fatal("Error connecting to network ")
		}
	}
}
