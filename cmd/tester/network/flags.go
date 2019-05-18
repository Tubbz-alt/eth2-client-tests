package network

import "github.com/urfave/cli"

var (
	TestnetName = cli.StringFlag{
		Name:  "testnet",
		Usage: "ID of the testnet to target",
	}
	Port = cli.IntFlag{
		Name:  "port",
		Usage: "Port served by all nodes in the testnet",
		Value: 9000,
	}
)
