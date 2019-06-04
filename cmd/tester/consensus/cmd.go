package consensus

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ethereum/eth2-client-tests/tester/consensus"
	"github.com/urfave/cli"
)

var (
	ConsensusTestCommand = cli.Command{
		Name:        "consensus",
		Usage:       "Tests consensus among hosts",
		Description: `Tests that all the hosts are in consensus`,
		Action:      consensusTest,
		Flags: []cli.Flag{
			Folder,
			Type,
			BlockchainFlag,
			TestOutputFile,

		},
	}
)

func consensusTest(cli *cli.Context) {
	folder := cli.String(Folder.Name)
	fileInfos, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	files := []string{}
	for _, fileInfo := range fileInfos {
		files = append(files, filepath.Join(folder, fileInfo.Name()))
	}
	switch cli.String(Type.Name) {
	case "finalized_block_root":
		consensus.CheckFinalizedBlockRoot(cli.String(TestOutputFile.Name), cli.String(BlockchainFlag.Name), files)
	case "finalized_state_root":
		consensus.CheckFinalizedStateRoot(cli.String(TestOutputFile.Name), cli.String(BlockchainFlag.Name), files)
	case "justified_block_root":
		consensus.CheckJustifiedBlockRoot(cli.String(TestOutputFile.Name), cli.String(BlockchainFlag.Name), files)
	case "justified_state_root":
		consensus.CheckJustifiedStateRoot(cli.String(TestOutputFile.Name), cli.String(BlockchainFlag.Name), files)
	default:
		log.Fatalf("Invalid type %s", cli.String(Type.Name))
	}

}
