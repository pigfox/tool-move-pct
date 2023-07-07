package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func getBalance(walletAddress string) (*big.Int, error) {
	// Convert the wallet address string to a common.Address
	address := common.HexToAddress(walletAddress)

	// Get the account balance
	balance, err := connection.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}
