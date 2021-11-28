package utils

import (
	"github.com/btcsuite/btcutil/bech32"
)

func DecodeBech32Address(addr string) (string, error) {
	hrp, _, err := bech32.Decode(addr)
	if err != nil {
		return "", err
	}
	return hrp, nil
}
