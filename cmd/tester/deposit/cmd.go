package deposit

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/eth2-client-tests/tester/deposit"
	"github.com/urfave/cli"
)

var (
	DepositContractCommand = cli.Command{
		Name:        "contract",
		Usage:       "Deploys deposit contract",
		Description: `Deploys deposit contract to the goerli testnet`,
		Action:      deployDepositContract,
		Flags: []cli.Flag{
			HttpPath,
			PrivKeyString,
			DepositsForChainStart,
			MinDepositAmount,
			MaxDepositAmount,
			OutputFile,
		},
	}

	SendDepositTxCommand = cli.Command{
		Name:        "sendTx",
		Usage:       "Send transaction to deposit contract",
		Description: `Sends tx to deposit contract`,
		Action:      sendDepositTx,
		Flags: []cli.Flag{
			HttpPath,
			Password,
			ContractFlag,
			PrivKeyString,
			KeystoreFlag,
			AmountFlag,
		},
	}
)

func deployDepositContract(cli *cli.Context) {
	contractAddress := deposit.DeployContract(cli.String(HttpPath.Name),
		cli.String(PrivKeyString.Name),
		cli.Int64(DepositsForChainStart.Name),
		cli.Int64(MinDepositAmount.Name),
		cli.Int64(MaxDepositAmount.Name),
		30)
	output := cli.String(OutputFile.Name)
	if output != "" {
		err := ioutil.WriteFile(output, []byte(contractAddress), 0644)
		if err != nil {
			log.Fatalf("Error writing file to disk %v", err)
		}
	}
}

func sendDepositTx(cli *cli.Context) {
	deposit.SendDepositTx(cli.Int64(AmountFlag.Name), cli.String(HttpPath.Name), cli.String(PrivKeyString.Name), cli.String(ContractFlag.Name), cli.String(KeystoreFlag.Name), cli.String(Password.Name))
}
