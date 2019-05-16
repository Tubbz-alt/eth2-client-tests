package network

import (
	"log"

	"github.com/ethereum/eth2-client-tests/tester/genesis"
	"github.com/ethereum/eth2-client-tests/tester/network"
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
		err := network.Connect(node.Ip, port)
		if err != nil {
			log.Fatal("Error connecting to network: ", err)
		}
	}
}
