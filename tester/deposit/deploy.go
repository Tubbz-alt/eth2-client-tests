package deposit

import (
	"bytes"
	"context"
	"fmt"
	"github.com/prysmaticlabs/go-ssz"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	prysmKeyStore "github.com/prysmaticlabs/prysm/shared/keystore"
	"github.com/prysmaticlabs/prysm/shared/params"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	contracts "github.com/prysmaticlabs/prysm/contracts/deposit-contract"
)

func DeployContract(httpPath string, privKeyString string, depositsForChainStart int64, minDepositAmount int64, maxDepositAmount int64, customChainstartDelay int) string {
	var rpcClient *rpc.Client
	var err error
	var txOps *bind.TransactOpts

	rpcClient, err = rpc.Dial(httpPath)

	if err != nil {
		log.Fatalf("Error while creating RPC client %v with path %s", err, httpPath)
	}

	client := ethclient.NewClient(rpcClient)

	if privKeyString != "" {
		privKey, err := crypto.HexToECDSA(privKeyString)
		if err != nil {
			log.Fatal(err)
		}
		txOps = bind.NewKeyedTransactor(privKey)
		txOps.Value = big.NewInt(0)
		txOps.GasLimit = 4000000
		// User inputs keystore json file, sign tx with keystore json
	}

	drain := txOps.From

	txOps.GasPrice = big.NewInt(10 * 1e9 /* 10 gwei */)

	// Deploy validator registration contract
	addr, tx, _, err := contracts.DeployDepositContract(
		txOps,
		client,
		big.NewInt(depositsForChainStart),
		big.NewInt(minDepositAmount),
		big.NewInt(maxDepositAmount),
		big.NewInt(int64(customChainstartDelay)),
		drain,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Wait for contract to mine
	for pending := true; pending; _, pending, err = client.TransactionByHash(context.Background(), tx.Hash()) {
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}

	log.Println(fmt.Sprintf("New contract deployed %s", addr.Hex()))

	return addr.Hex()
}

func SendDepositTx(depositAmount int64, httpPath string, privKeyString string, depositContractAddr string, prysmKeystorePath string, password string) {
	var rpcClient *rpc.Client
	var err error
	var txOps *bind.TransactOpts

	rpcClient, err = rpc.Dial(httpPath)
	if err != nil {
		log.Fatal(err)
	}

	client := ethclient.NewClient(rpcClient)
	depositAmount = depositAmount * 1e9

	privKey, err := crypto.HexToECDSA(privKeyString)
	if err != nil {
		log.Fatal(err)
	}
	txOps = bind.NewKeyedTransactor(privKey)
	txOps.Value = big.NewInt(depositAmount)
	txOps.GasLimit = 4000000

	depositContract, err := contracts.NewDepositContract(common.HexToAddress(depositContractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	validatorKeys := make(map[string]*prysmKeyStore.Key)

	// Load from keystore
	store := prysmKeyStore.NewKeystore(prysmKeystorePath)
	prefix := params.BeaconConfig().ValidatorPrivkeyFileName
	validatorKeys, err = store.GetKeys(prysmKeystorePath, prefix, password)
	if err != nil {
		log.Fatalf("Could not get keys from keystore %s: %v", prysmKeystorePath, err)
	}

	for _, validatorKey := range validatorKeys {
		data, err := prysmKeyStore.DepositInput(validatorKey, validatorKey)
		if err != nil {
			log.Fatalf("Could not generate deposit input data: %v", err)
		}

		serializedData := new(bytes.Buffer)
		if err := ssz.Encode(serializedData, data); err != nil {
			log.Fatalf("could not serialize deposit data: %v", err)
		}

		tx, err := depositContract.Deposit(txOps, serializedData.Bytes())
		if err != nil {
			log.Fatalf("unable to send transaction to contract %v", err)
		}

		log.Println(fmt.Sprintf("Deposit sent to contract address %v for validator with a public key %#x, tx hash %#x", depositContractAddr, validatorKey.PublicKey.Marshal(), tx.Hash()))
	}
}
