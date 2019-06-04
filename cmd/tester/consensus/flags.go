package consensus

import "github.com/urfave/cli"

var (
	Folder = cli.StringFlag{
		Name:  "folder",
		Usage: "Folder containing the logs",
	}
	Type = cli.StringFlag{
		Name:  "type",
		Usage: "Type of consensus test to perform [finalized_block_root|finalized_state_root|justified_block_root|justified_state_root]",
		Value: "finalized_block_root",
	}
	BlockchainFlag = cli.StringFlag{
		Name:  "blockchain",
		Usage: "[artemis|lighthouse|lodestar|prysm]",
		Value: "prysm",
	}
	TestOutputFile = cli.StringFlag{
		Name:  "testoutput",
		Usage: "JUnit XML output file",
	}
)
