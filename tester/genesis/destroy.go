package genesis

import (
	"log"
	"net/http"
)

func DestroyTestnet(testnetId string) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8000/testnets/"+testnetId, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error deleting a testnet", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("There was an error deleting the testnet: %d", resp.StatusCode)
	}
}
