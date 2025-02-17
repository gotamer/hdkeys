package hdkeys

import (
	"github.com/btcsuite/btcd/btcutil/bech32"
)

const CoinTypeNostr CoinType = 0x800004d5

func NostrAddress(privateKeyBytes []byte) (address string, err error) {

	bits5, err := bech32.ConvertBits(privateKeyBytes, 8, 5, true)
	if err == nil {
		address, err = bech32.Encode("nsec", bits5)
	}
	return
}

/*
func EncodePrivateKey(privateKeyHex string) (string, error) {
	b, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key hex: %w", err)
	}

	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}

	return bech32.Encode("nsec", bits5)
}

func EncodePublicKey(publicKeyHex string) (string, error) {
	b, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key hex: %w", err)
	}

	bits5, err := bech32.ConvertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}

	return bech32.Encode("npub", bits5)
}
*/
