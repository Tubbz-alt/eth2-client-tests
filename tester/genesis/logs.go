package genesis

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Reads remotely a log file from a node running in the testnet
func GetLogContents(testnetId string, node string) []byte {
	resp, err := http.Get("http://localhost:8000/log/" + testnetId + "/" + node)
	if err != nil {
		log.Fatal("Error requesting log", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("There was an error communicating with the testnet ", err)
	}
	logs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("There was an error reading the response from genesis", err)
	}
	return logs
}
