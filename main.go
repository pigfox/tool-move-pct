package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var connection *ethclient.Client
var currentConfig Config
var movePercent int

func setUp() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	movePercent = 50
	currentConfig = getTestConfig()
	getConnection()
}

func main() {
	setUp()
	// Replace the wallet addresses and private key with the actual values
	privateKey := ""
	fromAddress := ""
	toAddress := ""

	// Get the balance of the from address
	balance, err := getBalance(fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Balance of %s: %s ETH", fromAddress, balance)

	balance, err = getBalance(toAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Balance of %s: %s ETH", toAddress, balance)

	// Transfer all funds from the from address to the to address
	err = transferFunds(privateKey, toAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Funds transferred from %s to %s", fromAddress, toAddress)

	balance, err = getBalance(toAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Balance of %s: %s ETH", toAddress, balance)
}

func getConnection() {
	//client, err := ethclient.Dial("https://" + currentConfig.NetType + ".infura.io/v3/8cfea7aaa1384f1a87b6d5aa65119ea3")
	//client, err := ethclient.Dial("https://rpc.ankr.com/eth_sepolia/2a3e2875532c8e9908052db93d1877b1d8f20966f7a764f80e5633cde3c84b89")
	client, err := ethclient.Dial("https://muddy-powerful-shard.ethereum-sepolia.quiknode.pro/a32eb2172495349b879e8967eeed06d378c95405/")
	if err != nil {
		log.Fatal(err)
	}
	connection = client
}
