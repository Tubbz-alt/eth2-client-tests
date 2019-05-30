package prometheus

import "github.com/urfave/cli"

var (
	TestnetName = cli.StringFlag{
		Name:  "testnet",
		Usage: "ID of the testnet to target",
	}
	Query = cli.StringFlag{
		Name:  "query",
		Usage: "Prometheus query",
	}
	Pretty = cli.BoolFlag{
		Name:  "pretty",
		Usage: "Pretty print",
	}
)
