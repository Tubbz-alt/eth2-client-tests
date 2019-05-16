package genesis

import "github.com/urfave/cli"

var (
	FileOutputFlag = cli.StringFlag{
		Name:  "file",
		Usage: "File to write the testnet ID to ",
		Value: "",
	}
	BlockchainFlag = cli.StringFlag{
		Name:  "blockchain",
		Usage: "Blockchain ",
		Value: "prysm",
	}
	NumberOfNodesFlag = cli.IntFlag{
		Name:  "numNodes",
		Usage: "Number of nodes to deploy",
		Value: 3,
	}
	VolumesFlag = cli.StringSliceFlag{
		Name:  "volume",
		Usage: "Volumes to be mounted on all nodes. This flag can be used multiple times.",
		Value: nil,
	}
	PortsFlag = cli.StringSliceFlag{
		Name:  "ports",
		Usage: "Ports to be exposed on nodes. This flag can be used multiple times.",
		Value: nil,
	}
	TesnetIdFlag = cli.StringFlag{
		Name: "testnetId",
		Usage: "ID of the testnet to destroy",
	}
)
