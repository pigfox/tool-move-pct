package main

import "math/big"

type Config struct {
	NetType string
	ChainID *big.Int
}

func getMainConfig() Config {
	return Config{
		NetType: "mainnet",
		ChainID: big.NewInt(1),
	}
}

func getTestConfig() Config {
	return Config{
		NetType: "sepolia",
		ChainID: big.NewInt(11155111),
	}
}
