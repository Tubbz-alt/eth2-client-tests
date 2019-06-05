package genesis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/eth2-client-tests/tester/genesis"
	"github.com/urfave/cli"
)

var (
	GenesisCommand = cli.Command{
		Name:        "genesis",
		Usage:       "Commands to set up genesis",
		Description: `Commands to set up and test genesis`,
		Subcommands: []cli.Command{
			cli.Command{
				Name:        "up",
				Usage:       "Tests genesis is available",
				Description: `Tests that genesis is available on port 8000`,
				Action:      testGenesisAvailable,
				Flags:       []cli.Flag{},
			},
			cli.Command{
				Name:        "destroy",
				Usage:       "Destroys a tesnet",
				Description: `Destroys the nodes of a testnet`,
				Action:      destroyTestnet,
				Flags: []cli.Flag{
					TestnetName,
				},
			},
			cli.Command{
				Name:        "testnet",
				Usage:       "Deploys a new testnet",
				Description: `Deploys a new testnet to genesis`,
				Action:      deployTestnet,
				Flags: []cli.Flag{
					BlockchainFlag,
					FileOutputFlag,
					LogFolderFlag,
					NumberOfNodesFlag,
					PortsFlag,
					VolumesFlag,
					ContractFlag,
					ValidatorsPassword,
				},
			},
			cli.Command{
				Name: "build-status",
				Usage: "Provides information on the current state of the build in genesis, and waits for the build to complete",
				Description: "This command queries genesis for the status of the build, and displays the information, every 5s, until the build completes or errors out",
				Action: buildStatus,
				Flags: []cli.Flag{
					TestnetName,
				},
			},
		},
	}
)

func destroyTestnet(ctx *cli.Context) {
	testnet := ctx.String(TestnetName.Name)
	genesis.DestroyTestnet(testnet)
}

func deployTestnet(ctx *cli.Context) {
	blockchain := ctx.String(BlockchainFlag.Name)
	logFolder := ctx.String(LogFolderFlag.Name)
	output := ctx.String(FileOutputFlag.Name)
	volumes := ctx.StringSlice(VolumesFlag.Name)
	ports := ctx.StringSlice(PortsFlag.Name)
	contract := ctx.String(ContractFlag.Name)
	validatorsPassword := ctx.String(ValidatorsPassword.Name)

	genesis.DeployTestnet(blockchain, logFolder, genesis.Images[blockchain], volumes, ports, ctx.Int(NumberOfNodesFlag.Name), contract, validatorsPassword, output)
}

func testGenesisAvailable(ctx *cli.Context) {
	resp, err := http.Get("http://localhost:8000/servers")
	if err != nil {
		log.Fatal("There was an error contacting genesis", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("The genesis server returned an error: %d", resp.StatusCode)
	}
	log.Printf("Genesis server contacted successfully")
}

// {"progress":100.000000,"error":null,"stage":"Finished","frozen":false}
type BuildStatus struct {
  Progress float32 `json:"progress"`
  Error interface{} `json:"error"`
  Stage string `json:"stage"`
  Frozen bool `json:"frozen"`
}

func buildStatus(ctx *cli.Context) {
	for
	{
		resp, err := http.Get("http://localhost:8000/status/build/" + ctx.String(TestnetName.Name))
		if err != nil {
			log.Fatal("There was an error contacting genesis", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("There was an error reading the response from genesis", err)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("The genesis server returned an error: %d", resp.StatusCode)
		}

		var buildStatus BuildStatus
		err = json.Unmarshal(body, &buildStatus)
		if err != nil {
			log.Fatalf("Unexpected issue unmarshalling json %v", err)
		}
		log.Printf("%s", string(body))
		if buildStatus.Progress < 100 && buildStatus.Error == nil {
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
}
