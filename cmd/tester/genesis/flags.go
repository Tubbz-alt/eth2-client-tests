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
		Usage: "[artemis|lighthouse|lodestar|prysm]",
		Value: "prysm",
	}
	LogFolderFlag = cli.StringFlag{
		Name: "logFolder",
		Usage: "/var/log/testnet123",
		Value: "/var/log/testnet",
	}
	NumberOfNodesFlag = cli.IntFlag{
		Name:  "numNodes",
		Usage: "Number of nodes to deploy",
		Value: 4,
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
	TestnetName = cli.StringFlag{
		Name:  "testnet",
		Usage: "ID of the testnet to destroy",
	}
)
