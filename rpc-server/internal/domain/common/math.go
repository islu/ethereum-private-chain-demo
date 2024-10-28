package common

import "math/big"

func FromWei(wei *big.Int) *big.Float {
	eth := new(big.Float).SetInt(wei)
	weiInEth := new(big.Float).Quo(eth, big.NewFloat(1e18))
	return weiInEth
}
