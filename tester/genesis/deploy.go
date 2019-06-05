package genesis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type TestnetParameter struct {
	LogFolder           string `json:"logFolder"`
	InstrumentationPort string `json:"prometheusInstrumentationPort"`
	Contract            string `json:"contract"`
	ValidatorsPassword  string `json:"validatorsPassword"`
}

type TestnetResource struct {
	Cpus    string   `json:"cpus"`
	Memory  string   `json:"memory"`
	Volumes []string `json:"volumes"`
	Ports   []string `json:"ports"`
}

type Testnet struct {
	Servers    []int             `json:"servers"`
	Blockchain string            `json:"blockchain"`
	Nodes      int               `json:"nodes"`
	Images     []string          `json:"images"`
	Resources  []TestnetResource `json:"resources"`
	Params     TestnetParameter  `json:"params"`
}

func DeployTestnet(blockchain string, logFolder string, images []string, volumes []string, ports []string, nodes int, contract string, validatorsPassword string, output string) {
	mountPoint := ""
	if logFolder != "" {
		mountPoint = "/var/output"
		volumes = append(volumes, logFolder+":"+mountPoint)
	}
	testNet := Testnet{
		[]int{1},
		blockchain,
		nodes,
		images,
		[]TestnetResource{
			{"", "", volumes, ports},
		},
		TestnetParameter{
			mountPoint,
			"8088",
			contract,
			validatorsPassword,
		},
	}
	json, err := json.Marshal(testNet)
	if err != nil {
		log.Fatal("Error preparing testnet configuration", err)
	}
	log.Printf(string(json))
	resp, err := http.Post("http://localhost:8000/testnets", "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error sending a testnet configuration", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("There was an error deploying the testnet", err)
	}
	testnetId, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	log.Printf("Testnet deployed with id %s", testnetId)
	if output != "" {
		err := ioutil.WriteFile(output, testnetId, 0644)
		if err != nil {
			log.Fatal("There was an error saving testnet id to file", err)
		}
	}
}
