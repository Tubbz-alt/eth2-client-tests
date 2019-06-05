package deposit

import "github.com/urfave/cli"

var (
	HttpPath = cli.StringFlag{
		Name:  "http-path",
		Usage: "path to ETH1 client HTTP port",
		Value: "https://goerli.prylabs.net",
	}
	PrivKeyString = cli.StringFlag{
		Name:  "priv-key",
		Usage: "Hex string of the private key of a Goerli ETH account",
	}
	DepositsForChainStart = cli.IntFlag{
		Name:  "chain-start",
		Usage: "Number of validators required for chain start",
		Value: 8,
	}
	MinDepositAmount = cli.Int64Flag{
		Name:  "min-deposit",
		Usage: "Minimum deposit value allowed in contract",
		Value: 200,
	}
	MaxDepositAmount = cli.Int64Flag{
		Name:  "max-deposit",
		Usage: "Maximum deposit value allowed in contract",
		Value: 3.2 * 1e9,
	}
	OutputFile = cli.StringFlag{
		Name:  "output-file",
		Usage: "File to store contract address",
	}
	ContractFlag = cli.StringFlag{
		Name:  "contract",
		Usage: "Address of the deposit contract",
	}
	Password = cli.StringFlag{
		Name:  "password",
		Usage: "Password of the validator",
	}
	KeystoreFlag = cli.StringFlag{
		Name:  "keystore",
		Usage: "Path to keystore",
	}
	AmountFlag = cli.Int64Flag{
		Name:  "amount",
		Usage: "Deposit amount",
	}
)
