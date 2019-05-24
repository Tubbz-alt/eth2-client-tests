package network

import (
	"fmt"
	"log"

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
		},
	}
)

func touchNetwork(ctx *cli.Context) {
	testNet := ctx.String(TestnetName.Name)
	port := ctx.Int(Port.Name)
	nodes := genesis.GetNodes(testNet)
	for _, node := range nodes {
		stdout, stderr, err := docker.Exec(fmt.Sprintf("whiteblock-node%d", node.LocalId), []string{"lsof", "-i", fmt.Sprintf(":%d", port)})
		//stdout, stderr, err  := docker.ExecScript(fmt.Sprintf("whiteblock-node%d", node.LocalId), "network.sh", testNetwork, "bash network.sh")
		fmt.Printf("STDOUT: %s", stdout)
		fmt.Printf("STDERR: %s", stderr)
		if err != nil {
			log.Fatal("Error connecting to network: ", err)
		}
	}
}
